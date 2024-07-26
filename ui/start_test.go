package ui

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestSynchronize(t *testing.T) {
	dirA := "temp1"
	dirB := "temp2"

	dirAEntries := []string{
		"nothing",
		"everything",
		"anything",
	}

	dirBEntries := []string{
		"good",
		"bad",
		"evil",
	}

	t.Cleanup(func() {
		os.RemoveAll(dirA)
		os.RemoveAll(dirB)
	})

	for _, path := range dirAEntries {
		path = dirA + "/" + path

		fileObj, _ := files.GetFile(path, files.FilePerm)
		fileObj.Close()
	}

	for _, path := range dirBEntries {
		path = dirB + "/" + path

		fileObj, _ := files.GetFile(path, files.FilePerm)
		fileObj.Close()
	}

	Synchronize(dirA, dirB, nil, nil, nil, nil)

	dirAContents, _ := files.ListDir(dirA, true)
	dirBContents, _ := files.ListDir(dirB, true)

	if !slices.Equal(dirAContents, dirBContents) {
		t.Fail()
	}
}
