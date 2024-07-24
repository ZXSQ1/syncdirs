package app

import (
	"os"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestCopier(t *testing.T) {
	sourceFiles := []string{
		"temp1/java",
		"temp1/rust",
		"temp1/python",
	}

	destFiles := []string{
		"temp1/haskell",
		"temp1/go",
		"temp1/ruby",
	}

	t.Cleanup(func() {
		os.RemoveAll("temp1")
	})

	copier := NewCopier(sourceFiles, destFiles)
	copier.Copy(nil, nil, nil, nil)

	for _, destFile := range destFiles {
		if !files.IsExist(destFile) {
			t.Fail()
		}
	}
}
