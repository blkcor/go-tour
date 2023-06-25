// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x3f\x6f\xd3\x40\x14\xdf\xfd\x29\x9e\xc4\x5c\xc7\x76\x1a\x3b\xbd\x89\xd2\xa6\xa2\xa8\x85\x08\xbb\xea\x88\x2e\xf1\xab\xe3\xea\xec\xbb\xde\x3d\x37\x0e\x1b\x12\xea\x86\x18\x10\x03\x4c\x65\x61\x41\x02\x06\x24\x24\x24\xbe\x0d\x69\xf9\x18\xe8\xec\x34\x6e\x05\x6c\xa7\xdf\xfb\xdd\xbb\xdf\x9f\x8b\x51\x9f\xa3\x66\x0e\xc0\xd3\xaa\x3c\x94\x29\x32\x48\x71\x52\x65\x0e\xc0\x43\x22\x35\x96\x9a\x18\x0c\x3d\xcf\xb3\x0c\xe4\x69\x92\x17\x28\x2b\x62\x10\x5a\xe4\x58\xe7\x84\x77\xa1\x1d\x59\x12\xd6\x74\x1b\xdc\x56\xca\x3e\xb0\x8b\x27\xbc\x12\x34\xe6\x19\xc6\xf9\x73\x64\xe0\x5b\xfe\x21\xaf\x6f\x23\x16\x3a\x90\x59\xcc\xcf\x71\xcc\x69\xc6\xc0\x90\xd4\x3c\xc3\x9e\x90\x99\x69\x67\x7b\xb9\xc0\xc7\xbc\x40\x06\x5c\xa9\x0e\x1a\xd5\xc4\xc0\x15\xd2\x4a\x3f\x52\x42\xf2\xf4\xef\x25\x55\x83\x9b\x8e\xd1\xb8\x3f\xd2\x82\xc1\x8c\x48\xb1\x5e\xcf\x0f\x22\xd7\x73\x3d\xd7\x67\xd6\x74\xcf\x10\xa7\x7c\xba\xe6\xef\x17\x3c\xc3\x43\x5e\xb7\x6a\x07\x77\xf1\x6d\x21\xe4\x7c\x54\x93\xb1\x66\x01\x36\xc0\x3d\x55\x59\x77\xc4\xf5\x59\x95\x99\xe3\xec\x72\xe2\x13\x6e\xb0\x49\xe6\x41\xb2\x50\xc8\xa0\x58\x98\x33\x61\x97\x1a\xd4\x65\xe3\x50\x4b\x49\x00\xf7\x60\xf9\xe1\xd3\xf2\xe2\xdd\xaf\x9f\x97\xd7\xef\x5f\x5e\xbd\xfd\x7a\xf5\xea\xf3\xf2\xc7\x9b\xdf\xdf\x3e\x2e\x5f\x7f\x77\x00\xc6\xdc\x98\xb9\xd4\x29\x03\x3f\xe8\x6f\x0e\xc2\x68\xf8\xdf\x3b\xcb\x2f\x17\xd7\x97\x2f\x6c\xbb\xd2\x90\xe5\xdf\xd8\xed\xf7\xbd\xb0\x91\xd2\x46\x3b\x11\x32\x7b\x66\x50\x9f\xe7\x53\x74\x00\x12\x3e\x11\x38\xd6\x78\x92\xd7\xab\x99\xad\x7a\xc6\xb5\x41\x62\x50\xd1\xc9\xb0\x91\xa1\x4d\xf3\x1b\x18\x24\xba\xc2\xb6\xdc\xfd\x54\xe0\x8e\x2c\x4b\xd3\xf5\xfd\x44\x61\xb9\x82\xfa\x9e\xe3\x3c\x3a\x4e\x6c\x08\x31\x4e\xb5\x5d\x86\x69\xba\x98\x9e\x2e\x1c\x80\x7d\x63\x2a\xd4\xed\x7b\x1b\x9d\x96\x51\xad\x72\x8d\x0c\xa2\xc0\xf3\x1c\x67\x54\xf0\x5c\xb0\xb5\x21\x53\x90\x72\xcf\xce\xdc\xa9\x2c\xac\xa2\xe6\xfb\x6e\x86\x83\x55\xa8\xad\x37\xdf\x0b\xb6\xfc\x60\x6b\x18\x46\xf7\x3b\xe6\x3a\xc2\x63\x69\x66\x79\x56\xf1\x7c\x5e\xf9\x41\xa3\x22\x8e\x0f\x18\x50\xeb\x68\x4f\xcb\xe2\xdf\x1b\x12\x79\xd3\x7b\x30\x88\x42\x3f\xda\x1c\x0c\xfd\xd5\xf4\x4f\x00\x00\x00\xff\xff\xe6\x47\xff\xeb\x65\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 869, mode: os.FileMode(420), modTime: time.Unix(1687655531, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"configs/config.yaml": configsConfigYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
