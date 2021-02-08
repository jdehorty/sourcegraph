// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1000000000_init.down.sql (19B)
// 1000000000_init.up.sql (19B)
// 1000000001_initial_schema.down.sql (475B)
// 1000000001_initial_schema.up.sql (4.608kB)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __1000000000_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initDownSql,
		"1000000000_init.down.sql",
	)
}

func _1000000000_initDownSql() (*asset, error) {
	bytes, err := _1000000000_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000000_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initUpSql,
		"1000000000_init.up.sql",
	)
}

func _1000000000_initUpSql() (*asset, error) {
	bytes, err := _1000000000_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000001_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x39\x45\xee\x91\x95\xb5\x51\x02\xb6\x15\x9b\x45\x77\x21\xd2\x21\x0c\x98\xa4\x26\x29\x78\x7c\x11\xb1\x82\x32\x68\x77\xb3\xf8\xef\xff\x37\x95\xdc\xab\x56\x30\x56\x9f\xba\x23\x57\x6d\x2d\x07\xae\x76\x5c\x0e\xaa\xd7\x3d\xcf\x90\x10\xb2\x99\x22\x86\x92\x4d\x82\x29\x1a\x1c\xcd\xb9\x24\x00\xf1\x37\x11\xac\x87\x75\x58\x4c\xe8\x30\xd8\x0b\xcd\xeb\x4d\x75\x90\x14\x4f\x7d\xb3\xb4\xe5\x67\xe7\x1c\xf0\x3a\x3f\xaa\x6f\x84\xd5\x27\x50\x92\xf3\x84\xc0\x3b\x4a\xad\x7b\x28\x76\xb4\xc5\x9a\xe5\xf8\xb9\xff\x8d\x38\x0c\x84\xc0\x2b\x22\x18\xdb\x76\x4d\xa3\xb4\x60\xf7\x00\x00\x00\xff\xff\xe1\xb5\x30\xdc\xdb\x01\x00\x00")

func _1000000001_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initial_schemaDownSql,
		"1000000001_initial_schema.down.sql",
	)
}

func _1000000001_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1000000001_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_initial_schema.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2c, 0xd2, 0x4, 0x8c, 0x45, 0xac, 0xd5, 0xee, 0x57, 0x8f, 0xba, 0x64, 0x82, 0x72, 0x50, 0x52, 0x95, 0x91, 0x23, 0xf5, 0x42, 0xf6, 0x3, 0x65, 0xb, 0x7, 0x31, 0xdc, 0xe5, 0x12, 0x12, 0x91}}
	return a, nil
}

var __1000000001_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x5b\x73\xdb\xba\x11\x7e\xd7\xaf\xd8\x39\x0f\xb5\x34\x23\xd3\xef\x76\xe7\xa4\x8a\xc4\xe4\xb0\xc7\xa6\x52\x49\x99\xe6\xb4\xd3\xd1\x40\xe4\x8a\x44\x02\x02\x0c\x00\x5a\x56\x2f\xff\xbd\xb3\x00\x78\xb3\xe5\x24\x73\xa6\xf5\x8b\x35\xc2\xe2\xdb\x6f\xef\x0b\xbd\x8d\xdf\x27\xe9\xdd\x64\xb2\xdc\xc4\x8b\x5d\x0c\xf1\xa7\x5d\x9c\x6e\x93\x75\x0a\xc9\x3b\x48\xd7\x3b\x88\x3f\x25\xdb\xdd\x16\x2c\xaf\xd0\x64\x4c\x60\x7e\xb8\xfb\x9e\x6c\x5d\xec\xad\x2e\xaa\xef\xca\x65\xdc\xe2\x93\xbd\x9b\x4c\xae\xaf\x61\x83\x99\xd2\xb9\x01\x8d\xb5\x32\xdc\x2a\x7d\x06\xc9\x2a\x34\x73\x38\x28\x5b\x42\xc9\x8d\x55\x9a\x67\x4c\x00\x93\x39\xd4\x1a\x0d\x4a\x3b\x87\xc6\x70\x59\x00\x83\x46\xf2\xaf\x0d\x0e\x6f\xef\xe9\xfa\x1e\x92\x15\xa1\x4f\x1b\xa9\x51\x30\x8b\x39\x58\x05\xb6\x1c\x49\x26\xab\x68\xd6\x52\xdd\x2d\xde\xde\xc7\xee\xd0\xdd\x37\x30\x9d\x00\x00\x5c\x5f\xc3\xae\xbc\x0c\x1f\x39\x01\x9e\xc3\x81\x17\x06\x35\x67\xc2\x99\x98\x7e\xbc\xbf\x87\x0f\x9b\xe4\x61\xb1\xf9\x0d\x7e\x8d\x7f\x9b\x4f\x86\x40\x74\x79\x0e\x56\xf3\x42\xb3\xea\x9a\xcb\x1c\x9f\x30\x87\xa3\xd2\x70\x64\xc6\x02\x46\x45\x04\x1a\x0b\x7c\xaa\xe1\xc8\x85\x45\xcd\x65\xe1\x15\xd1\xcd\xe0\xb8\x4e\x4f\xc0\x5e\xae\xd3\xed\x6e\xb3\x48\xd2\x1d\x64\x25\x66\x5f\x3c\x45\xa9\x24\x56\xb5\x3d\xc3\xf2\x97\x78\xf9\x2b\x4c\xa7\x0e\x61\xfd\x21\xde\x2c\x76\xeb\xcd\xf4\x8f\x3f\xcf\xe0\xea\xea\xf6\xd6\x43\xce\x66\x93\x99\x8f\x47\x2c\x8f\x4a\x67\x08\xb6\x64\xd6\x47\x02\x98\xc6\xe0\xe7\xa8\x75\xd7\xc7\x34\xf9\xcb\xc7\x18\x92\x74\x15\x7f\x1a\x78\xcd\x6b\xf6\xb2\x7b\x9e\x3f\xc1\x3a\x1d\x9c\x3a\x06\x41\xcd\x52\x23\xb3\xd8\x7a\x02\xbc\x27\x8c\xf3\xc4\xb3\x4c\xe8\x1d\x01\x07\x66\x30\x07\x25\x87\x7e\x32\x1d\xa7\xcb\x64\x28\x21\xc7\x34\xe0\xe3\x36\x49\xdf\x43\xc1\x25\x4c\x85\x3a\xa1\xf6\xae\x99\xdd\xde\x3a\x4f\xd0\x81\xbb\xb5\x57\xb5\x21\xb6\xc3\x2c\x65\xfa\xc0\xad\x66\xfa\x0c\x15\x5a\x96\x33\xcb\x80\x1d\x54\x63\x01\x1f\x51\x5a\x13\xc1\xd6\x2a\x8d\x39\x70\x09\x0c\x0c\xd6\x4c\x3b\x2b\xd9\x41\x20\x30\x03\xdc\x02\x37\xa0\x8e\x16\x25\x11\x22\x17\xe4\x04\x4f\x66\x57\x8d\xb0\xbc\x16\xd8\x42\x8d\x33\xb3\x53\xf7\xbf\xcf\xcb\x87\xb1\x25\xb6\xe4\xc6\x73\x98\xfb\xcf\x19\x93\x70\x40\x60\xf2\x3c\x30\xff\xcf\xdb\x75\xda\x93\x3a\x95\x3c\x2b\xe1\xc4\x85\x20\x49\x8d\xb6\xd1\x12\xf3\x56\xc1\xa9\x44\x09\x5f\x1b\xd4\x67\x0a\xa2\x37\x6f\xee\xca\x39\x40\xfb\x08\xfb\xd0\xd2\xd7\x85\x56\x4d\x8d\x79\x28\xf2\xcf\x46\xc9\x03\xa8\x1a\x35\xb3\x4a\x1b\x78\x33\x87\x37\x7f\x98\xc3\x9b\x7f\xcf\x5b\x05\x74\xe7\x4f\x3f\x47\xb0\x23\xba\xa6\x54\x8d\xc8\x09\xd6\x54\x4c\x08\x70\x04\x95\x14\xe7\x39\xd4\x9a\x57\x44\xbe\x31\x08\x19\x33\x48\xc1\xf0\x42\x82\x1b\x6b\xc0\x34\x59\x09\xcc\xdc\x06\xdc\x16\x1e\xfe\xf5\xd3\x67\xf6\xc8\xf6\x8f\xa8\x0d\x57\xd2\xfc\x74\x0b\x7f\x8f\xa2\xe8\x1f\xff\x19\x08\x08\x26\x8b\x86\x15\x48\x87\xf4\xf7\x42\xa0\x6e\x84\xd8\x6b\xfc\xda\xa0\xb1\x17\x11\x98\x94\xca\x32\x1b\x14\x3c\x43\x70\xff\x3a\x77\x7b\x8f\xb4\x51\xbd\x58\xb9\x9d\x2c\x37\xdf\xac\xdd\x56\x6e\xdf\x7d\x18\x57\x6f\xfb\xf5\xb4\xfd\x10\x94\x25\x54\xb0\xbd\x16\xab\x40\xd5\x96\x57\xfc\x9f\x08\x7f\xfd\x25\xde\xc4\x90\x09\xd6\x18\x34\x70\xe2\xb6\x0c\x84\xfb\xc0\x85\x88\xf5\x41\x7d\x56\xc4\x2f\x59\x51\xb5\x0e\x33\xce\xd7\xf0\xfb\x24\x85\xe7\xcc\xda\x52\xf5\x79\x06\xea\x11\xb5\x1b\x62\xc0\x8c\x51\x19\x77\xb3\xc0\x91\x62\xc3\xf2\x99\x2a\x0d\xd4\x31\xe7\xc0\x23\x8c\xa0\x10\xea\xc0\x84\x38\xcf\x28\x79\x35\x52\x31\x73\x59\x08\x24\x05\xb2\xa9\xd0\x4f\xa4\x47\x26\x1a\x97\x44\x85\x72\xd3\x28\x54\x07\x17\x67\x68\x6a\x67\x63\xae\x4e\x32\x9a\x5c\x5f\x7b\x62\x9d\xb6\x96\x0a\x57\x92\xae\x77\x7d\xcd\x8d\xbc\xd1\x84\x72\x28\x54\xda\x91\x2b\xf8\x64\xd5\xd6\x4c\x63\xfc\x4c\xd3\x78\x24\x03\x15\x69\x60\x60\x6a\xcc\xf8\x91\x67\x03\x90\x39\x28\x0d\x42\xa9\x2f\x4d\xed\x06\x60\xd6\x68\x8d\xd2\xf7\x76\x50\xc7\xb1\x1b\xd8\xd1\xa2\xa6\x36\x55\x32\x03\x07\xc4\xae\xd5\x92\x74\x4e\x96\x74\x63\xec\x35\x22\x4e\x49\x0b\xfe\x6c\xe2\x32\xeb\xbe\x71\xe1\x08\xa7\x2e\x4c\x57\x06\x32\x1a\x08\x5c\xc9\x79\xdb\x0f\xf1\x89\x55\xd4\x0e\x09\x51\x33\x97\xd7\x08\x59\xc9\x64\x81\xbe\xbd\x16\xac\x29\x10\x0e\x2c\xfb\x42\x32\x23\x33\x0e\x48\xf1\xe8\x58\x8f\x3a\x29\xf5\x44\x34\xfb\x5a\x71\x4a\x8f\xae\x9d\x2e\xda\x55\x82\xe7\x28\x2d\x3f\x72\xd4\x8e\x06\x69\xf5\x57\x88\xb1\xcb\xbd\x16\x9d\xf2\x0c\xf3\xd0\x76\xb8\x01\xa9\x2c\x30\x49\x31\x3a\x6a\x55\x01\x93\xca\x96\xa8\x5b\x05\x6e\x08\xcc\xe1\xd0\x58\xd0\x8c\x0e\xe0\x73\x63\x6c\xbf\xc2\xf4\x7a\x7d\x07\x0f\x3c\x39\x0d\x13\x8b\x05\xea\xf1\x26\xe1\xf6\x32\xcb\xaa\xba\x77\xb3\xe7\xe3\x3d\xea\x31\x9c\xa3\x77\xc9\x43\xbc\xdd\x2d\x1e\x3e\xec\xfe\xf6\x7c\x73\x08\x58\x47\xa1\x98\x25\x9b\x9c\x53\x42\x5a\xbf\x16\x2b\x8f\xec\x65\x72\xd5\xd0\x64\xab\x35\x66\x9c\x5a\xe3\x05\xfc\x45\x5f\x75\x5d\xf1\x7a\xbf\xf6\x63\x86\x1f\x69\xba\x44\xa3\x26\xf7\xaa\xe1\xe3\xf2\x98\x3a\x57\x13\xb9\x8a\x51\x56\xd4\xb5\xe0\x99\x2f\xac\xd5\xdb\xd9\xc8\x88\xce\x02\x38\xb1\x90\x6f\x14\xbd\x54\x59\xec\x62\x54\x86\x0b\x03\x25\x15\x3b\x83\x54\x20\x94\x2c\x90\xb2\x92\x1b\x0b\x37\x94\xf8\x8f\x4c\xf0\x9c\x34\xb8\xd1\xe6\x74\xcc\xa1\x54\x27\x7c\x6c\x43\xd8\x8f\x10\xd9\x08\x41\x66\x8e\x39\x50\xc2\x90\x2f\xda\xee\x32\x6a\x48\xae\x0f\xb1\xd0\x89\x7c\xb2\xcf\x3c\xac\x5b\x65\x5e\x73\x4f\xa5\x8c\xa5\x5c\x40\x69\xc5\x19\xbe\x48\x75\x92\x61\x8d\x0a\xc9\x3c\xec\x0c\x4d\x9d\xbb\xc8\xd4\xa8\xb9\xca\xa9\xa7\x89\xb3\x2b\xa6\x2c\x53\x8d\xf4\xe4\xa8\x01\xb4\x0a\x06\xfc\x7c\x71\x99\x08\x92\x17\x55\x4e\xa6\xe5\x28\xd0\x62\x1e\x56\x08\x9a\xb3\x96\xb6\x03\x7b\x99\x61\xe7\x26\x6a\x74\xff\x77\xd7\xb9\x4d\xe9\x07\xd2\xcb\xb9\xcd\x6f\x6d\xa4\xd2\xfb\xf2\x07\x32\x2a\xb1\x2e\x69\x4a\xf6\x88\xbe\x87\x86\x46\xd4\xaa\x31\x5c\x66\xc1\x4c\xa5\x79\xc1\x25\xa3\xcd\xe0\x5b\xc4\x62\x69\x1a\x8d\xe4\x05\x25\x3d\xc5\xd1\x08\x39\x72\x14\xb9\x73\xb2\x6f\xff\xe4\x77\xda\x6a\x98\x0e\x6a\x5e\xbc\x10\x9c\x3a\x77\xcd\xec\xc3\xa5\xfe\xa5\xe0\xae\xd0\xdf\x74\xda\xa6\x5a\xb2\x75\x65\x3d\x83\x45\xba\x82\xe9\x88\xec\xf8\xe8\xb2\x41\xad\xcc\xac\x83\x5e\x6f\x5e\xd1\x12\x1a\xc8\x6b\x9a\xc6\xc7\xaf\x6b\x6b\xe5\xbc\xc6\x59\x70\xe5\xbb\xf5\x26\x4e\xde\xa7\xb4\xff\xf6\x7b\xc3\x9e\xe7\x33\xd8\xc4\xef\xe2\x4d\x9c\x2e\xe3\x6d\xbf\xf2\xd0\xf7\xeb\x14\x56\xf1\x7d\xbc\x8b\x61\xb9\xd8\x2e\x17\xab\x18\x56\x24\xb9\xa1\x51\x32\x7f\x89\x39\xe4\x31\x02\x1d\xbc\x82\x7e\x07\xec\x45\x3b\x7f\x2f\xfe\x64\xfc\x04\x2b\xcf\x35\xea\x30\x9c\x6a\xa6\x2d\xa7\x94\xea\x17\x75\x38\xf8\xd6\xe6\x66\xff\x16\x11\x4a\x6b\x6b\x73\x7b\x73\x93\xab\xcc\x44\xdd\xaf\x03\x51\xa6\xaa\x1b\x7a\x68\x1b\x7b\xe3\xd6\xf6\xeb\xc1\x0f\x07\x37\xbd\x0e\x33\xd9\xc6\xf7\xf1\x72\x17\xaa\x65\xdf\x9f\x4c\xaf\x46\x93\xf9\x6a\x0e\x57\x04\x71\x35\x66\x7b\xb0\x1a\xf1\xb5\xe7\xe2\xe0\xc9\x3c\x5a\x27\x47\xc0\xfb\x90\x6d\x7b\x0f\xb5\x4e\x9f\x6d\x04\x7e\xb3\xf4\x87\x6d\x66\xce\xee\xbe\x0b\x18\x82\xf2\xe3\xa8\x6d\x14\xbf\x09\x7d\x31\xf0\x3f\xa4\xe3\x72\xca\xdc\x4d\x26\xcb\xf5\xc3\x43\xb2\xbb\x9b\xfc\x37\x00\x00\xff\xff\x60\x0c\x23\xbc\x00\x12\x00\x00")

func _1000000001_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initial_schemaUpSql,
		"1000000001_initial_schema.up.sql",
	)
}

func _1000000001_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1000000001_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_initial_schema.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbd, 0x80, 0xef, 0x57, 0xde, 0xe0, 0x13, 0x78, 0xc8, 0xad, 0x78, 0x9, 0x5, 0xa6, 0x95, 0x19, 0x55, 0xb0, 0x53, 0xd7, 0x79, 0x95, 0xbb, 0xe5, 0xed, 0xaa, 0xda, 0xc5, 0x43, 0xb1, 0xc, 0x63}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"1000000000_init.down.sql":           _1000000000_initDownSql,
	"1000000000_init.up.sql":             _1000000000_initUpSql,
	"1000000001_initial_schema.down.sql": _1000000001_initial_schemaDownSql,
	"1000000001_initial_schema.up.sql":   _1000000001_initial_schemaUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"1000000000_init.down.sql":           {_1000000000_initDownSql, map[string]*bintree{}},
	"1000000000_init.up.sql":             {_1000000000_initUpSql, map[string]*bintree{}},
	"1000000001_initial_schema.down.sql": {_1000000001_initial_schemaDownSql, map[string]*bintree{}},
	"1000000001_initial_schema.up.sql":   {_1000000001_initial_schemaUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
