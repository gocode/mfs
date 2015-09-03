package mfs

const (
	PathSeperator = "/"
)

type FileSystem struct {
	rootDir *Dir
	pwd     *Dir
}

func New() *FileSystem {
	return &FileSystem{rootDir: &Dir{Name: "/"}}
}

func (fs *FileSystem) CreateFile(path string) *File {

}

func (fs *FileSystem) CreateDir(path string) {

}

func (fs *FileSystem) ChangeDir(path string) error {

}

type Dir struct {
	Name  string
	dirs  []Dir
	files []File
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
