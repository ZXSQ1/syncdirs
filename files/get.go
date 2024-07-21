package files

import (
	"io/fs"
	"os"
	"path/filepath"
)

/*
description: gets a file object
arguments:
  - path: the string path to the file

return:
  - *os.File: the file object
  - error: an error
*/
func GetFile(path string, perm fs.FileMode) (*os.File, error) {
	var file *os.File
	var err error = nil

	if !IsExist(filepath.Dir(path)) {
		os.MkdirAll(filepath.Dir(path), DirPerm)
	}

	if IsExist(path) {
		stat, _ := os.Stat(path)
		perm = stat.Mode()

		file, err = os.OpenFile(path, os.O_RDWR, perm)
	} else {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, fs.FileMode(perm))
	}

	return file, err
}
