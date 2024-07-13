package files

import (
	"os"
	"slices"
	"testing"
)

func TestListDir(t *testing.T) {
	tempDir := "temp"
	pathsToCreate := []string{
		tempDir+"/name/first",
		tempDir+"/name/second",
		tempDir+"/address",
		tempDir+"/gender",
	}

	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	for _, pathToCreate := range pathsToCreate {
		fileObj, _ := GetFile(pathToCreate)
		fileObj.Close()
	}

	entries, err := ListDir(tempDir)

	if err != nil {
		t.Fail()
	}

	for _, entry := range entries {
		if !slices.Contains(pathsToCreate, entry) {
			t.Fail()
		}
	}
}
