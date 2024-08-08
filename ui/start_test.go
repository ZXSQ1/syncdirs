package ui

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/app"
)

func TestSynchronize(t *testing.T) {
	dirA := "temp1"
	dirAEntries := []string{
		dirA + "/java",
		dirA + "/bedrock",
		dirA + "/minecraft",
	}

	dirB := "temp2"
	dirBEntries := []string{
		dirB + "/bednorock",
		dirB + "/nojava",
		dirB + "/minecraft",
	}

	t.Cleanup(func() {
		for _, dir := range []string{dirA, dirB} {
			os.RemoveAll(dir)
		}
	})

	for _, entries := range [][]string{dirAEntries, dirBEntries} {
		for _, entry := range entries {
			fileObj, _ := os.Create(entry)
			fileObj.Close()
		}
	}

	Synchronize(dirA, dirB)

	lister := app.NewLister([]string{dirA, dirB})
	lister.List()

	if !slices.Equal(lister.Get(dirA), lister.Get(dirB)) {
		t.Fail()
	}
}
