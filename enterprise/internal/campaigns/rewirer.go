package campaigns

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/repos"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/campaigns"
	"github.com/sourcegraph/sourcegraph/internal/db"
	"github.com/sourcegraph/sourcegraph/internal/types"
)

// RewirerMapping maps a connection between ChangesetSpec and Changeset.
// If the ChangesetSpec doesn't match a Changeset (ie. it describes a to-be-created Changeset), ChangesetID is 0.
// If the ChangesetSpec is 0, the Changeset will be non-zero and means "to be closed".
// If both are non-zero values, the changeset should be updated with the changeset spec in the mapping.
type RewirerMapping struct {
	ChangesetSpecID int64
	ChangesetSpec   *campaigns.ChangesetSpec
	ChangesetID     int64
	Changeset       *campaigns.Changeset
	RepoID          api.RepoID
}

type RewirerMappings []*RewirerMapping

func (rm RewirerMappings) Hydrate(ctx context.Context, store *Store) error {
	changesetSpecs, _, err := store.ListChangesetSpecs(ctx, ListChangesetSpecsOpts{
		IDs: rm.ChangesetSpecIDs(),
	})
	if err != nil {
		return err
	}
	changesets, _, err := store.ListChangesets(ctx, ListChangesetsOpts{IDs: rm.ChangesetIDs()})
	if err != nil {
		return err
	}
	changesetsByID := map[int64]*campaigns.Changeset{}
	changesetSpecsByID := map[int64]*campaigns.ChangesetSpec{}

	for _, c := range changesets {
		changesetsByID[c.ID] = c
	}
	for _, c := range changesetSpecs {
		changesetSpecsByID[c.ID] = c
	}

	for _, m := range rm {
		if m.ChangesetID != 0 {
			m.Changeset = changesetsByID[m.ChangesetID]
		}
		if m.ChangesetSpecID != 0 {
			m.ChangesetSpec = changesetSpecsByID[m.ChangesetSpecID]
		}
	}
	return nil
}

// ChangesetIDs returns a list of unique changeset IDs in the slice of mappings.
func (rm RewirerMappings) ChangesetIDs() []int64 {
	changesetIDMap := make(map[int64]struct{})
	for _, m := range rm {
		if m.ChangesetID != 0 {
			changesetIDMap[m.ChangesetID] = struct{}{}
		}
	}
	changesetIDs := make([]int64, len(changesetIDMap))
	for id := range changesetIDMap {
		changesetIDs = append(changesetIDs, id)
	}
	return changesetIDs
}

// ChangesetSpecIDs returns a list of unique changeset spec IDs in the slice of mappings.
func (rm RewirerMappings) ChangesetSpecIDs() []int64 {
	changesetSpecIDMap := make(map[int64]struct{})
	for _, m := range rm {
		if m.ChangesetSpecID != 0 {
			changesetSpecIDMap[m.ChangesetSpecID] = struct{}{}
		}
	}
	changesetSpecIDs := make([]int64, len(changesetSpecIDMap))
	for id := range changesetSpecIDMap {
		changesetSpecIDs = append(changesetSpecIDs, id)
	}
	return changesetSpecIDs
}

// RepoIDs returns a list of unique repo IDs in the slice of mappings.
func (rm RewirerMappings) RepoIDs() []api.RepoID {
	repoIDMap := make(map[api.RepoID]struct{})
	for _, m := range rm {
		repoIDMap[m.RepoID] = struct{}{}
	}
	repoIDs := make([]api.RepoID, len(repoIDMap))
	for id := range repoIDMap {
		repoIDs = append(repoIDs, id)
	}
	return repoIDs
}

type changesetRewirer struct {
	mappings RewirerMappings
	campaign *campaigns.Campaign
	tx       *Store
	rstore   repos.Store
}

func NewChangesetRewirer(mappings RewirerMappings, campaign *campaigns.Campaign, tx *Store, rstore repos.Store) *changesetRewirer {
	return &changesetRewirer{
		mappings: mappings,
		campaign: campaign,
		tx:       tx,
		rstore:   rstore,
	}
}

// Rewire uses RewirerMappings (mapping ChangesetSpecs to matching Changesets) generated by Store.GetRewirerMappings to update the Changesets
// for consumption by the background reconciler.
//
// It also updates the ChangesetIDs on the campaign.
func (r *changesetRewirer) Rewire(ctx context.Context) (changesets []*campaigns.Changeset, err error) {
	// First we need to load the associations.
	associations, err := r.loadAssociations(ctx)
	if err != nil {
		return nil, err
	}

	changesets = []*campaigns.Changeset{}

	for _, m := range r.mappings {
		// If we don't have access to a repository, we return an error. Why not
		// simply skip the repository? If we skip it, the user can't reapply
		// the same campaign spec, since it's already applied and re-applying
		// would require a new spec.
		repo, ok := associations.accessibleReposByID[m.RepoID]
		if !ok {
			return nil, &db.RepoNotFoundErr{ID: m.RepoID}
		}

		// If a Changeset that's currently attached to the campaign wasn't matched to a ChangesetSpec, it needs to be closed/detached.
		if m.ChangesetSpecID == 0 {
			changeset := m.Changeset

			// If the changeset is currently not attached to this campaign, we don't want to modify it.
			if !changeset.AttachedTo(r.campaign.ID) {
				continue
			}

			if err := r.closeChangeset(ctx, changeset); err != nil {
				return nil, err
			}

			continue
		}

		spec := m.ChangesetSpec

		if err := checkRepoSupported(repo); err != nil {
			return nil, err
		}

		var changeset *campaigns.Changeset

		if m.ChangesetID != 0 {
			changeset = m.Changeset
			if spec.Spec.IsImportingExisting() {
				r.attachTrackingChangeset(changeset)
			} else if spec.Spec.IsBranch() {
				r.updateChangesetToNewSpec(changeset, spec)
			}
		} else {
			if spec.Spec.IsImportingExisting() {
				changeset = r.createTrackingChangeset(repo, spec.Spec.ExternalID)
			} else if spec.Spec.IsBranch() {
				changeset = r.createChangesetForSpec(repo, spec)
			}
		}
		changesets = append(changesets, changeset)
	}

	return changesets, nil
}

func (r *changesetRewirer) createChangesetForSpec(repo *types.Repo, spec *campaigns.ChangesetSpec) *campaigns.Changeset {
	newChangeset := &campaigns.Changeset{
		RepoID:              spec.RepoID,
		ExternalServiceType: repo.ExternalRepo.ServiceType,

		CampaignIDs:       []int64{r.campaign.ID},
		OwnedByCampaignID: r.campaign.ID,
		CurrentSpecID:     spec.ID,

		PublicationState: campaigns.ChangesetPublicationStateUnpublished,
		ReconcilerState:  campaigns.ReconcilerStateQueued,
	}

	// Copy over diff stat from the spec.
	diffStat := spec.DiffStat()
	newChangeset.SetDiffStat(&diffStat)

	return newChangeset
}

func (r *changesetRewirer) updateChangesetToNewSpec(c *campaigns.Changeset, spec *campaigns.ChangesetSpec) {
	if c.ReconcilerState == campaigns.ReconcilerStateCompleted {
		c.PreviousSpecID = c.CurrentSpecID
	}
	c.CurrentSpecID = spec.ID

	// Ensure that the changeset is attached to the campaign
	c.CampaignIDs = append(c.CampaignIDs, r.campaign.ID)

	// Copy over diff stat from the new spec.
	diffStat := spec.DiffStat()
	c.SetDiffStat(&diffStat)

	// We need to enqueue it for the changeset reconciler, so the
	// reconciler wakes up, compares old and new spec and, if
	// necessary, updates the changesets accordingly.
	c.ResetQueued()
}

func (r *changesetRewirer) createTrackingChangeset(repo *types.Repo, externalID string) *campaigns.Changeset {
	newChangeset := &campaigns.Changeset{
		RepoID:              repo.ID,
		ExternalServiceType: repo.ExternalRepo.ServiceType,

		CampaignIDs:     []int64{r.campaign.ID},
		ExternalID:      externalID,
		AddedToCampaign: true,
		// Note: no CurrentSpecID, because we merely track this one

		PublicationState: campaigns.ChangesetPublicationStatePublished,

		// Enqueue it so the reconciler syncs it.
		ReconcilerState: campaigns.ReconcilerStateQueued,
		Unsynced:        true,
	}

	return newChangeset
}

func (r *changesetRewirer) attachTrackingChangeset(changeset *campaigns.Changeset) {
	// We already have a changeset with the given repoID and
	// externalID, so we can track it.
	changeset.AddedToCampaign = true
	changeset.CampaignIDs = append(changeset.CampaignIDs, r.campaign.ID)

	// If it's errored and not created by another campaign, we re-enqueue it.
	if changeset.OwnedByCampaignID == 0 && changeset.ReconcilerState == campaigns.ReconcilerStateErrored {
		changeset.ResetQueued()
	}
}

func (r *changesetRewirer) closeChangeset(ctx context.Context, changeset *campaigns.Changeset) error {
	if changeset.CurrentSpecID != 0 && changeset.OwnedByCampaignID == r.campaign.ID {
		// If we have a current spec ID and the changeset was created by
		// _this_ campaign that means we should detach and close it.

		// But only if it was created on the code host:
		if changeset.Published() {
			// Store the current spec also as the previous spec.
			//
			// Why?
			//
			// When a changeset with (prev: A, curr: B) should be closed but
			// closing failed, it will still have (prev: A, curr: B) set.
			//
			// If someone then applies a new campaign spec and re-attaches that
			// changeset with changeset spec C, the changeset would end up with
			// (prev: A, curr: C), because we don't rotate specs on errors in
			// `updateChangesetToNewSpec`.
			//
			// That would mean, though, that the delta between A and C tells us
			// to repush and update the changeset on the code host, in addition
			// to 'reopen', which would actually be the only required action.
			//
			// So, when we mark a changeset as to-be-closed, we also rotate the
			// specs, so that it changeset is saved as (prev: B, curr: B) and
			// when somebody re-attaches it it's (prev: B, curr: C).
			// But we only rotate the spec, if applying the currentSpecID was
			// successful:
			if changeset.ReconcilerState == campaigns.ReconcilerStateCompleted {
				changeset.PreviousSpecID = changeset.CurrentSpecID
			}
			changeset.Closing = true
			changeset.ResetQueued()
		} else {
			// otherwise we simply delete it.
			return r.tx.DeleteChangeset(ctx, changeset.ID)
		}
	}

	// Disassociate the changeset with the campaign.
	changeset.RemoveCampaignID(r.campaign.ID)
	return r.tx.UpdateChangeset(ctx, changeset)
}

type rewirerAssociations struct {
	accessibleReposByID map[api.RepoID]*types.Repo
}

// loadAssociations retrieves all entities required to rewire the changesets in a campaign.
func (r *changesetRewirer) loadAssociations(ctx context.Context) (associations *rewirerAssociations, err error) {
	if err := r.mappings.Hydrate(ctx, r.tx); err != nil {
		return nil, err
	}

	associations = &rewirerAssociations{}
	// Fetch all repos involved. We use them later to enforce repo permissions.
	//
	// 🚨 SECURITY: db.Repos.GetRepoIDsSet uses the authzFilter under the hood and
	// filters out repositories that the user doesn't have access to.
	associations.accessibleReposByID, err = db.Repos.GetReposSetByIDs(ctx, r.mappings.RepoIDs()...)
	if err != nil {
		return nil, err
	}

	return associations, nil
}
