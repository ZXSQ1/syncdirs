package sync

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestSynchronize(t *testing.T) {
	pathDirA := "temp1"
	entriesDirA := []string{
		"something/good",
		"something/bad",
		"something/to_eat/pizza",
		"something/to_eat/tomato",
		"something/to_drink/coffee",
		"rust",
	}

	pathDirB := "temp2"
	entriesDirB := []string{
		"something/good",
		"something/bad",
		"something/to_eat/burger",
		"something/to_eat/mango",
		"something/to_eat/tomato",
		"something/to_drink/tea",
		"go",
	}

	t.Cleanup(func() {
		os.RemoveAll(pathDirA)
		os.RemoveAll(pathDirB)
	})

	for _, path := range entriesDirA {
		path = pathDirA + "/" + path
		os.MkdirAll(path, 0644)
	}

	for _, path := range entriesDirB {
		path = pathDirB + "/" + path
		os.MkdirAll(path, 0644)
	}

	Synchronize(pathDirA, pathDirB)

	contentsDirA, _ := files.ListDir(pathDirA, true)
	contentsDirB, _ := files.ListDir(pathDirB, true)

	for _, path := range contentsDirA {
		if !slices.Contains(contentsDirB, path) {
			t.Fail()
		}
	}
}
