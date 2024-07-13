package files

import (
	"fmt"
	"os"
)

/*
description: lists the files in a directory recursively
arguments:
	- directory: the path of the directory to list the files of
return:
	- []string: a string slice containing the paths
	- error: an error if there is any
*/
func ListDir(directory string) ([]string, error) {
	var result []string
	var err error = nil

	if !IsExist(directory) {
		return result, fmt.Errorf("directory not found: %s", directory)
	} else if !IsDir(directory) {
		return result, fmt.Errorf("file not a directory: %s", directory)
	}

	dirEntries, err := os.ReadDir(directory);
	
	for _, entry := range dirEntries {
		entryPath := directory + "/" + entry.Name()

		if IsDir(entryPath) {
			recursiveEntries, _ := ListDir(entryPath)
			result = append(result, recursiveEntries...)
		} else {
			result = append(result, entryPath)
		}
	}

	return result, err
}
