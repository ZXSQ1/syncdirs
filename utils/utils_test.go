package utils

import (
	"fmt"
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
	dirPath, ok := ValidateDir(dirPath)

	if !ok {
		t.Fail()
	}

	dirPath = dirPath + "1"
	_, ok = ValidateDir(dirPath)

	if ok {
		t.Fail()
	}

	os.RemoveAll(dirPath)
	_, ok = ValidateDir(dirPath)

	if ok {
		t.Fail()
	}

	fmt.Println(dirPath)
}
