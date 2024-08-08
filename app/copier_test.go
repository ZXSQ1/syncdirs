package app

import (
	"fmt"
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

	for _, path := range sourceFiles {
		fileObj, _ := files.GetFile(path, files.FilePerm)
		fileObj.Close()
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
	copier.Copy(func(data CopierData) {
		fmt.Println(data)
	})

	for _, destFile := range destFiles {
		if !files.IsExist(destFile) {
			t.Fail()
		}
	}
}
