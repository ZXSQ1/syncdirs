package files

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/utils"
)

func TestListDir(t *testing.T) {
	tempDir := "temp"
	pathsToCreate := []string{
		tempDir + "/name/first",
		tempDir + "/name/second",
		tempDir + "/address",
		tempDir + "/gender",
	}

	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	for _, pathToCreate := range pathsToCreate {
		fileObj, _ := GetFile(pathToCreate)
		fileObj.Close()
	}

	entries, err := ListDir(tempDir, false)

	if err != nil {
		utils.PrintError(err.Error())
	}

	for _, entry := range entries {
		if !slices.Contains(pathsToCreate, entry) {
			t.Fail()
		}
	}

	entries, err = ListDir(tempDir, true)

	if err != nil {
		utils.PrintError(err.Error())
	}

	for index, pathToCreate := range pathsToCreate {
		pathsToCreate[index] = pathToCreate[len(tempDir)+1:]
	}

	for _, entry := range entries {
		if !slices.Contains(pathsToCreate, entry) {
			t.Fail()
		}
	}
}
