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
