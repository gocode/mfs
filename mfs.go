package mfs

const (
	PathSeperator = "/"
)

type FileSystem struct {
	rootDir *Dir
}

func New() *FileSystem {
	return &FileSystem{rootDir: &Dir{Name: "/"}}
}

func (fs *FileSystem) CreateFile(path string) *File {

}

func (fs *FileSystem) CreateDir(path, name string) {

}

func (fs *FileSystem) ReadDir(path string) *[]Dir {

}

type Dir struct {
	Name  string
	dirs  []*Dir
	files []*File
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

type File struct {
	Name   string
	data   []byte
	offset int
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
