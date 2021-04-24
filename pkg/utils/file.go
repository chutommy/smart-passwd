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

// FileName returns the name of the file without the path.
func (f *File) FileName() string {
	return f.Name + "." + f.Type
}

// FilePath returns the path of the file including its name.
func (f *File) FilePath() string {
	return filepath.Join(f.Path, f.FileName())
}
