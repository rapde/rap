package zipfs

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var _ Finfo = &Dinfo{}

type Dinfo struct {
	fi     os.FileInfo
	finfos []os.FileInfo
	files  map[string]Finfo
	fnames []string
}

func newDir(fi os.FileInfo) *Dinfo {
	return &Dinfo{fi: fi, files: make(map[string]Finfo, 0)}
}

func (d *Dinfo) FileInfo() os.FileInfo {
	return d.fi
}

func (d *Dinfo) Open() (http.File, error) {
	return &Dir{Dinfo: d}, nil
}

func (d *Dinfo) addFile(fn string, fi Finfo) {
	fname := filepath.Base(fn)

	if d.files[fname] == nil {
		d.files[fname] = fi
		d.fnames = append(d.fnames, fname)
	}
}

type DirInfo struct {
	name    string
	modTime time.Time
}

func (f *DirInfo) Name() string {
	return f.name
}

func (f *DirInfo) Size() int64 {
	return 0
}

func (f *DirInfo) Mode() os.FileMode {
	return os.ModeDir | 0755
}

func (f *DirInfo) IsDir() bool {
	return true
}

func (f *DirInfo) ModTime() time.Time {
	return f.modTime
}

func (f *DirInfo) Sys() interface{} {
	return nil
}
