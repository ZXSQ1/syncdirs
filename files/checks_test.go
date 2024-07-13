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
