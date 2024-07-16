package files

import (
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
func GetFile(path string) (*os.File, error) {
	var file *os.File
	var err error = nil

	if !IsExist(filepath.Dir(path)) {
		os.MkdirAll(filepath.Dir(path), DirPerm)
	}

	if IsExist(path) {
		file, err = os.Open(path)
	} else {
		file, err = os.Create(path)
	}

	return file, err
}
