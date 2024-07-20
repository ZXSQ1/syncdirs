package app

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestPathDiffererAB(t *testing.T) {
	dir1, dir2 := "temp1", "temp2"
	set1, set2 := []string{"haskell", "bash", "python"}, []string{"haskell", "rust", "go"}

	t.Cleanup(func() {
		os.RemoveAll(dir1)
		os.RemoveAll(dir2)
	})

	for dir, entries := range map[string][]string{dir1: set1, dir2: set2} {
		os.MkdirAll(dir, files.DirPerm)

		for _, entry := range entries {
			path := dir + "/" + entry
			os.Create(path)
		}
	}

	differer := NewPathDiffererAB(dir1, dir2, set1, set2)
	differer.Differ()

	for _, path := range differer.GetFound() {
		if !slices.Contains([]string{
			"temp1/bash", "temp1/python",
			"temp2/rust", "temp2/go",
		}, path) {
			t.Fail()
		}
	}
}
