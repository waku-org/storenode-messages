// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1_setup.down.sql (106B)
// 1_setup.up.sql (580B)
// doc.go (74B)

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var __1_setupDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x4c\xa9\xf0\x2d\x4e\x37\xb4\xe6\x02\xcb\x86\x38\x3a\xf9\xb8\x22\xc9\x16\x57\xe6\x25\x87\xe4\x17\x64\x26\x07\x97\x24\x96\x94\x16\xe3\x50\x95\x9b\x59\x5c\x9c\x99\x97\xee\x9b\x5a\x5c\x9c\x98\x9e\x5a\x6c\xcd\x05\x08\x00\x00\xff\xff\x21\x92\x2c\xdf\x6a\x00\x00\x00")

func _1_setupDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_setupDownSql,
		"1_setup.down.sql",
	)
}

func _1_setupDownSql() (*asset, error) {
	bytes, err := _1_setupDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_setup.down.sql", size: 106, mode: os.FileMode(0664), modTime: time.Unix(1715268524, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3b, 0x6c, 0x2, 0x65, 0x57, 0xd3, 0xeb, 0xf4, 0xfc, 0x58, 0x8a, 0x59, 0x1e, 0x6c, 0xa, 0x59, 0xbc, 0x22, 0x76, 0x8d, 0x78, 0x88, 0x84, 0xdd, 0x12, 0xe, 0x90, 0x96, 0xe7, 0xe6, 0x26, 0xe0}}
	return a, nil
}

var __1_setupUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xbd\x6e\x83\x30\x14\x85\xe7\xf8\x29\xee\x08\x12\x4b\xe7\x4c\x2e\xb8\xc5\x2a\x98\xca\x76\x9a\x64\x74\xc0\xa2\x96\xc2\x8f\x72\x8d\xd4\xbe\x7d\x55\x1a\xa1\x94\x52\x96\xcc\xc7\xe7\xf3\x3d\x5f\x2c\x19\xd5\x0c\x34\x7d\xcc\x18\xf0\x27\x10\x85\x06\x76\xe0\x4a\x2b\xc0\xcf\xb6\xd4\x5d\xef\x4a\xe5\x8d\x1f\x10\x02\xb2\x29\xcf\x03\x7a\x7b\xe1\x15\x70\xa1\xd9\x33\x93\xe3\x7b\xb1\xcb\xb2\x88\x6c\xfa\xe1\x84\xc3\x69\x6c\xc0\x1b\x95\x71\x4a\x7f\xc5\x67\x83\x5e\x7d\x23\x5d\x63\xd1\x9b\xa6\x5f\x62\xbc\x4a\x9e\x53\x79\x84\x17\x76\x84\x60\xfa\x2d\x82\x1b\x76\x48\x42\xd8\x73\x9d\x16\x3b\x0d\xb2\xd8\xf3\x64\x4b\x08\x59\x99\xd1\x38\x44\xd7\xd6\xb9\x45\x34\xb5\xbd\x7b\x46\xf3\xc3\x49\x0d\xbe\x2f\xc6\x58\xaf\x0e\x44\xdf\x5d\x6c\xdb\x55\xf6\x9f\xf2\xd5\xf5\x42\x38\x36\x2b\xea\xff\x52\x61\xe6\xed\xe6\xc4\x08\xa6\x0f\x97\xbc\x5d\xb5\x71\x91\xb0\xc3\x4c\x9b\xab\x3e\x72\xac\x1f\xa0\x10\x73\x83\xc1\x74\x49\xc2\x54\x1c\x6e\xc9\x57\x00\x00\x00\xff\xff\x4d\xb6\x1f\x1f\x44\x02\x00\x00")

func _1_setupUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_setupUpSql,
		"1_setup.up.sql",
	)
}

func _1_setupUpSql() (*asset, error) {
	bytes, err := _1_setupUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_setup.up.sql", size: 580, mode: os.FileMode(0664), modTime: time.Unix(1715270769, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0xbe, 0x16, 0x81, 0x21, 0xb1, 0x90, 0xd0, 0xf2, 0x7c, 0x5c, 0x19, 0x62, 0xf4, 0x6a, 0xb7, 0xa4, 0x2, 0x54, 0xc8, 0xec, 0x2c, 0xdb, 0x3b, 0x11, 0xe9, 0x85, 0x6a, 0x73, 0x76, 0xe4, 0xf}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0664), modTime: time.Unix(1715177003, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
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
	"1_setup.down.sql": _1_setupDownSql,

	"1_setup.up.sql": _1_setupUpSql,

	"doc.go": docGo,
}

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
	"1_setup.down.sql": &bintree{_1_setupDownSql, map[string]*bintree{}},
	"1_setup.up.sql":   &bintree{_1_setupUpSql, map[string]*bintree{}},
	"doc.go":           &bintree{docGo, map[string]*bintree{}},
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
