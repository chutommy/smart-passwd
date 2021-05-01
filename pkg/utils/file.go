package utils

import (
	"path/filepath"
)

// File represents metadata of a file.
type File struct {
	Name string
	Type string
	Path string
}

// NewFile constructs a new File.
func NewFile(path string, name string, fType string) *File {
	return &File{
		Name: name,
		Type: fType,
		Path: path,
	}
}

// FileName returns the name of the file without the path.
func (f *File) FileName() string {
	return f.Name + "." + f.Type
}

// FilePath returns the path of the file including its name.
func (f *File) FilePath() string {
	return filepath.Join(f.Path, f.FileName())
}
