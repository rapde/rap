package zipfs

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileSystem zip file system
type FileSystem struct {
	r    Reader
	dirs map[string]*Dinfo
}

// FileSystemCloser zip file system with a closer
type FileSystemCloser struct {
	FileSystem
	Filename string
	c        io.Closer
}

// Options with FileSystem
type Options struct {
	Prefix string
	Ignore []string
}

// Finfo for FileSystem
type Finfo interface {
	FileInfo() os.FileInfo
	Open() (http.File, error)
}

// New file system with zip data
func New(data []byte, opts *Options) (*FileSystem, error) {
	b := bytes.NewReader(data)
	r, err := zip.NewReader(b, b.Size())
	if err != nil {
		return nil, err
	}

	dirs := newDirs(r.File, time.Now(), opts)

	fs := &FileSystem{
		r:    &reader{r},
		dirs: dirs,
	}

	return fs, nil
}

// Open zip file return file system closer
func Open(name string, opts *Options) (*FileSystemCloser, error) {
	rc, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(name)
	if err != nil {
		rc.Close()
		return nil, err
	}

	dirs := newDirs(rc.File, fi.ModTime(), opts)
	fs := FileSystem{
		r:    &readCloser{rc},
		dirs: dirs,
	}

	return &FileSystemCloser{fs, name, rc}, nil
}

// Open implement http.FileSystem
func (fs *FileSystem) Open(name string) (file http.File, err error) {
	name = strings.Trim(name, "/")
	if name == "" {
		name = "."
	} else {
		name = strings.ToLower(name)
	}

	d := fs.dirs[name]
	if d != nil {
		return d.Open()
	}

	d = fs.dirs[filepath.Dir(name)]
	if d == nil {
		return nil, os.ErrNotExist
	}

	f := d.files[filepath.Base(name)]
	if f == nil {
		return nil, os.ErrNotExist
	}

	return f.Open()
}

// Stat return fileinfo
func (fs *FileSystem) Stat(abspath string) (os.FileInfo, error) {
	for _, f := range fs.r.File() {
		if f.Name == abspath {
			return f.FileInfo(), nil
		}
	}

	return nil, os.ErrNotExist
}

// Close implement io.Closer
func (z *FileSystemCloser) Close() error {
	return z.c.Close()
}

func newDirs(files []*zip.File, modTime time.Time, opts *Options) map[string]*Dinfo {
	dirs := make(map[string]*Dinfo)

	// opts
	if opts == nil {
		opts = &Options{}
	}

	if opts.Ignore == nil {
		opts.Ignore = Ignore
	}

	// ignore files
	ig, _ := NewIgnore(opts.Ignore)

	// prefix
	prefix := opts.Prefix
	if 0 < len(prefix) {
		prefix = strings.ToLower(prefix)
		prefix = strings.Trim(prefix, "/") + "/"
	}

	// root directory
	dirs["."] = newDir(&DirInfo{name: "/", modTime: modTime})

	for _, f := range files {
		fi := f.FileHeader.FileInfo()
		org := strings.Trim(f.FileHeader.Name, "/")
		fn := strings.ToLower(org)

		// prefix check
		if 0 < len(prefix) {
			if !strings.HasPrefix(fn, prefix) {
				continue
			}
			fn = strings.TrimPrefix(fn, prefix)
			fn = strings.Trim(fn, "/")
			org = org[len(org)-len(fn):]
		}

		// ignore file
		if ig != nil && ig.MatchString(fn) {
			continue
		}

		if fi.IsDir() {
			if fn == "" {
				fn = "."
			}

			dirs[fn] = newDir(fi)

			if fn == "." {
				continue
			}
		}

		dn := filepath.Dir(fn)
		if dirs[dn] == nil {
			mkpath(dirs, filepath.Dir(org), fi.ModTime())
		}

		d := dirs[dn]
		d.addFile(fn, &ZipFile{f})
	}

	return dirs
}

func mkpath(dirs map[string]*Dinfo, fn string, t time.Time) {
	subdir := strings.Split(fn, "/")
	parent := dirs["."]
	dn := ""

	for _, d := range subdir {
		dn += strings.ToLower(d)

		if dirs[dn] == nil {
			dir := newDir(&DirInfo{name: d, modTime: t})
			dirs[dn] = dir
			if parent != nil {
				parent.addFile(dn, dir)
			}
		}

		parent = dirs[dn]
		dn += "/"
	}
}
