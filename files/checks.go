package files

import "os"

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

return: a boolean that indicates whether the file is a file or not
*/
func IsFile(file string) bool {
	stat, _ := os.Stat(file)

	return !stat.IsDir()
}

/*
description: checks if the file is a directory
arguments:
  - file: the string path to check

return: a boolean that indicates whether the directory is a dir or not
*/
func IsDir(file string) bool {
	return !IsFile(file)
}
