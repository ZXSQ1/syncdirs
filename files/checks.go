package files

import (
	"os"
)

const FilePerm = 0644
const DirPerm = 0755

/*
description: checks if the file exists or not
arguments:
  - file: the string file path to check

return: a boolean that indicates whether the file exists or not
*/
func IsExist(file string) bool {
	_, err := os.Stat(file)

	return !os.IsNotExist(err)
}

/*
description: checks if the file is a file
arguments:
  - file: the string path to check

return:
  - bool: a boolean that indicates whether the file is a file or not
  - error: an error object
*/
func IsFile(file string) (bool, error) {
	stat, err := os.Stat(file)

	if err != nil {
		if err == os.ErrPermission {
			return false, nil
		}

		return false, err
	}

	return !stat.IsDir(), nil
}

/*
description: checks if the file is a directory
arguments:
  - file: the string path to check

return:
  - bool: a boolean that indicates whether the directory is a dir or not
  - error: an error object
*/
func IsDir(file string) (bool, error) {
	isFile, err := IsFile(file)

	if err != nil {
		return false, err
	}

	return !isFile, nil
}
