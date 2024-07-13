package files

import (
	"os"
	"testing"
)

func TestIsExist(t *testing.T) {
	tempFile := "tempFile"

	os.Create(tempFile)

	if !IsExist(tempFile) {
		t.Fail()
	}

	os.Remove(tempFile)

	if IsExist(tempFile) {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	tempPath := "temp"

	t.Cleanup(func() {
		os.RemoveAll(tempPath)
	})

	os.Create(tempPath)

	if !IsFile(tempPath) {
		t.Fail()
	}

	os.Remove(tempPath)
	os.Mkdir(tempPath, 0644)

	if IsFile(tempPath) {
		t.Fail()
	}
}
