package utils

import (
	"os"
	"path"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestError(t *testing.T) {
	errMessage := "error!"

	if Error(errMessage) != errCol.Sprint("E: ")+errMessage {
		t.Fail()
	}
}

func TestValidateDir(t *testing.T) {
	dirPath := "./.cache/syncdirs/testDir"

	t.Cleanup(func() {
		cacheDir := path.Clean(path.Dir(dirPath))
		os.RemoveAll(cacheDir)
	})

	os.MkdirAll(dirPath, files.DirPerm)
	dirPath, ok := ValidateDir(dirPath)

	if !ok {
		t.Fatal("if 1")
	}

	_, ok = ValidateDir(dirPath + "1")

	if ok {
		t.Fatal("if 2")
	}

	os.RemoveAll(dirPath)
	_, ok = ValidateDir(dirPath)

	if ok {
		t.Fatal("if 3")
	}
}
