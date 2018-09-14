// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdf\x73\xdb\xb8\xf1\x7f\xe7\x5f\xb1\xa3\x97\xaf\xf3\x9d\x58\xb4\x7d\x77\x6d\xc6\xaa\x3b\x75\xe5\xf4\xac\x49\xfc\x63\x22\x5f\x32\xf7\xa4\x81\xc0\x15\x89\x1a\x04\x78\x00\x28\x45\xbd\xf1\xff\xde\x59\x00\x14\x49\x89\x4a\xd2\xf4\x1e\xaa\x07\x5b\xc2\xfe\xc0\xee\x7e\xf6\x17\x99\xa6\x30\xd5\xd5\xd6\x88\xbc\x70\x70\x71\x76\xfe\x06\xe6\xac\xb4\xb5\xca\x61\x7e\x33\x87\xa9\xd4\x75\x06\xf7\xcc\x89\x35\xc2\x54\x97\x55\xed\x84\xca\xe1\x09\x59\x09\xac\x76\x85\x36\x76\x9c\xa4\x69\x92\xa6\xf0\x5e\x70\x54\x16\x33\xa8\x55\x86\x06\x5c\x81\x70\x5d\x31\x5e\x60\x43\x79\x0d\x1f\xd1\x58\xa1\x15\x5c\x8c\xcf\xe0\x84\x18\x46\x91\x34\x7a\x35\x21\x15\x5b\x5d\x43\xc9\xb6\xa0\xb4\x83\xda\x22\xb8\x42\x58\x58\x09\x89\x80\x9f\x39\x56\x0e\x84\x02\xae\xcb\x4a\x0a\xa6\x38\xc2\x46\xb8\xc2\xdf\x13\xb5\x90\x25\xf0\x6b\xd4\xa1\x97\x8e\x09\x05\x0c\xb8\xae\xb6\xa0\x57\x5d\x46\x60\x2e\x1a\x4d\x9f\xc2\xb9\xea\x32\x4d\x37\x9b\xcd\x98\x79\x83\xc7\xda\xe4\xa9\x0c\xac\x36\x7d\x3f\x9b\xbe\xbd\x9f\xbf\x3d\xbd\x18\x9f\x45\xa1\x5f\x94\x44\x6b\xc1\xe0\x6f\xb5\x30\x98\xc1\x72\x0b\xac\xaa\xa4\xe0\x6c\x29\x11\x24\xdb\x80\x36\xc0\x72\x83\x98\x81\xd3\x64\xf4\xc6\x08\x8a\xdb\x6b\xb0\x7a\xe5\x36\xcc\x20\xa9\xc9\x84\x75\x46\x2c\x6b\xd7\x8b\x59\x63\xa2\xb0\x3d\x06\xad\x80\x29\x18\x5d\xcf\x61\x36\x1f\xc1\xdf\xaf\xe7\xb3\xf9\x6b\x52\xf2\x69\xf6\x74\xfb\xf0\xcb\x13\x7c\xba\xfe\xf0\xe1\xfa\xfe\x69\xf6\x76\x0e\x0f\x1f\x60\xfa\x70\x7f\x33\x7b\x9a\x3d\xdc\xcf\xe1\xe1\x1f\x70\x7d\xff\x2b\xbc\x9b\xdd\xdf\xbc\x06\x14\xae\x40\x03\xf8\xb9\x32\xe4\x81\x36\x20\x28\x9a\x98\xf9\xd0\xcd\x11\x7b\x26\xac\x74\x30\xc9\x56\xc8\xc5\x4a\x70\x90\x4c\xe5\x35\xcb\x11\x72\xbd\x46\xa3\x28\x13\x2a\x34\xa5\xb0\x84\xaa\x05\xa6\x32\x52\x23\x45\x29\x1c\x73\xfe\xe8\xc0\xaf\x71\x42\x2c\x77\x82\x17\x0c\x25\x7c\x44\x85\xff\x12\x0c\xfe\x52\xae\xc3\xb7\xbf\xe5\x25\x13\x72\xcc\x75\xf9\xd7\x24\xb1\x5b\xe5\xd8\x67\xb8\x82\x51\x65\xb4\xd3\x3f\x8c\x26\x49\x52\x31\xfe\x4c\x16\xf0\x92\xad\x4b\x8a\xe4\x24\x49\x44\x59\x69\xe3\x60\x94\x6b\x9d\x4b\x4c\x59\x25\x52\xa6\x94\x8e\x36\x8c\xbd\xf0\x68\xb2\x63\xf3\xbf\xf9\x69\x8e\xea\xd4\x6e\x58\x9e\xa3\x49\x75\xe5\x59\x07\xc5\x92\x24\x50\xe1\x24\x37\x15\x1f\xe7\xcc\xe1\x86\x6d\x03\x99\x2f\x72\x54\x8b\xa8\x65\x1c\xb5\x8c\x75\x85\x8a\x55\x62\x7d\xd1\x50\x5e\xc1\x15\xfc\x9e\x00\x08\xb5\xd2\x97\xfe\x1b\x80\x13\x4e\xe2\x25\x8c\xa6\xb2\xb6\x0e\x0d\xdc\x31\xc5\x72\x34\xf0\xf1\xee\x13\x33\x08\xb7\x28\x2b\x34\x70\xfd\x38\x1b\x4d\x3c\xff\x3a\xd4\xce\x25\x8c\xd6\x67\xe3\xf3\xf1\x59\x3c\xe6\x5a\x39\xc6\x5d\xa3\x95\x3e\x8a\x95\xa4\x78\x2f\xc6\x91\x9f\x3e\xb5\x91\x97\x30\xa2\xb4\xb7\x97\x69\x9a\x0b\x27\xd9\x92\x42\x9e\x36\x28\xa4\xbc\x64\xa7\x21\xbc\x1d\x31\x24\x68\x2e\x61\x74\x88\x55\x64\x7a\xa1\x7f\xfe\x0f\x7e\x76\x68\x14\x93\x8b\x4c\x73\xdb\xd8\xf6\x1d\xd7\x66\x68\xb9\x11\x3e\xac\xe4\x91\x36\x08\x6c\xa9\x6b\x07\xdf\x16\xb5\x97\x04\xc0\xf2\x02\x4b\xb4\x97\x70\xfb\xf4\xf4\x38\x9f\xec\x9f\xd0\x01\xd7\xca\xd6\xfe\x64\x14\xeb\x98\x2e\x4c\xff\x69\xb5\xf2\x6a\x2a\xa3\xb3\x9a\x1f\xa3\xbf\x4c\x92\xc4\xa2\x59\x0b\x8e\x3b\xb3\x82\xc3\x54\x9e\x42\x4a\x92\x5f\x0b\xdf\xf8\x18\xf0\xc0\xe1\xe9\xa6\xe2\x30\x35\xc8\x1c\x36\x72\x27\xbd\x9f\x77\x36\x7f\x05\x06\x5d\x6d\x94\xdd\x23\x7d\xc0\x4a\x6e\x5f\x75\x40\xdf\xe5\xa8\xaf\x81\x31\xab\xc4\x98\x22\xdd\x64\x5e\xfb\xa9\x6a\x07\x97\x30\xf2\x55\xb2\x3e\x4f\xa3\x3d\xa3\x1e\xcf\x52\x67\x5b\x62\xfa\xff\xf6\xf8\x25\x42\xdc\x73\xcc\xa0\x33\x02\xd7\xa1\x6b\x58\xc7\x5c\x6d\xa9\xd3\xee\xbc\xa4\x8e\x00\xc2\x59\x78\xae\x97\xc8\xb5\x5a\x89\xdc\x37\x15\xae\x95\x42\xee\xc4\x5a\xb8\xed\x2e\x12\x3f\xa3\xdb\x85\xa1\xfd\xde\x8f\x41\x7b\xfe\xfd\x01\xc8\xf1\xcb\x01\x18\xf4\x34\x43\x89\x0e\x07\xf0\xbb\xf1\x84\x9d\xe1\xbd\x9f\x7d\xdb\x7b\xa4\xef\x37\x3f\x5a\xf2\x1f\x7b\xb0\xc3\x8a\x81\x14\xd6\x11\x4e\x51\xd0\x0e\x40\xf0\x9e\x58\x4e\xfa\xbf\x8f\x41\x41\xb4\x3f\x1a\x8e\x94\x6c\xfc\xba\x47\xb5\x51\x4d\x63\xf4\xcd\xd5\x94\xbe\x34\x63\x93\x60\x95\x00\xaa\xcc\x0e\x5c\x3f\xa3\x8b\x4b\xc8\xac\xc3\x7e\xd2\x1e\x1f\x38\x19\xcf\xff\x30\x07\xa3\xb9\x03\xbe\xbd\x24\x49\x89\xd6\xd2\x74\xdb\x6f\x03\x6d\x43\xb9\x67\x25\x36\xdb\x4c\x53\x65\x4e\xc3\x12\xdb\x2e\x83\x99\x67\xa6\xdd\x41\xe5\x7e\x1a\xc0\x15\x9c\x4f\x1a\x0d\x4f\x45\xe4\xa5\xc9\xdc\x8c\x76\x1f\x07\xcf\xd1\xbb\xfa\x31\xf2\xcd\x2b\xe4\xad\xd0\x15\x5c\x4c\x8e\x5a\xeb\x03\xd5\x69\x80\x05\xfa\x95\x43\x1b\xbf\xd5\x75\xcd\xde\x30\xdb\x35\x9a\xd6\x28\xbf\xf0\xd1\x5e\x85\xd6\x25\xa1\x13\x69\x09\xfa\xf9\xc0\x81\x0c\x1d\x13\xd2\xee\x47\x22\x8a\x82\x41\x5b\x69\x65\x31\x78\x14\x88\x33\x87\xe5\x8e\x71\xdf\x85\x5e\xc3\xf9\x96\x68\x4b\xad\x9f\x69\x6f\xab\x86\x63\x3d\xa8\x7a\x2f\x34\x33\xdb\xd3\x2b\x54\x68\xa3\x5b\xeb\xb0\x3c\x74\xbe\xeb\xca\x8d\xf7\xfe\x8b\x0e\xed\x37\xa2\x2e\x22\xcc\xd1\x76\xd9\xb9\xfb\xff\x6c\x30\xdd\x69\x9a\xb8\xce\xe8\xed\x57\xbd\x3a\xec\x66\xed\x0d\x53\x5d\xcb\xac\xe7\xdb\x12\x1b\xc5\x31\x39\x87\x70\x9d\xef\x06\x08\x89\x76\xb3\x20\x1a\x12\x27\xcc\x71\xec\x62\x97\x82\xdf\x8f\x93\xff\x2b\x0c\xa2\xd0\xfb\xc1\xfe\x89\x15\x55\x41\x36\x94\x6e\x87\x36\x77\x99\x5a\x63\x6e\xf6\x72\xad\xeb\xbc\xc8\x7a\x36\x0c\x64\xe6\x00\x66\x17\x93\x21\xd4\x6d\x2f\xd0\x03\xd2\xbb\x40\xff\x30\x64\x74\x27\xfb\xfe\xb7\x4d\x1f\x90\xef\x2c\x22\x4e\x37\x7b\x08\x7d\x3d\xa2\xae\xc3\x7f\x05\x3f\x1e\xef\x7a\xbd\x46\x39\x58\x6a\xbb\xee\x79\x0a\x52\x3c\x23\x84\x3d\xf7\xab\x9d\xba\x99\x6e\x7a\x05\xef\xea\x25\x1a\x85\x0e\xad\x5f\xa2\x36\xda\x3c\x23\x75\xd5\x0c\xed\x18\xa6\x5a\x39\xa3\x25\x54\x92\xa9\x9d\x94\x05\x5a\x8a\x33\x74\xf4\x90\xb6\x6b\xb1\x08\x77\x8c\x17\x42\x21\x19\x3b\xee\x79\xfb\xc6\x2e\x9a\x0b\x3d\x02\x5d\x4b\xe2\x8a\xfd\x95\x89\x11\xb8\x7c\x18\x82\x87\x01\x8a\x6f\x18\x07\xc2\xc2\xed\x75\x5b\x73\x85\xc8\x8b\x05\x5b\x33\x21\xd9\x52\x48\xe1\xb6\x01\x80\x8e\x41\x2b\xb6\x34\x82\xc7\x7e\x5c\xdb\xbd\xb1\x87\x8e\x02\xb4\x88\x4c\x57\xf0\xd3\x24\x21\xf4\xa2\x2c\x37\x98\xa1\x72\x82\x49\x4b\x0a\xea\xf8\xb4\xcb\xc9\x15\x12\xef\xa6\x43\x03\x78\x27\x6a\x2d\xc4\xa4\xac\xb6\xf4\xb0\x53\x06\x15\xf3\xf9\x2d\x30\xce\xd1\xda\xae\x39\x3b\x96\x7d\x7c\x0b\x6d\xdd\x17\xe4\x3c\xb9\x5b\x0c\x7e\x7a\xd3\x63\xec\x71\x19\x4f\xee\x16\x40\x98\xf8\x62\xcd\x1c\xc2\x33\x6e\x1b\xd1\x31\x3c\xd1\xb8\x2d\x6b\xeb\xc2\xfe\x10\x33\xbd\x36\xe1\x7d\x83\x6b\x41\x17\xca\x3a\xa6\x38\xf6\xef\x09\x2a\x17\xa4\x72\x1f\x9b\xe7\x37\x76\x97\xba\xcd\x6b\x04\xde\x4d\xd1\x78\x39\x65\x2f\x21\xaf\x95\xdc\x02\x83\x92\x85\x54\x58\x35\xaf\x7e\x50\x66\xfe\x4d\x08\xae\x28\x7f\x7b\xc9\x1a\xd5\x2d\xbc\xba\x4e\xda\xfe\x74\xbc\x4a\x3b\xc9\xd9\x01\x50\x84\xe9\x67\x2b\xc6\x11\x98\xd4\x2a\x6f\xdf\x2c\xc5\x1c\xf0\xd0\xf9\xc3\xe3\x55\x5f\x2b\xf1\x5b\x8d\x72\x0b\xc2\x67\xd6\x2a\xf8\x43\x8f\xa7\xd9\xb1\x36\x17\xae\xec\xa6\x44\xcc\x31\x0b\x9b\x42\xf0\xc2\xbf\xec\x32\xc2\xe2\x41\x7b\xda\x4d\x9b\x6e\x52\x96\x8d\x70\x33\x6d\xd2\x94\xa6\x5f\xb7\x87\x5c\x3f\xce\x60\x1e\x56\xe2\xce\x7c\x6c\x77\xdf\x30\x3a\xd3\x14\xc2\x9c\x24\xe8\x1a\xe9\x66\x20\x1f\xca\xed\xcf\xd4\x15\xe8\x0a\x4d\x58\xad\x69\xc9\x7b\x78\x77\x64\x9d\x69\x54\x0d\xac\xe4\xed\x0e\x1d\x33\xca\xb1\xbc\x49\xca\x5c\xd0\x86\x57\x69\x2b\x9c\x36\xdb\x1d\x63\x0c\x6c\x2e\x5c\x27\x1b\xce\x27\xfb\x8a\x0a\x66\x8b\x66\x80\x90\x26\xae\xcb\x52\xb8\x21\x2d\x81\xd2\x96\xdf\xf1\x1e\xe6\x0c\xa2\x77\x95\x4b\x64\x0a\x36\x05\x2a\x58\xd6\x42\x0e\xaa\x25\xe6\x05\xcd\x29\x6c\xab\x34\xaa\xbe\xa1\x43\xbd\xf2\xb2\xd9\xbe\xac\x3f\x5c\x64\x41\xee\xc7\x9e\xdc\xc7\x16\xe1\xdc\x77\xb3\x2c\x8c\xb7\xb2\x12\x12\x0f\x6c\xd0\xfd\x6a\xe9\xe8\x99\x06\x09\xd3\xb6\xd3\x5e\xc1\x45\xe2\x15\xfc\xa9\x27\xf5\x28\x99\x23\xe4\x40\xb8\x10\x84\xc0\x98\xf9\xf4\x49\xc1\xd4\xca\xbf\x2a\x8c\xe3\xa2\xdb\x41\x1a\xc1\x2b\xf8\x73\xf3\x14\x94\xec\xb9\xd4\x49\x0a\x4f\x1a\xc8\x95\xe8\xcd\xa2\xfb\x04\xd8\x94\x40\xf2\xef\x00\x00\x00\xff\xff\x8b\x52\x39\x73\xe0\x16\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 5856, mode: os.FileMode(420), modTime: time.Unix(1536859088, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
