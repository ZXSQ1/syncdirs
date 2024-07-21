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
		fileObj, _ := GetFile(pathToCreate, FilePerm)
		fileObj.Close()
	}

	entries, err := ListDir(tempDir, false)

	if err != nil {
		utils.PrintError(err.Error())
	}

	if !slices.Equal([]string{"temp/address", "temp/gender", "temp/name/first", "temp/name/second"}, entries) {
		t.Fail()
	}

	entries, err = ListDir(tempDir, true)

	if err != nil {
		utils.PrintError(err.Error())
	}

	for index, pathToCreate := range pathsToCreate {
		pathsToCreate[index] = pathToCreate[len(tempDir)+1:]
	}

	if !slices.Equal([]string{"address", "gender", "name/first", "name/second"}, entries) {
		t.Fail()
	}
}
