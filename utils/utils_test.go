package utils

import (
	"os"
	"path"
	"testing"
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

	os.MkdirAll(dirPath, 0644)

	if _, ok := ValidateDir(dirPath); !ok {
		t.Fail()
	}

	dirPath = dirPath + "1"

	if _, ok := ValidateDir(dirPath); ok {
		t.Fail()
	}

	os.RemoveAll(dirPath)

	if _, ok := ValidateDir(dirPath); ok {
		t.Fail()
	}
}
