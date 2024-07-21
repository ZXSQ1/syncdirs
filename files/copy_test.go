package files

import (
	"os"
	"slices"
	"testing"
)

func TestCopy(t *testing.T) {
	sourcePath := "tempSrc"
	destPath := "tempDst"
	sourceData := []byte("This is a test.")

	t.Cleanup(func() {
		os.Remove(sourcePath)
		os.Remove(destPath)
	})

	sourceObj, _ := GetFile(sourcePath, FilePerm)
	sourceObj.Write(sourceData)
	sourceObj.Close()

	err := Copy(sourcePath, destPath)

	if err != nil {
		t.Fail()
	}

	buffer := make([]byte, len(sourceData))

	destObj, _ := GetFile(destPath, FilePerm)
	destObj.Read(buffer)

	if !slices.Equal(buffer, sourceData) {
		t.Fail()
	}
}
