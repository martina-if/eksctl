// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/schema.json (41.487kB)

package v1alpha5

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x7d\x6b\x73\xdc\xb6\x92\xf6\x77\xff\x0a\xd4\xf8\xd4\x49\x5c\x25\x6a\x4e\x52\xf5\xbe\xbb\xe5\x64\x55\x3b\x96\x7c\xd1\xda\x96\xb4\x1a\x27\xa9\x8d\xe4\x2a\x63\xc8\x1e\x0e\x22\x12\x60\x00\x50\xf2\x38\xf1\x7f\xdf\x02\x08\x72\x78\x01\x2f\xe0\x50\xb6\xbc\x9f\x6c\x71\x88\x46\xa3\xbb\xf1\xa0\xd1\x68\x34\xff\x7a\x84\xd0\xec\x1f\x1c\xd6\xb3\xa7\x68\xf6\x78\x1e\xc0\x9a\x50\x22\x09\xa3\x62\x7e\x1c\xa5\x42\x02\x3f\x66\x74\x4d\xc2\xd9\x81\x7a\x51\x6e\x13\x50\x2f\xb2\xd5\x1f\xe0\xcb\xec\xd9\x3f\x84\xbf\x81\x18\xab\xc7\x1b\x29\x93\xa7\xf3\xf9\x1f\x82\x51\x2f\x7b\xea\x31\x1e\xce\x03\x8e\xd7\xd2\xfb\xd7\xbf\xcd\xb3\x67\x8f\xb3\x76\xa5\xae\x66\x4f\x91\xe2\x03\xa1\x59\xde\x67\xc4\xd2\xe0\x37\x2c\xfd\x4d\xf1\x13\x42\xb3\x84\xb3\x04\xb8\x24\x20\x4a\x4f\x11\x9a\xf9\x59\xa3\x37\x2c\x0c\x09\x0d\x2b\xbf\xf5\x0e\xae\xe8\x28\x6f\x5d\x34\xfd\x6c\xfe\xf7\xf9\x60\xd7\x3f\xac\x81\x73\x08\xce\x79\x00\x7c\xf6\x14\x5d\xb5\xf2\x60\x7e\x78\x5f\xb4\xc5\x41\xa0\x7b\xc6\xd1\x45\x79\x14\x6b\x1c\x09\x28\x5e\x0a\x40\xf8\x9c\x24\xea\x3d\xc5\xb1\xcf\xa8\xc4\x84\x0a\xe4\x6b\x15\xa0\x04\x73\x1c\x83\x04\x2e\x10\x87\x08\x4b\x08\x90\x64\xa8\x24\xab\x82\xd0\x47\x8f\x50\x09\x51\x44\xfe\xf0\x36\x32\x8e\xbc\x7d\x09\x3f\x2a\x09\xa2\xa9\xa3\xa6\xe0\x5b\x55\x05\x14\xaf\x22\x78\xb7\x4d\x6a\x3f\x20\x34\x23\x12\xe2\xfa\xc3\x92\xc9\x09\xc9\x55\x1f\x07\xd5\x5f\x03\x58\xe3\x34\x92\xea\x85\x59\xe9\x97\xcf\xe5\xd7\x0a\x12\x98\x73\xbc\x1d\xab\xe0\x32\xe7\x53\x6a\x17\x78\xb7\x16\x8c\x61\xa1\xc8\x08\xd9\x51\xc7\xae\xe4\xad\x9a\xce\x10\xa0\xa4\x5e\x0e\x7f\xa6\x84\x43\x50\x15\x51\x0c\x12\x07\x58\xe2\xa6\x7c\xda\xcc\x01\x27\xe4\x57\xe0\x22\x63\xf9\x2f\x9b\xce\x2c\x6a\xaf\x28\xbd\xf2\x03\x7c\xc4\x71\x12\xe9\x4e\xae\xaa\x76\x02\x37\xc2\x97\xd1\x21\x61\xf3\xdb\x1f\x70\x94\x6c\xf0\xff\x2b\x9b\xcb\xfb\x47\x16\xc3\x99\xe1\x5b\x4c\x22\xbc\x22\x11\x91\xdb\xdf\x19\x7d\x08\x16\x7b\x50\x06\x1c\x0b\x44\x22\x07\xc0\xb3\x93\x5d\x63\x1e\x62\x09\x17\x9c\xad\x49\x34\x78\xc8\xf6\x2e\x5f\x54\x68\xed\x35\xda\x90\xc8\x61\xc3\x7c\x49\xa4\x9d\x02\xc1\xb1\x93\xa0\x4e\x17\x6f\xed\x84\x6e\x08\x0d\xee\xd9\x5a\xab\xf3\xae\xd7\x50\x63\x4c\x71\x08\xc1\x19\x0b\xe0\x25\x67\x69\xb2\x9f\xd6\xde\xd6\xa8\x0d\xd5\x5b\x6d\xd0\x15\x30\x5a\x02\xa0\x2b\x45\x31\xd4\xfc\xa1\x54\xe0\x10\xde\x7f\x3f\xd7\xff\xce\x35\xff\x84\x86\x1e\x2d\xde\x78\x82\x30\x0d\xd0\x95\x19\x19\xda\xfd\x50\x34\x82\x1b\xe1\x99\x9f\x75\x3b\x31\x7f\x52\x65\xa1\x07\x1f\x15\x4b\x3f\x63\xb4\xe1\xb0\xfe\x8f\xeb\x59\x3b\x27\xd7\xb3\xa3\x3a\xe3\x3f\xcf\xf1\x91\xe6\xaf\xd1\xbe\xc9\xd4\xf5\xec\xa8\x39\x08\x45\xc0\x6e\x5c\x05\x84\xba\x98\xea\x5b\xd8\x61\x6e\x95\x1c\x9d\xc6\x24\x26\xb5\x85\x17\x8c\x23\x42\xd7\x8c\xc7\x58\x3d\xd2\x82\xcc\xa7\x02\x12\xca\x50\x2c\xda\xb6\x99\x88\x93\xba\x7b\x7b\x1d\x68\x0b\x43\x94\x28\xc0\xe7\x20\xc5\x73\xea\xf3\x6d\xce\xc0\x00\x6d\x2e\x1b\xcd\xec\xd4\x25\x96\x69\x43\x9f\x9d\x06\xb2\xcc\x9a\x58\xc9\xdd\x26\xbe\x13\xad\x5f\x2f\x8e\xc7\xba\x50\x1a\x39\x0f\xac\xab\xbf\x6d\x0a\xd4\xc0\xbb\xc6\xb3\xdd\xca\x3b\x51\xb1\x63\x99\xeb\x5c\xf6\xed\xab\x6e\xa7\xc2\x9b\xfa\xaa\x2d\x66\x93\xb8\x90\x18\x09\xa2\x2c\xd8\xb8\x78\x07\xca\xa1\x5b\x01\xe2\x90\x44\xd8\x87\x00\xdd\x11\xb9\x41\x46\x6f\x68\x71\x71\x3a\xd8\x79\x74\x26\x6c\x73\x1b\x9f\xd3\x20\x61\x84\x4a\x31\x64\x63\x90\x70\x72\x8b\x25\x2c\x7c\x1f\x44\xc3\xb8\x73\x70\x59\x31\x16\x01\x6e\x99\x17\x49\xba\x8a\x88\xef\x4a\xc0\xc9\x80\xab\x4c\xb6\xf5\x3d\x89\x6a\x37\x2c\x0a\x44\xe1\x9d\xe3\x84\x20\x01\xfc\x16\x38\x02\x23\x55\x84\x75\x6f\x65\x50\x1b\xac\xde\x51\xc4\x6d\x2a\x56\x4e\xd2\x00\xe5\xe6\xb3\x8d\x05\xcf\x3f\x82\x9f\x2a\x72\x97\x2c\x82\xc5\xe5\x59\x8f\x23\xd5\xe9\xa2\xd6\xa8\x5d\x00\x8f\x89\x50\x68\x22\x9e\xb1\x94\x06\x98\x6f\xc7\x50\x57\x92\x20\xbe\xd2\x31\x4b\xab\xb6\x8b\x9c\xd7\xcd\x9d\x94\x96\x15\xaa\x7b\xb9\xc2\x86\xc1\x3d\x04\x58\xa2\x30\x91\xd0\x14\x1e\x9c\x9f\x9e\x1c\xdf\xd3\xbc\xab\x0d\x79\xf8\x50\xfa\xad\xa6\x46\xcf\xc1\xb6\x6c\xc3\xef\xb0\xa3\x09\x51\x01\x47\x11\x3a\x5d\xbc\x45\x58\x4a\x4e\x56\xa9\x04\x81\xd8\x1a\xe1\x7c\x42\x3b\xc2\x40\x1f\xb5\x96\x79\x5f\xb3\xe8\x01\x28\x80\x29\x65\x12\x57\x03\x7e\xdd\xb2\xb8\xbf\x8d\x75\x29\x80\x69\x23\xf0\xd7\x67\xbb\x9d\x63\x29\xb1\xbf\xb9\x60\x11\xf1\x1b\xf3\xc4\x0e\x01\xa7\x34\x22\x14\x4e\x98\x9f\xc6\x40\x1b\x1d\xda\xd4\x81\x12\x4d\x1e\x05\xa6\x8d\x5a\x7b\xb3\x7e\xd5\xff\xe4\x86\x08\x64\x6c\x4b\xa1\xb4\x16\xbe\x8b\x23\x3c\xbe\x97\x5e\x89\x2c\x2e\xcf\x1e\x56\x8c\x24\xc2\x2b\x88\xbe\x59\x63\xa3\x38\x86\xb1\x81\x86\x56\x82\x22\xc1\xfe\xb4\x54\x13\xe7\x25\xc4\x8d\xfe\x88\xfd\x4e\x03\x9e\xba\x36\x40\x12\x87\xdf\x96\x89\x38\x2d\x9d\xda\x88\xac\x36\xd0\x9c\x27\x07\x76\xac\xee\x9a\xed\x6d\xd8\xd8\x63\x1f\x9d\xdb\x23\xad\x90\x29\x97\x4b\x8a\x08\x8e\x0d\x9a\x19\x30\x43\xf9\x2e\x53\x07\x02\xb2\x5d\x4e\xca\xc7\x78\xd1\xae\xd4\x07\x2d\xa7\xcb\xba\xcd\xb7\x2e\xaa\xdc\xd1\x0d\x74\x32\x9e\x9c\xf8\x84\xca\xc8\xd4\xad\x3c\x8c\x86\xd4\x1c\x05\xdf\x45\xc9\x26\x64\x1d\x25\xeb\x3d\xc3\xa8\xcf\x17\x0e\x61\x29\x12\xd3\x7f\xa2\x31\x39\x68\x1b\x06\xa6\x24\xf9\xcd\x21\x5e\x39\xf4\xb3\xdf\x71\xd1\x44\x28\x6a\x94\x62\xe1\xeb\x5e\x70\xac\x38\xaf\x0d\x81\x02\xc7\x51\xb1\x75\x1f\xb3\xf9\x1f\x44\xcc\x36\x81\xce\x16\xef\x86\x20\x92\xda\x3f\xdd\xe1\xe1\xdb\x48\x27\x45\xe4\xc4\x27\x44\xa4\xb3\xc5\x3b\x64\xc8\x56\xa1\x1a\xb1\xa4\xba\x00\x0e\xc3\xa5\x7e\x7a\x36\xe1\x0e\x47\x7c\xcc\xa7\x45\x03\x5f\x75\xb2\x26\x3e\x96\xb0\x48\xe5\x86\x71\x22\xb7\x27\x96\x23\x88\x61\x9e\xfc\x3e\xee\x7a\x1e\x7d\x9a\xda\x7d\xf4\x6f\xce\x26\x41\x65\x27\x4b\x2d\x06\x33\x44\xd4\x07\x55\xf5\x5a\xd9\x9f\xcc\xe4\x11\x07\x1c\x78\x8c\x46\xdb\x49\x22\x08\x03\xc8\x59\x0d\x3e\x5d\x51\x70\x8a\x0d\x8f\x5b\xb2\x5a\x4e\xb4\x40\xde\x31\x7e\x73\x6f\xcb\x54\x16\x10\x7e\xf0\x1c\x3b\x59\x74\xae\x86\xe6\x30\x27\x44\x63\xd3\x89\x76\x9c\x33\xea\x48\x18\x4b\x71\x83\xe1\x0e\x42\x36\x73\xfc\xf5\xe2\x78\x10\xf8\xa6\x92\x2d\xa2\x88\xa9\x29\x7c\x7a\x71\xfb\xff\x47\x9d\x54\xf8\x24\xe0\xc3\xf6\xb3\x21\x91\x9b\x74\x75\xe8\xb3\xf8\xef\x3b\xc0\xb7\xa0\x2c\x40\xfc\x9d\x25\xb3\xfc\x9d\xdc\x84\x7f\xa7\x92\x44\xe2\x6f\x92\x50\x90\x87\xa7\x17\x67\xd0\x12\xa5\xf1\xdb\x4f\x64\x3a\x7a\x6f\x9c\xe3\xd8\x51\xfb\xa3\xe4\xf8\xf8\xf4\xe4\x72\xbf\x58\xf9\x3e\x43\x1d\x7f\x18\xbd\x66\x1c\xed\x8c\x15\xa9\x61\x20\x2c\x04\xf3\x49\xb6\xf9\x3d\x40\x70\x18\x1e\x22\xc9\x50\x2a\x20\x3b\xf6\x12\x90\x60\xae\x2c\x4b\xbf\xac\x08\xe4\xa6\x66\xec\x0b\x29\x9a\x74\x8b\x70\xe0\x6d\x58\xd3\x7c\x87\x98\xf0\x17\x64\xcb\xaa\x53\x32\x3a\xc3\xc5\x4a\x8e\xe2\x81\xd9\x3b\x25\x4f\xb3\x03\x56\xb3\x73\x36\x17\x93\xfb\x32\xe1\x46\x01\x7e\xaa\x96\xf4\x2c\x55\x62\x52\x27\x66\x83\x79\x76\x94\xbd\x1c\xdf\x47\xc3\xc2\x12\x0e\x9e\x96\x3e\x04\x28\xeb\x41\xe7\xa8\xa0\xe5\x4b\x67\x63\x1d\x4a\xaa\x7f\xa4\x0d\xb7\xa0\xdf\x5c\x96\xb6\x19\x56\x63\x32\x9f\x04\x98\x03\x02\x22\x37\xc0\xf3\x55\xa1\x34\x53\xd4\x48\x9a\x13\x6a\x97\xf2\x71\x80\xe4\x06\x04\x68\x22\x37\xb0\x85\x00\xad\xb6\x68\xf1\xbb\x6e\xe7\x33\x7a\x0b\x94\x00\xad\x84\xd6\xfa\xa5\xf7\x45\x19\x1b\xb9\xf2\x93\x4a\xda\x86\x5e\xbe\x5a\xcd\xde\xa2\x4b\xfb\x62\x31\xc0\xbc\x0f\x3a\x16\xde\x1a\xbc\x74\x2d\x76\x9d\x00\x32\xa1\xef\x12\x46\x6c\x85\x23\x83\xac\xda\xf1\xc0\x51\x84\xfc\x0d\x89\x72\x17\x64\x5e\xc5\x64\x47\x97\xc6\x9d\x7e\xc5\xd3\xa9\xa5\x63\x0e\x0b\x85\x35\xc4\x33\x5d\xe0\xab\x32\x42\xb6\x56\x26\x8c\x0c\x8f\x28\xc9\x98\x3c\x74\x9a\x4a\x83\x68\xf4\x9f\x67\x38\x67\x29\x74\x8d\xeb\x74\xf1\x16\x71\x16\xc1\x77\x02\x2d\x2e\xcf\xf2\x15\x5b\x32\xc4\x53\x8a\x12\x16\x08\xc4\xa8\x64\x39\xcf\x6e\xe3\xdd\x8b\x76\x3f\x12\x43\x04\xbe\x64\x7c\xca\x1c\xe0\xa5\xa1\x39\x85\xeb\x96\x2d\x37\x5a\xe3\x3c\x8d\x40\xa8\x81\x67\x3c\x23\xe5\x3b\x46\x0c\xeb\x04\x77\xe1\x6f\x20\x48\x23\xd8\x43\xce\xfb\xf5\xe4\xb2\xcc\x7d\x29\xf7\xa5\x4b\xae\x77\x1b\xe2\x6f\x8a\x49\x24\x36\x2c\x8d\x82\xdc\xb0\x02\x86\x68\xb6\x0f\x45\x3a\x13\x4c\x9f\x1c\x9b\x69\x97\x49\x04\x82\x42\x26\x87\xe8\x74\x8d\x28\xa3\x7a\x26\xde\x92\x00\x82\x03\x0d\x58\xf9\x8a\xa7\x16\x27\xd5\x30\x8f\x3f\xde\x91\x28\x42\x2b\x50\x7d\x05\x6e\x0a\x7a\x20\x2c\x5b\x35\xfd\x80\x83\xed\x15\x19\xfe\x22\xb2\x0b\x21\x12\x87\x7a\x88\x8b\xdf\x96\x88\x83\x60\x29\xf7\xc1\x6d\xf3\xe2\x40\xe9\x5e\x8e\x38\x6d\x00\x6e\xc5\xb5\x6e\x57\x65\xba\xf0\x7d\x86\x1f\xc2\x98\x9c\x94\x84\x86\x42\x9b\x4c\x05\x35\x0a\x28\xb1\x03\xd5\x30\x90\x1a\xd9\x49\x87\x9f\x50\x40\xf6\x20\x7f\x21\x3b\x5a\x1e\xec\x34\x3c\xe8\x04\x8d\x8a\x78\x5f\xa7\x2b\xe0\x14\x24\x08\xa4\x99\x46\x85\x19\x95\xd6\xdd\xda\xa2\xe0\x06\x62\x13\xf4\x30\x30\xa9\x64\x44\x0e\x48\x1b\xa7\x05\x39\xb4\xe6\x2c\x46\x19\x10\x4f\x28\x89\x71\xf4\x27\x3a\xf1\x6b\xcb\x93\x98\x14\x14\xf6\xf0\x5d\x86\x42\xc2\x58\xa7\x25\x07\x84\x97\x64\x50\x7e\xdf\x8a\x31\x29\x24\xc7\x49\x73\x87\x81\xda\x1d\xc4\xfc\xe5\x2e\x83\xbb\x3a\xa5\x42\xe2\x28\x22\x34\x44\x18\xfd\x77\x4a\xfc\x1b\x21\x31\x97\xb9\x8b\x5f\x5c\x13\x09\x89\x64\x89\x98\x3f\x26\xc5\xfb\x1e\xf6\xfe\x2c\xde\xf7\xcc\xfb\x1e\xa1\xde\x96\xa5\xdc\x33\xeb\xb8\xdb\x55\x92\xc6\x4d\x91\x91\xbd\x5e\xcf\x8e\x7a\xc6\xd5\x7e\xc5\x44\x69\x00\x57\x51\xb9\x43\xc6\xe7\xf9\xdb\x9d\x42\x7e\xae\x6f\xb8\xa2\x4b\x48\x58\x97\x40\xd7\x51\xfa\x71\x7a\x81\x29\xaa\xd7\xb3\xa3\x12\x0f\xed\x83\xe7\x90\xb0\x61\x03\x57\x74\xbe\xe5\x41\x3b\x41\x16\xaf\x0e\x76\x67\x23\x07\x1d\x73\x74\x12\x2c\x33\x97\xe3\x74\x34\xc2\x76\xe2\x5d\xbe\x70\xac\xef\x51\x2b\x83\x7f\x49\xe4\x79\xa2\xb6\xa8\xbb\x83\x42\x1d\xd3\x88\x08\xbd\x51\xbf\x93\x2c\x27\x55\xbd\x87\xd4\xd0\x04\x91\x8c\x6f\x0f\xd1\xd5\x4b\x2d\x48\xf4\x32\x25\x81\x9a\xf9\x99\x5c\x4b\xf3\xad\x74\x11\xb0\x4f\x49\x5f\x94\xf1\x92\x45\x34\x79\xbe\x9e\x1d\x95\xc7\xb5\xb3\x83\x1c\x84\x6b\x89\xc4\x25\x3c\x6e\x73\x97\x76\x56\xd3\xe2\xe7\xb4\xa5\xcb\x6d\x11\xe6\x2b\x22\x39\xe6\x5b\xf4\x5f\xcb\xf3\xb3\xf9\xff\x2c\xde\xbe\x29\x32\x85\xc5\x01\x12\xa9\xbf\x41\x58\x20\x1d\xcd\xb3\xdc\x2f\x67\x5c\x67\x94\xeb\x14\x63\x02\xae\x27\x77\xf7\xc9\x80\xc5\x43\xca\x05\xdc\xb8\xf7\x3a\x71\x80\x0c\xc7\xe4\x05\x8e\x49\x34\x6d\x76\xec\xc3\xbe\xa4\x1e\x80\x50\x62\x3b\xc6\x09\xf6\x89\x6c\x1d\xb9\x32\x8a\x10\xf8\x9e\x17\xb7\x0b\xcd\xb5\x5e\xdd\xd6\xd0\x4b\x7d\x5d\xc5\x61\x52\x2d\x3c\xe8\x7d\x4c\xef\x9e\x20\xc6\x1f\x97\xe4\x53\xab\x44\x3a\xb5\x13\x13\x3a\xba\xed\xe4\xb9\x92\x26\xfa\x6d\x52\x25\x2c\xb5\x60\xea\xc7\xf3\x6d\xe4\xf5\x6a\xd7\x12\xb9\x13\x03\x0b\x2e\x14\xe6\xb8\x5c\xbe\xfa\xe6\x42\x43\xfd\x79\x98\x2c\x4a\x63\x18\xa2\xfa\x2e\xf7\x2b\x24\x21\x5e\x6d\xa5\x63\x80\xa9\xa5\x55\x89\xeb\x7f\xff\xd7\x54\xc1\xa4\x1d\x6a\xb7\xe1\x48\x07\xdc\x59\xe6\x89\x65\xda\xd9\xa5\xda\x89\xf1\x35\x7b\x6c\x22\x51\xe7\xa4\xa8\x9b\x60\x0d\x6c\x27\xdd\xde\x62\x8a\x9e\xbf\x5e\x7a\x8d\x1a\x08\xe8\xdd\xf9\xc9\x39\xfa\x15\x47\x24\x28\x0e\x38\x69\x8c\x93\x04\x02\xb4\x26\x90\xf9\x01\x01\x92\x1b\xce\xee\x14\x11\xe0\x9c\x0d\xcf\x4b\xbb\x9f\xde\xab\xee\x02\x48\x4e\x7c\x71\xcc\x22\xb5\xa7\xae\xa6\x24\xb7\xf8\x0b\x21\xc7\x34\x8d\x30\x57\xa6\x31\xd8\x6d\x28\x37\x9a\x12\x2b\xe3\x8c\xff\xaf\xef\x2e\x38\x4d\xcf\xb2\x34\x2c\x83\x99\xc4\x74\x75\xd0\x74\xb5\xcd\x22\xa9\x3e\xd6\x3e\x7f\x7e\x0f\x5e\xd7\xa7\xd0\x45\x00\x76\xa5\x24\x02\xe6\x8b\xf7\xdf\x6f\xa4\x4c\xc4\xd3\xf9\x5c\xfd\x75\x88\xef\xc4\x21\x8e\xf1\x27\x46\x0f\x7d\x16\xcf\x17\xbf\x2d\x75\xdd\x9d\x17\x79\x9b\xb9\xda\x55\x08\x39\xff\x45\x00\xd7\xfe\xfe\x1c\xdf\x09\x6f\x67\x02\x1e\x16\x9e\x19\x93\x5f\x18\xd8\xa1\xb2\xf4\xe1\x7b\x9b\xbe\x61\xec\xb6\x23\x5f\x88\xf5\xeb\xd9\x91\x45\x72\xcd\x9d\x4e\x9e\xe9\x38\x20\xe4\xf4\xe5\x33\xe9\xa6\xc8\x8c\x72\xb2\x78\x4b\xee\xc5\x24\x56\x9e\xed\xb5\x4e\x4f\x34\xd0\x1d\x9f\x9e\x5c\x3a\xee\xd2\xca\x2d\xab\xea\xbb\xc7\x0d\xd4\x1e\x41\xeb\x65\x02\x3e\x59\x6f\xd1\x95\x9f\x0a\xc9\x62\xb4\x78\x7b\xba\x2b\x2f\x93\x3d\xf3\x70\x4c\x3c\x91\x26\x09\xe3\x72\xfe\xe4\x00\x5d\xeb\xac\x13\x4f\x88\xf8\x7a\x96\xff\xa5\xfe\xc7\x38\xba\xd6\xf7\xd6\x88\x7f\x3d\x73\xf2\x5c\x72\x26\x1a\xc1\x21\x0b\x03\x6a\xba\xec\x58\x55\xd3\xe4\x00\xfd\xf3\xcf\x94\xc9\x9f\x72\xae\xb2\xbf\xca\x4f\xf3\x27\x8c\x9b\x87\x19\x97\xd9\xff\x5d\x77\x96\xf7\xb3\x5f\x15\x61\xd7\xca\x89\x3a\xd6\xa0\x96\x72\x55\x0d\x6a\xce\x2b\xd0\x43\xdc\x4e\x77\x99\xf2\x1b\x12\x13\x99\x95\x49\xca\xc2\xfa\xda\xaa\x88\x8f\x16\xbf\xef\x4c\x5a\x99\x83\x81\xfd\xf9\xe3\x4f\x8c\x82\x87\xef\x30\x07\x2f\x33\x9e\xec\x07\xb7\x88\x66\xd6\x6d\xc3\x74\x87\x74\x64\x0a\x27\x35\xb8\x6d\x8f\xf1\xae\x98\x94\x11\x70\xe6\xdf\xc0\xc0\xdc\xd1\x02\x77\x9e\x95\x9b\x5a\x89\xfb\x11\x16\x82\xf8\x6f\x18\x0e\x9e\xe1\x48\x79\xf2\xfc\x0c\xc7\x0f\x53\xd9\x0b\x93\xf6\x0b\x48\x1f\xd9\xac\x0c\xbf\x22\xcb\x05\x54\x42\x2e\x56\xf7\xb0\x96\x34\xd7\xaf\x52\x67\xe2\x2d\xe2\xd4\x41\xd0\x93\xb3\xe5\x1e\xf8\x7c\x75\x9c\x81\x1d\x0e\x02\x0e\x62\x67\xc7\xb7\x89\xef\xd1\x62\xef\x32\x7f\x6c\x90\xd2\xf4\xe9\x05\x54\x78\xa6\xc9\x93\xec\xb8\x5b\x39\xf3\x27\x67\x4b\x14\x31\x76\x53\xad\xfa\x34\x22\x68\x3f\xbc\xf7\xeb\xd9\x51\x75\x04\xba\x4a\x5c\x3f\x47\xbd\x88\x39\x45\x0c\x0d\x56\xe2\x3c\x91\x24\x26\x9f\xa0\xd5\x7f\x69\x89\x89\x54\xe4\x93\x15\x41\x15\xe8\xea\xf9\xb3\xa5\x8e\x91\xc7\xe4\x93\x76\xe5\x7a\xfd\xdf\xe7\xc7\x3f\x36\x3d\x47\x58\x09\x8f\xe5\x7c\xd5\xdc\xdb\x21\xea\xca\xd9\x19\xec\xca\x0e\xe4\xe2\x7a\x76\x54\x1f\x60\x3b\x52\xdd\x43\x7c\x72\x9a\xeb\x6a\x16\xc2\x17\x1c\xd6\xe4\xe3\xbd\x90\x9e\x3c\xa6\x9a\x13\x16\x27\x44\x64\xd7\xca\x06\x57\xd2\xdb\x49\xda\x4a\xc3\xda\xdd\x4d\xba\x82\x08\xe4\x73\x9d\xa0\x5c\x2f\x78\xdb\xd1\x97\x43\x81\x17\x03\x71\xe4\x13\xa0\x0f\xa6\xbb\x0f\x66\x4b\x56\xf3\x44\xc9\x27\x42\x43\x4f\x6e\xc0\x33\xef\x39\x16\xb7\x6c\xf1\x2f\x9b\x64\x0b\xd4\x52\x4c\xfd\xec\xb3\x00\x8e\xcc\x4f\x3f\xcf\xf5\x5f\x86\xbf\x76\xf3\xff\xe6\x43\xdf\x17\x2c\x10\x17\xc0\x95\xcd\x8c\x8b\x80\xff\x5f\x89\x9e\xb3\x5b\xe0\x9c\x04\xf0\x2c\x3f\x23\x3e\x66\x71\x8c\x5d\x8b\xdd\x56\xec\xf0\xdc\x90\x44\x1f\xb2\x9d\xf6\x87\xef\x04\x2a\x8e\xa0\x13\xe5\x56\x64\xaf\x3b\x19\x77\x41\x34\xb3\xd7\x8c\xb2\x31\xd7\x36\xfa\xd6\x01\x27\xbc\x31\xd6\x07\xe9\x02\x82\xce\x66\x84\x00\xad\x60\xcd\x38\xd4\x46\x58\xe0\x64\x56\xb8\x09\x1a\x97\x78\x87\xc8\x74\x64\x17\x2d\x62\xdd\xef\x14\xa6\xc2\x98\xc9\x8d\xb8\xca\x2f\x36\xec\x9c\xb1\x56\x0f\x31\x15\xe0\x99\xd7\x3d\x93\xd7\xe9\xad\x19\xf7\x34\x64\xe3\x68\x57\xcf\xf5\x89\xf6\xcc\x8a\x3f\x9d\x04\x66\xf8\xea\x75\x18\x07\x33\x73\x3d\x3b\x6a\x8e\x51\xfb\x90\x1d\x4c\x0e\x3c\xb2\x2a\x5f\xaa\x19\x78\xb5\x6a\x77\x7a\xf5\xb2\xe5\x8e\xe0\xb8\x83\xb0\x2e\x5d\xe7\x29\x18\x20\x90\x10\x9b\xbc\x06\x65\x96\x73\x4d\xc4\x48\x45\x0d\x26\x6a\x1d\xe4\xb7\x7c\x44\x27\xb1\xed\xaa\xef\xb7\xc3\x3d\x0f\x41\x6a\xbb\xf9\x9a\x75\xe5\x86\x6d\xcd\x33\x66\xb3\x3d\xf2\xc4\x1b\xf3\x41\xa4\xad\x12\xcc\x0e\x13\x4d\x51\xe2\xfe\x7d\x5f\x07\x8d\xd3\xf3\x8b\xd6\xad\x7d\xa7\x8f\x92\x35\x7f\x1d\x8b\xd7\xb0\x3d\x3d\x19\x5c\x12\xa6\x41\x61\xc0\x86\xa8\xa3\xf5\x37\x71\x4a\xdd\xe0\xda\x7d\x43\x55\xe9\x5e\x1f\x6e\xa2\x5b\xcc\x09\xa6\xd9\xf5\xd2\xa7\xe8\xc3\xf5\x2c\x4c\x7e\xbc\x9e\x7d\x40\x44\xa0\x97\xa6\xfc\xcf\x45\xca\x13\x26\x00\x2d\x97\x27\xe8\x7b\xc3\xdd\x93\x03\xf5\x2e\x61\x3f\x98\x77\x2f\x38\xbb\x25\x82\x30\x0a\x01\x52\xc6\xa0\x5e\xd6\xaf\x08\x3f\x7f\xe5\xdd\x86\xb3\x34\xdc\x24\xa9\x44\x45\xa8\x01\xbd\x3a\x31\xaf\xc9\xfc\xb5\x63\x16\xe9\xc7\x6e\x19\xe1\xb6\xc1\x64\xde\x5f\x16\xda\x0e\x93\x1f\xb3\xff\xe4\xbb\x96\xfe\xf1\x95\x9b\x13\xf6\x43\xa3\xb9\x7d\xc8\xe5\x56\xc2\x6f\xb6\x6a\x97\x42\xa5\xa5\x6c\xb6\x6c\x11\x4c\xc9\x5c\xc2\xe4\xc7\xea\x6f\x40\xd3\xb8\xf9\xf1\x88\xfa\x6b\x0a\x2c\xd9\x0f\xf5\x47\xc2\x6f\x3e\x92\x3f\x94\x51\xb1\xf4\xad\x89\x47\x35\x1b\x75\xcd\xa0\x18\x9d\x50\x61\x8f\x01\xb4\x07\x36\xda\x62\x29\x43\x93\x29\xea\xe9\x10\x9d\xb9\x13\x35\xb7\xaa\x23\x64\x67\xd9\xf1\x59\x36\x90\x7d\xa7\x24\x6d\xf1\x3c\x3b\xde\xd9\xf1\xc4\x8e\xac\x1d\x8b\x46\x3b\x98\xdb\x57\x89\xf6\xbd\x75\x33\x66\xd0\xf4\x55\x86\xc4\xe8\x3b\x7c\x84\x9a\x77\x5a\x0b\xd2\xb5\x9d\x2e\xf4\xed\x08\x87\x6c\x91\xed\xd1\xf0\xee\xf0\x92\xf9\x71\x92\x3a\xf1\x95\xe4\xe5\x52\x19\x26\xb9\xc1\x52\x57\x05\x28\x4e\x5d\x74\x6a\x72\xd3\x95\x1e\x58\x32\x7e\x74\x3f\xba\x9b\xc6\x11\xf1\x33\xfb\x29\x4f\xcf\xe7\xc3\x16\x41\x4c\xe8\x71\xfe\x7d\xab\x51\x7e\x4d\x7e\x43\x6e\xf2\x18\x5f\x51\x5d\x0f\xd3\x2d\xba\x2a\xdb\x59\x71\x2b\x6f\x17\x2b\xdf\x25\x25\xcc\xcb\x6f\x7a\x4c\x54\xfe\x9e\x3f\x2e\x75\xe2\xb1\xb5\x97\x53\x72\x0b\x0a\x56\x58\x6b\x86\xcc\xf7\x65\xe6\x7a\x76\x64\x1d\xee\x3e\x57\x1a\xac\xfa\xb6\xa9\x71\xc2\xb9\xa4\xe3\x1b\x15\x3b\x57\x9b\xc5\xb2\xa5\xa2\x15\x16\x10\xa0\xdd\x97\x45\x86\x5f\xc9\xda\xa3\x0b\xfb\x0c\x1a\xf8\x05\x86\x07\x5d\xa6\x7b\xb7\x88\xeb\x3b\x29\xce\x05\x17\x06\x1e\x24\x8c\x2a\xe6\xe0\x40\xfb\xde\x4e\x6d\xc6\x7d\xa8\xc1\xad\x2f\xb5\xb3\x5c\x04\x01\xa3\x17\xf9\x9d\x09\xe7\x33\xad\x6a\xf3\x91\x33\xbe\xab\xc2\xb4\xc5\x4e\x3a\xd4\xdc\xa5\x25\x07\x21\x77\xca\x68\x42\xd8\x69\xfb\x0c\xc3\x2e\x9d\xca\x0d\x63\xfa\xe9\xb5\x02\x4a\x9b\x1d\xb4\xa3\x4b\xb4\x3a\xa5\x21\x1f\xfb\xe5\x1e\x9c\x24\x6f\xa1\x19\x4d\x74\x69\x7b\xc1\xe1\x96\xc0\xdd\x38\x12\xa9\x64\x4b\x1f\x47\x23\x5d\x09\x1f\xb8\xcc\xee\x0e\x8d\x6c\xdf\xfa\x21\xc7\x41\xcd\x61\x35\x4e\xe8\xb0\x1e\xd9\xee\xa3\x04\x4e\x71\xd4\x91\xed\xd1\xd9\x7e\x2d\x5a\xcf\x9f\x3b\xdb\x91\x18\x87\xf0\x2c\x25\x51\x30\x52\xce\x1f\x2f\xdb\x8b\x12\xef\xf9\x7d\x9a\x0a\x6f\x76\xcb\x6a\x91\x60\x8b\x1d\x59\x26\x47\xbb\xcd\xd7\x8c\xa1\x26\xeb\x9a\xca\x0f\xac\xb3\xb6\x2e\x26\xbb\x79\xde\x07\xda\x29\xa8\x19\x7d\x51\xd0\x4e\xa4\x05\xd7\x7a\xd2\x09\x5a\x52\x54\xcb\x91\x09\x0b\xde\xb7\x21\x62\xb5\xd9\x43\x72\xb6\xd4\xae\x9c\x93\xf6\xfa\x0f\x34\x8d\x57\x6d\xe1\x5c\x46\x4f\x40\x6d\x77\x9f\x61\x01\x7b\xe5\x23\xe5\x84\x2e\x80\xfb\x40\x25\x0e\x61\xb1\x62\xb7\xb0\x37\x5d\x91\x30\x69\xaa\xc4\x11\x46\x97\x92\x63\x09\xe1\xb8\xcf\x80\x25\x4c\xe6\x26\x73\xc1\x58\x33\xc3\xa1\x9d\x1f\x37\xe8\xa8\x18\x8a\x4d\x4f\x7d\xf2\x77\x14\x6b\xe7\x18\xfb\x45\x39\x21\x06\xd8\x77\x41\x57\xaa\xe3\xdd\x89\x73\x71\xca\xab\x1e\x7b\xc5\x63\x87\x3b\xde\x5d\x9d\x35\x8e\x6f\x6b\xbd\x5c\xcf\x8e\xaa\xec\x58\x6e\x2b\x94\x0f\x4a\x07\xef\xc4\x4e\x4f\x1e\xe4\x89\x56\xc6\x1c\x88\x72\xf1\xd9\x3c\xcc\x89\xcc\x65\x79\x93\x02\x30\xee\x34\x76\x54\x07\xad\x1b\x96\x37\xcc\xc7\xd1\x3e\xd9\x05\xe6\x8b\x57\xb8\xc6\x03\x52\x66\x1f\x15\x1f\xc2\xda\x67\xa8\x23\x69\x97\xf4\x2a\x79\xda\x72\xae\xaf\x44\xb0\xd4\xa5\x33\x27\x90\x41\x56\x37\xaa\xc2\xa9\x29\xe4\x8a\x63\x46\x43\xbd\xd8\xee\x0a\x8e\x22\x42\x47\xe7\x9a\x4c\xdf\x61\xbb\xb4\x9c\xb0\x78\x37\x35\xed\x42\xb6\x5a\xdf\x24\x80\xe8\x33\x2a\x39\x8b\x44\x63\x2e\x74\x24\x3f\x0c\x09\xf7\x0d\xa5\xd9\x82\x68\xcb\x57\xc3\x76\x7f\x11\x1b\xb7\xf3\xca\x8a\x96\xbe\x86\x51\x2b\x74\xd1\x78\xec\xe1\x70\x41\xe0\x02\xcb\xd6\xbd\x57\xa7\x8f\xa0\x4b\xda\x55\x4a\xd6\x9e\x7e\xbd\x0c\xb2\xb1\x46\xaf\xb5\xd7\x2a\x16\xab\xb6\x5a\xb5\xd0\x2f\x9c\x89\xf7\x10\x1a\x44\x76\x69\x3d\xd5\x15\x5e\x1f\x41\xec\x13\x3f\x71\xa1\x5e\x99\x42\xe7\xcd\xf2\x4c\xed\x77\x18\x59\x1c\x13\x99\xb7\x78\x8b\x29\x59\x83\x68\xe6\xed\xb8\x40\xfa\xb1\x26\x69\x3e\x7d\x20\x36\xe8\x45\x94\x7e\x44\x71\x4e\x39\x5f\x60\x5f\x12\xa9\x6b\x0e\x21\x46\x91\x29\x4a\xe4\x84\xe3\xe3\x7b\xb1\xce\x26\x7d\x28\xb8\x47\xc2\x83\xea\x28\x2b\x9c\x27\x19\xba\x01\x48\x90\xe4\xd8\xbf\x41\x6c\xad\x39\xfb\x4e\x20\xb1\xa5\x3e\x4a\x38\xd3\x7b\xde\x9f\x32\x0c\x24\x02\xa9\x7d\xdf\x2d\x8e\xcc\x47\x30\xcd\x11\x1f\xa1\x21\xf2\xbc\x90\x48\x4f\xb5\xf2\x24\x0e\xf5\x40\xb3\x47\x94\x49\x10\x1e\x87\xb5\x5a\x95\x14\x71\x27\xb9\x3d\x1c\x46\xef\xeb\x8b\x91\x55\x33\x31\x25\x92\x76\xa5\xfc\xee\x36\xc0\x75\x35\x43\x63\x0f\x99\xe5\x64\x17\xee\x01\xbd\x82\x28\x46\xf9\x74\xc8\xbe\x73\xb0\x76\x15\xf1\xbd\xf4\x39\xe0\x53\x6d\x38\x38\xa7\xed\x17\x26\x87\x4c\x5d\xb5\x1f\xe3\xa9\x2f\x33\xfe\x24\x2b\x7d\x5d\x27\x66\x41\xf6\x39\x13\x9f\x83\xce\x0c\xdb\x00\x0a\x20\x89\xd8\x16\xdd\xc0\x16\x61\xb1\x7b\xd7\x49\x58\xf7\xd1\xe5\xb0\xc4\x54\xe5\x45\x29\xd1\xef\x2b\xb0\x1c\xab\x2b\x6a\x74\x96\x81\x9d\xca\xc8\x65\xb5\x0d\xd5\x1b\x80\x67\x9d\x6d\x36\x19\xd9\x0c\x6d\x92\xd5\xd4\xa5\x08\x9a\x92\x4f\x5e\x57\xae\xa8\x2e\x9b\x61\x58\xa9\x2c\x72\x3e\xad\xaa\x15\xd0\x14\x06\x29\x8c\x1a\x7e\x5a\xfa\xe5\x39\xab\xac\xe2\x96\x4a\xf9\x6d\x8b\x38\x4b\x65\x92\xca\x01\x8e\x64\x97\x29\x9f\x6b\x22\x28\x20\x5c\xd7\x7d\xdd\x16\xe5\xa6\x13\xce\x94\x03\x02\x41\x5e\x18\x12\x49\x88\x13\x7d\x91\x0d\x7d\x9f\x7d\xc6\x6f\x57\xef\x1e\xf9\x59\x4e\xca\x13\x74\x92\x4d\x42\xbd\x14\x1f\xce\x7f\x2e\x55\xa4\x54\x43\xf7\x94\xb9\x1d\x39\xcd\x92\xaf\xcf\xe0\x00\x0c\xce\xb2\xf6\xf6\x50\x83\x29\x0e\xae\xcb\x71\xa2\x65\xb9\x1e\xe7\x21\x3a\xc6\x14\xad\x00\x61\xb4\xe2\x98\xfa\x9b\x03\x5d\xe4\x5a\x7f\x64\x43\xfb\x42\x1b\x5c\x89\xd4\x0f\xfe\x64\xc1\x34\x7d\xf5\x97\xf4\xd7\x3e\xf9\x1e\xa2\x51\x4e\xbe\x62\xe1\x97\xcb\x37\xa8\x8b\xf5\x17\x6a\x15\xfd\x88\xe3\x24\x82\x03\x84\x93\xc4\x0b\xe0\xd6\x49\x2e\xd3\x75\x34\x41\x09\x0e\x23\x36\x9b\x95\x1d\x58\x21\x60\x6a\x58\x0e\x40\x62\x12\x99\x92\x93\x7f\x36\xca\xc4\x16\xd5\x29\x41\xbd\x51\x87\x3b\x1c\x04\x65\x5f\xbc\x54\x91\x72\x0c\x0e\xdf\x17\x2b\x15\xe0\xbd\xac\x16\x78\x6d\xaf\x38\xac\x27\xc6\x1e\xf6\xfc\x6e\x03\x28\x24\xd2\xcc\x30\x94\xd2\x00\xb8\x29\x32\x9d\xf3\x5d\x0b\x1a\x93\x08\x44\x51\xf6\x3f\x9b\x89\x6a\x01\xfa\xa7\xde\x02\x41\x60\x3e\xcc\x15\x63\xe7\xa8\xd8\x74\xac\xe0\x38\xf9\xa9\x9f\x9d\x5e\xb0\x80\x18\x93\x7d\xf7\x63\x9a\x86\x19\x45\xf9\x5b\x09\xca\x0a\x0c\x8c\xf9\x1b\x4c\x43\xc7\xd4\xfa\x7d\x48\xf7\x8e\x5b\x39\xfe\x7b\x2e\xe4\x4a\x97\xbb\x45\xb2\xac\x4a\xed\x62\x77\xe9\xf1\x8e\x2b\x2d\xd2\x83\xdd\x26\x64\xee\x6c\x46\xf7\xd5\x75\x7f\xfd\x45\x2c\x37\x0f\xf2\x5c\xe3\x52\xf9\x89\xe4\x16\x90\xe6\x50\x5f\xa3\x31\xf1\xe4\x9a\x23\x68\x0a\xcc\x67\x3f\xe8\x2a\xbf\xb9\x4b\xa9\xa5\x14\x33\xaa\xde\x53\x36\xb6\x26\x34\x40\xa5\x4a\xf5\x95\xc0\x07\x4e\x92\x68\x6b\x04\x79\x75\xad\x93\x7f\x3d\xb1\x15\x12\x4c\x25\xa1\x15\x16\x70\x3d\x7b\xef\xa4\xd9\xaf\x3a\x86\xec\xba\x42\x69\x1c\xd5\xda\x43\x6a\x3c\xd9\xff\xde\x77\xde\x09\x5d\x2e\x5f\x0d\x8b\xb7\x76\x29\x53\x35\xcf\x97\x91\xfc\xca\xe4\x72\xf9\x4a\xef\x48\x77\x5f\x4a\xc0\xa9\xdc\x00\x95\xfa\x63\xbb\x4e\x72\xde\x97\x7c\xef\x2c\x49\xf9\x3e\xb0\xfa\xce\x28\x5c\xb1\xa4\xbc\x23\xc3\x69\x43\xff\x5a\xd7\x26\xe5\xb7\xb2\x10\x57\x20\xc0\xe8\x37\x24\xf2\x3f\x77\xf9\xbf\x4f\x19\x0f\xe7\xbc\x51\x3a\x7d\x00\xf2\x7c\x1d\xc6\xfa\x25\x2e\xda\x73\x86\x06\xae\x64\x8a\xc4\xfd\x2c\x64\x23\x29\x4f\xe0\xd7\x2a\x43\x3c\x68\x78\x53\x0d\x3c\xb7\xad\x8d\x75\xe1\x36\xfc\x86\xce\xb9\xff\xc5\xc3\x16\xf5\x62\xec\xbb\xba\x3c\x19\x44\xde\x4f\x48\xa2\xbf\xd7\x8a\xd7\xbb\x04\x9f\x83\x14\xe6\xee\xcd\xa0\xec\xa4\x1b\xd8\x2e\x2e\xcf\x86\xa7\x25\x99\xf7\x87\x9e\x75\x39\x59\x53\x1b\x2f\xfb\x7e\x31\xa9\x99\xb5\xf1\xfa\xed\x12\x41\x21\xa5\xfc\x53\x5d\xee\xd5\x61\xdd\xa8\x57\x74\x35\xa2\x9c\x64\x49\x99\x2d\xd8\xd3\x38\x99\xa7\xe8\xf4\x22\x2f\x27\x85\x08\xcd\x3e\xed\x4b\x99\xc4\x95\xfb\x60\xbd\xe7\xed\xdd\x64\x1e\xe5\xaa\xfe\xfc\xe8\xf3\xa3\xff\x0d\x00\x00\xff\xff\x96\x2b\xff\xb7\x0f\xa2\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x2a, 0x79, 0x82, 0x32, 0xc0, 0xfd, 0x33, 0xa9, 0x2b, 0xe0, 0x2a, 0x2e, 0x1c, 0x63, 0x43, 0xf8, 0xfd, 0xa9, 0xf9, 0xa9, 0x9, 0x63, 0xc6, 0x3b, 0x81, 0xed, 0xc7, 0xe3, 0xdf, 0x81, 0x6e}}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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
