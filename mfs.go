package mfs

import (
	"errors"
	"strings"
)

const (
	PathSeperator = "/"
)

type FileSystem struct {
	rootDir *Dir
}

type Dir struct {
	Name  string
	dirs  []*Dir
	files []*File
}

type File struct {
	Name   string
	data   []byte
	offset int
}

func New() *FileSystem {
	return &FileSystem{rootDir: &Dir{Name: "/"}}
}

func (fs *FileSystem) CreateFile(path string) (*File, error) {
	rd := fs.rootDir
	p := strings.Split(path, "/")
	for i := 0; i < len(p)-1; i++ {
		rd = rd.getDir(p[i])
		if rd == nil {
			return nil, errors.New("path does not exist")
		}
	}
	f := &File{Name: p[len(p)-1]}
	rd.files = append(rd.files)
	return f, nil
}

func (fs *FileSystem) CreateDir(path, name string) error {
	p := strings.Split(path, "/")
	rd := navigate(fs.rootDir, p[len(p)-1:])
	if rd == nil {
		return errors.New("path does not exist")
	}

	rd.dirs = append(rd.dirs, &Dir{Name: p[len(p)-1]})
	return nil
}

func (fs *FileSystem) ReadDir(path string) ([]*Dir, error) {
	rd := navigate(fs.rootDir, strings.Split(path, "/"))
	if rd == nil {
		return nil, errors.New("path does not exist")
	}
	return rd.dirs, nil
}

func navigate(dir *Dir, path []string) *Dir {
	for _, p := range path {
		dir = dir.getDir(p)
		if dir == nil {
			return nil
		}
	}

	return nil
}

func (dir *Dir) getDir(name string) *Dir {
	for _, d := range dir.dirs {
		if d.Name == name {
			return d
		}
	}

	return nil
}

func (dir *Dir) getFile(name string) *File {
	for _, f := range dir.files {
		if f.Name == name {
			return f
		}
	}

	return nil
}

func (f *File) Write(p []byte) (int, error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *File) Read(p []byte) (int, error) {
	n := 0
	for i := f.offset; i < len(f.data); i++ {
		p[n] = f.data[i]
		n++
	}
	return n, nil
}

func (f *File) ReadAt(p []byte, offset int64) (int, error) {
	n := 0
	for i := offset; i < int64(len(f.data)); i++ {
		p[n] = f.data[i]
		n++
	}
	return n, nil
}
