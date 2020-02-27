// Code generated for package executor by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/.gitignore
// template/client/angular/readme.txt
// template/client/angular/static/readme.txt
// template/go.mod
// template/server/cmd/web/app.go
// template/server/cmd/web/main.go
// template/yoga.json
package executor

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

var _templateGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\x31\x6f\xdc\x30\x0c\x85\x77\xfd\x0a\x15\xd9\x0c\x48\x9e\xdb\xb1\x4d\x0b\x74\xea\x70\x45\xd7\x40\x96\xde\xd9\x4a\x24\x51\x10\xa9\xe4\xdc\xa1\xbf\xbd\xf0\xdd\x25\xc3\xf9\x16\x43\x78\xef\x23\x45\xd2\xe2\x60\x71\x82\x3a\x7f\xff\xa9\xc1\x86\x94\xd4\x60\x99\xb6\xe3\x9a\xe2\xa4\x06\x2b\x60\x51\x83\xa5\x2e\x4a\x3d\xe8\x03\xa0\x17\x91\xfa\x65\x1c\x17\xa4\x6a\xe7\x28\x4b\x9f\xac\xa7\x3c\xc6\xb9\x50\x83\x39\xc6\x04\x1e\xf5\x91\x9a\xce\xd4\xa0\xdd\x44\x5d\xf4\xd9\x8c\x65\xd6\x67\xdb\x6e\xa9\x3c\xe5\x1a\x13\x82\xa6\x2e\xb5\x8b\xf2\x29\xa2\xc8\xe8\xca\xdc\x93\x6b\x63\x88\xbc\xd3\x46\xc9\x75\xa7\x51\x17\x23\xec\xd5\x83\xfe\x55\xd2\xaa\x71\x8a\x2c\xac\xe3\x51\x7f\x75\x7f\x91\xf4\x9b\x63\xdd\x7a\xd9\x85\x4d\x9b\x6b\xae\x6d\x05\x54\x94\x80\xe2\x23\xf8\x96\x2c\x14\xf0\x94\x29\xf4\x04\xde\xd8\xda\xe8\x18\xd3\x47\x2f\xb7\xb8\x5f\x1a\x65\x98\x0b\x85\x66\xf0\x8a\x22\x3c\xd8\x67\xa6\x5d\x11\x5c\x81\x60\x32\x1c\xf7\x06\x53\x53\x9f\x63\xb9\x92\xea\x41\xff\x7c\xfc\xce\xda\x95\xa0\x11\xa2\x50\x63\x35\xda\x18\xe0\x94\xad\x8d\x9e\xe1\x45\x59\x9f\x1c\x73\x75\xb2\x28\xeb\x3f\x8f\x6a\xb0\xc9\xf5\xe2\x17\x65\x19\x22\xb1\xcc\xbc\x69\xdc\xa7\x14\x33\xcc\x1b\xb5\x17\xae\xce\xe3\x9a\x5b\x1b\xfd\xe7\xf0\x8d\x02\x94\x7d\x65\x4f\x01\xe3\xa0\x3e\xbd\x1f\xdf\x13\x5c\x8a\xf9\x90\xc5\xf1\xcb\xad\x76\xb9\xf3\x46\xc4\x49\x50\x38\x52\xb9\xd2\x76\x89\x2c\xd4\xd6\x71\xd8\x6e\xcf\x91\xfd\xed\x2c\x2c\x3b\x66\xe3\x9d\x5f\xb0\x1b\x29\x95\x02\x2f\x36\x91\x7f\xd9\x7b\xaf\x68\x6e\xde\xc5\xa4\x38\x55\xa0\x5d\x43\x23\x15\x9b\x68\xde\xfd\xda\x9a\x4d\xc0\xd4\xe7\x7b\xe6\xea\x5a\x31\x68\x8d\xda\x3d\x77\x5b\x0b\xe4\xbb\xce\x5a\xb7\xc9\x9d\xb7\x65\xdd\x20\xfd\xe3\xde\x3b\xb1\x8f\x87\xa7\x83\x50\xdb\x55\xfe\x7b\xe9\x79\x62\x1b\x26\xf5\x3f\x00\x00\xff\xff\x0a\xb5\xf1\x4d\x9d\x03\x00\x00")

func templateGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_templateGitignore,
		"template/.gitignore",
	)
}

func templateGitignore() (*asset, error) {
	bytes, err := templateGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/.gitignore", size: 925, mode: os.FileMode(420), modTime: time.Unix(1582776643, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateClientAngularReadmeTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xc9\xc8\x2c\x56\x48\xcb\xcf\x49\x49\x2d\x52\x28\xcf\xcc\xc9\x51\xc8\x48\x2c\x4b\x55\x48\xcc\x4b\x2f\xcd\x49\x2c\x52\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\x03\x04\x00\x00\xff\xff\x51\xe1\x8b\x7b\x29\x00\x00\x00")

func templateClientAngularReadmeTxtBytes() ([]byte, error) {
	return bindataRead(
		_templateClientAngularReadmeTxt,
		"template/client/angular/readme.txt",
	)
}

func templateClientAngularReadmeTxt() (*asset, error) {
	bytes, err := templateClientAngularReadmeTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client/angular/readme.txt", size: 41, mode: os.FileMode(420), modTime: time.Unix(1582776643, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateClientAngularStaticReadmeTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x81\x0d\x80\x30\x08\x04\xc0\x55\x7e\x28\x17\x40\x69\x85\x84\x14\xd2\xf2\xce\xef\x5d\xe6\x07\x33\x43\xc7\xc6\xb1\x64\x28\x4c\xbe\x81\x9b\x1e\x8a\x64\x17\x1b\x39\x21\xeb\x65\xc8\x86\x54\x85\x3f\xd2\x9e\xeb\x0f\x00\x00\xff\xff\x6c\xce\x12\x93\x3b\x00\x00\x00")

func templateClientAngularStaticReadmeTxtBytes() ([]byte, error) {
	return bindataRead(
		_templateClientAngularStaticReadmeTxt,
		"template/client/angular/static/readme.txt",
	)
}

func templateClientAngularStaticReadmeTxt() (*asset, error) {
	bytes, err := templateClientAngularStaticReadmeTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client/angular/static/readme.txt", size: 59, mode: os.FileMode(420), modTime: time.Unix(1582776643, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGoMod = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcd\x4f\x29\xcd\x49\x55\x48\x2c\x28\xe0\xe2\x4a\xcf\x57\x30\xd4\x33\x34\xe6\x02\x04\x00\x00\xff\xff\x38\x91\x92\x26\x14\x00\x00\x00")

func templateGoModBytes() ([]byte, error) {
	return bindataRead(
		_templateGoMod,
		"template/go.mod",
	)
}

func templateGoMod() (*asset, error) {
	bytes, err := templateGoModBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/go.mod", size: 20, mode: os.FileMode(420), modTime: time.Unix(1582776643, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServerCmdWebAppGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xd1\x6e\xdd\x36\x0c\x7d\xb6\xbe\x82\xd3\x93\x14\x38\x76\x81\xbe\x5d\x20\x0f\x41\x92\x2e\x03\xd6\xb4\x68\x82\x15\x43\x51\x0c\x8a\xcd\xd8\x5a\x65\xc9\xa3\xe4\x18\x41\x90\x7f\x1f\x28\xf9\x26\x77\xcb\x30\xb4\x2f\xd7\xb6\x28\x1e\x1e\x1e\x1e\xde\xd9\x74\xdf\xcc\x80\x30\x19\xeb\x85\xb0\xd3\x1c\x28\x81\x12\x95\x44\xdf\x85\xde\xfa\xa1\xfd\x33\x06\x2f\x45\x25\x5d\x18\xf8\xe1\x31\xb5\x63\x4a\xb3\x14\x5a\x88\xb6\x85\xdf\xc3\x60\x72\x36\x98\x79\x76\xb6\x33\xc9\x06\x2f\xd2\xc3\x8c\x25\x14\x13\x2d\x5d\x82\x47\x51\x8d\xc6\xf7\x0e\x29\xc2\x64\xe6\x2f\x31\x91\xf5\xc3\x57\x46\x6a\x2e\x4b\xe0\xdd\xe2\x3b\xf1\x94\x41\xcf\x4d\x32\xd0\x2f\xd3\xf4\x00\x3d\xbf\x16\x90\x82\x7a\xfe\x72\xc0\xa8\xbf\x19\xb7\x20\x14\xb8\x92\x7d\x85\x6b\xae\x7c\x46\x68\x12\x82\x01\x8f\x2b\x93\x13\x77\x8b\xef\x60\x8b\xaa\x2e\x50\x84\xdb\x10\x9c\x2e\x3c\x1f\x45\xf5\xc0\xcf\xdd\x49\xfe\x7e\x14\xd5\x33\xe1\x1d\x4c\xe6\x1b\xaa\xff\xa1\xad\x6b\x51\x3d\x89\x8a\xb9\x6e\xa7\x8c\xd3\xdb\x68\x6e\x1d\x9e\x05\x8a\x8a\xb1\x9b\x9f\x31\x31\x7b\x5d\x4a\x35\x7b\xfc\x2f\xb2\x35\xb3\x6d\x39\x5b\x7e\x85\x13\x38\x80\x79\x7d\x33\xdf\xc8\xf5\xdf\x59\x87\xd7\x48\xf7\x48\x2a\x7f\x9f\x5b\x52\xb2\x73\x16\x7d\x6a\x8d\x1f\x16\x67\xa8\x8d\xc9\x24\xdb\xb5\x52\xeb\x26\x5f\xbd\xbc\xb9\xf9\x28\x2a\xc2\xb4\x90\x07\x86\x2e\x92\xe5\x18\x94\xdf\x34\x22\xc4\x0c\x5b\x04\x53\x06\x8e\x58\x11\x5d\xe2\x4a\x03\x12\x05\x62\xc1\xee\x02\xc1\x6c\xd2\x58\xc3\xf8\xd2\x35\x19\x3f\x20\xbc\x70\x86\x2c\xe5\x8b\x60\xea\x1f\x19\x3a\x0b\xe7\xc2\xd0\x7c\x24\xeb\x93\xf3\x4a\x7e\xc6\xdb\x8d\x00\xd8\x08\xe6\xde\x58\xc7\x32\x42\xf0\x90\xdd\xf9\xf6\xcd\xdb\x37\x52\x3f\x77\x91\xa1\x7f\xb5\x31\xa1\x3f\xf5\x7d\xe1\x28\x77\xf9\x52\x0d\xde\x3a\xbd\x79\x6a\x13\x1f\xd0\xf7\x73\xb0\x3e\xfd\xbb\xbb\x2d\xae\xd6\x82\xf8\x09\xe3\x1c\x7c\xc4\xcf\x64\x13\x52\x0d\x04\x47\xdb\xf9\x5f\x0b\xc6\xa4\xb9\xad\xb5\xb9\x44\xd3\x23\x29\x56\x37\x29\x79\x16\x7c\x42\x9f\x8e\x6f\x1e\x66\x94\x35\xc8\x83\x8d\x28\x6b\xa4\x45\x31\x09\xeb\xc4\xc5\x58\x9a\x6c\xe0\x1d\xc8\x4b\x74\x2e\xc8\x62\x24\xa4\x2c\x25\xe7\x34\x57\xb8\x5e\xf0\x2e\x22\xa9\x55\x37\xe5\x55\xf5\xc5\x46\xf6\x8e\x87\x01\x3f\x9d\x70\xa3\x59\xe8\x88\xbe\xbf\x20\x52\x6b\x5d\xda\xb8\x4e\x26\x2d\xf1\x17\x9f\x90\xbc\x71\xc5\x2e\x17\x3c\xbf\x9a\x33\x9b\xfc\xaa\x74\x9e\xc2\x93\x28\x92\x3c\x43\xfc\xb7\x10\x5c\x1f\xac\x4f\x35\x4c\x18\x23\xff\x77\x94\xa5\xc8\x8a\x10\xc6\xb9\x86\x3f\x9e\xc9\xbf\x37\x14\x47\xe3\x0e\x77\xa7\x3c\x1e\x65\x76\x91\xdc\xed\x51\x9e\xb4\x28\x2e\x29\x94\xd6\x7a\x83\x55\x0c\xa9\x4b\xd9\xfd\x28\xaf\x10\x7b\xec\xc1\x7a\x08\xd4\x23\x41\x0a\xfb\x5d\x83\xb3\x0f\x9f\xae\x81\x7d\xe9\x42\x67\x1c\xf4\x78\x8f\x2e\xcc\x13\xee\xe7\x7d\xb8\x93\x23\xbc\xda\xe3\x57\x27\xa5\xa9\x6c\x34\xce\xff\x11\x77\xbc\xb2\xc7\x69\xd7\x61\x8c\xc7\xec\x12\x0a\xee\xf8\xd4\xb9\xb0\x1e\x7f\x20\x3b\x58\xcf\x6e\x39\x62\x53\x7f\x5f\xd2\x7b\x4c\x63\xe8\xe3\x0f\x66\x95\x2b\x07\x59\x23\xcb\x4c\xdb\xf0\xff\x0e\x00\x00\xff\xff\x87\x39\x4d\x63\x0a\x06\x00\x00")

func templateServerCmdWebAppGoBytes() ([]byte, error) {
	return bindataRead(
		_templateServerCmdWebAppGo,
		"template/server/cmd/web/app.go",
	)
}

func templateServerCmdWebAppGo() (*asset, error) {
	bytes, err := templateServerCmdWebAppGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/server/cmd/web/app.go", size: 1546, mode: os.FileMode(420), modTime: time.Unix(1583001232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServerCmdWebMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc9\xb1\x0a\x02\x31\x0c\x06\xe0\xb9\x79\x8a\xd0\xa9\x05\xb9\x07\x10\x1c\x5d\x45\x70\x72\x0c\x47\x2c\xc5\x5e\x13\x7e\x5a\x1d\xc4\x77\x17\x6f\xfd\x3e\x97\xf5\x29\x45\x79\x93\xda\x89\xea\xe6\x86\xc1\xb1\x59\x89\x44\x8f\xd9\xd7\x3d\x52\xe6\x0f\x05\x71\xe7\xe3\x89\x2f\xfa\xbe\x5b\x91\x34\x30\x35\x53\x50\xe0\xaf\xe2\xbe\xdc\x14\x2f\x4d\x99\x42\xb3\xb2\x5c\x51\xfb\x68\x3d\xc5\x33\x60\x88\x07\x56\x20\xd3\x97\x7e\x01\x00\x00\xff\xff\xf9\x16\x97\x06\x71\x00\x00\x00")

func templateServerCmdWebMainGoBytes() ([]byte, error) {
	return bindataRead(
		_templateServerCmdWebMainGo,
		"template/server/cmd/web/main.go",
	)
}

func templateServerCmdWebMainGo() (*asset, error) {
	bytes, err := templateServerCmdWebMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/server/cmd/web/main.go", size: 113, mode: os.FileMode(420), modTime: time.Unix(1582998841, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateYogaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\x2a\x4b\x2d\x2a\xce\xcc\xcf\x53\xb2\x52\x50\x32\xd4\x33\xd0\x33\x50\xe2\xaa\x05\x04\x00\x00\xff\xff\xbc\x36\xa3\xb2\x1a\x00\x00\x00")

func templateYogaJsonBytes() ([]byte, error) {
	return bindataRead(
		_templateYogaJson,
		"template/yoga.json",
	)
}

func templateYogaJson() (*asset, error) {
	bytes, err := templateYogaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/yoga.json", size: 26, mode: os.FileMode(420), modTime: time.Unix(1582776643, 0)}
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
	"template/.gitignore":                       templateGitignore,
	"template/client/angular/readme.txt":        templateClientAngularReadmeTxt,
	"template/client/angular/static/readme.txt": templateClientAngularStaticReadmeTxt,
	"template/go.mod":                           templateGoMod,
	"template/server/cmd/web/app.go":            templateServerCmdWebAppGo,
	"template/server/cmd/web/main.go":           templateServerCmdWebMainGo,
	"template/yoga.json":                        templateYogaJson,
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
	"template": &bintree{nil, map[string]*bintree{
		".gitignore": &bintree{templateGitignore, map[string]*bintree{}},
		"client": &bintree{nil, map[string]*bintree{
			"angular": &bintree{nil, map[string]*bintree{
				"readme.txt": &bintree{templateClientAngularReadmeTxt, map[string]*bintree{}},
				"static": &bintree{nil, map[string]*bintree{
					"readme.txt": &bintree{templateClientAngularStaticReadmeTxt, map[string]*bintree{}},
				}},
			}},
		}},
		"go.mod": &bintree{templateGoMod, map[string]*bintree{}},
		"server": &bintree{nil, map[string]*bintree{
			"cmd": &bintree{nil, map[string]*bintree{
				"web": &bintree{nil, map[string]*bintree{
					"app.go":  &bintree{templateServerCmdWebAppGo, map[string]*bintree{}},
					"main.go": &bintree{templateServerCmdWebMainGo, map[string]*bintree{}},
				}},
			}},
		}},
		"yoga.json": &bintree{templateYogaJson, map[string]*bintree{}},
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
