package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	var build struct {
		Number int64
	}
	commit := os.Getenv("BUILDKITE_COMMIT")
	if commit == "" {
		panic("BUILDKITE_COMMIT env var not set")
	}
	{
		body, err := json.Marshal(map[string]interface{}{
			"commit":  "HEAD",
			"branch":  "master",
			"message": "Open source repository commit " + commit[0:7],
			"author": map[string]interface{}{
				"name": "Sourcegraph Bot",
			},
			"meta_data": map[string]interface{}{
				"OSS_REPO_REVISION": commit,
			},
		})
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", os.ExpandEnv("https://api.buildkite.com/v2/organizations/sourcegraph/pipelines/enterprise/builds?access_token=$BUILDKITE_TOKEN"), bytes.NewReader(body))
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&build)
		if err != nil {
			panic(err)
		}
	}
	for {
		req, err := http.NewRequest("GET", os.ExpandEnv(fmt.Sprintf("https://api.buildkite.com/v2/organizations/sourcegraph/pipelines/enterprise/builds/%v?access_token=$BUILDKITE_TOKEN", build.Number)), nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		var build struct {
			State string
		}
		err = json.NewDecoder(resp.Body).Decode(&build)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
		bs := build.State
		switch bs {
		case "passed":
			os.Exit(0)
		case "running", "scheduled":
			time.Sleep(1 * time.Second)
			continue
		default:
			fmt.Println("enterprise build ended with status:", bs)
			os.Exit(1)
		}
	}
}
