package files

import (
	"fmt"
	"os"
)

/*
description: lists the files in a directory recursively
arguments:
  - directory: the path of the directory to list the files of
  - trimBase: will trim the directory path out of the paths if specified
  - ex. temp/main/thing --> main/thing

return:
  - []string: a string slice containing the paths
  - error: an error if there is any
*/
func ListDir(directory string, trimBase bool) ([]string, error) {
	var result []string
	var err error = nil

	if !IsExist(directory) {
		return result, fmt.Errorf("directory not found: %s", directory)
	} else if isDir, _ := IsDir(directory); !isDir {
		return result, fmt.Errorf("file not a directory: %s", directory)
	}

	dirEntries, err := os.ReadDir(directory)

	for _, entry := range dirEntries {
		entryPath := directory + "/" + entry.Name()

		if isDir, _ := IsDir(entryPath); isDir {
			recursiveEntries, _ := ListDir(entryPath, false)
			result = append(result, recursiveEntries...)
		} else {
			result = append(result, entryPath)
		}
	}

	if trimBase {
		for index, path := range result {
			result[index] = path[len(directory)+1:]
		}
	}

	return result, err
}
