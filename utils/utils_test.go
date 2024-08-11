package utils

import (
	"os"
	"path"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestValidateDir(t *testing.T) {
	dirPath := "./.cache/syncdirs/testDir"

	t.Cleanup(func() {
		cacheDir := path.Clean(path.Dir(dirPath))
		os.RemoveAll(cacheDir)
	})

	os.MkdirAll(dirPath, files.DirPerm)
	dirPath, ok := ValidateDir(dirPath)

	if !ok {
		t.Fail()
	}

	_, ok = ValidateDir(dirPath + "1")

	if ok {
		t.Fail()
	}

	os.RemoveAll(dirPath)
	_, ok = ValidateDir(dirPath)

	if ok {
		t.Fail()
	}
}
