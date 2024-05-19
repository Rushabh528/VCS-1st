package JSN

import (
	"os"
	"path/filepath"
)

func Ls() ([]string, error) {
	// returns working dir in string
	wd, err := os.Getwd()
	// if err != nil {
	// 	return nil, err
	// }
	// Opens dir and return *file type
	// file is a struct defined in os pkg and various methods defined
	path := filepath.Join(wd, "files")

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	// ReadDir and return []fs.DirEntry
	// fs.DirEntry is an interface
	// To implement fs.DirEntry you need to implement 4 methods
	// Name() String, IsDir() bool, Type() FileMode, Info (FileInfo, error)
	files, err0 := dir.ReadDir(0)
	if err0 != nil {
		return nil, err
	}
	list := []string{} // slice of names of file
	for _, file := range files {
		if !file.IsDir() {
			list = append(list, file.Name())
		}
	}
	return list, nil
}
