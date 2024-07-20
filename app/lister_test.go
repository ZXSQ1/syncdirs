package app

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestLister(t *testing.T) {
	dirs := []string{"temp1", "temp2"}

	t.Cleanup(func() {
		for _, dir := range dirs {
			os.RemoveAll(dir)
		}
	})

	for _, path := range []string{
		dirs[0] + "/main/thing",
		dirs[0] + "/jars/thing",
		dirs[0] + "/what/thing",
		dirs[1] + "/you/dumb",
		dirs[1] + "/no/you",
		dirs[1] + "/no/YOU",
	} {
		os.MkdirAll(path, files.DirPerm)
	}

	lister := NewLister(dirs)
	lister.List()

	for _, path := range lister.Get(dirs[0]) {
		if !slices.Contains([]string{
			"main/thing", "jars/thing", "what/thing",
		}, path) {
			t.Fail()
		}
	}

	for _, path := range lister.Get(dirs[1]) {
		if !slices.Contains([]string{
			"you/dumb", "no/you", "no/YOU",
		}, path) {
			t.Fail()
		}
	}
}
