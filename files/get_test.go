package files

import (
	"os"
	"testing"
)

func TestGetFile(t *testing.T) {
	temp := "temp"

	t.Cleanup(func() {
		os.Remove(temp)
	})

	file, err := GetFile(temp, FilePerm)
	defer file.Close()

	if err != nil {
		t.Fail()
	}

	if file.Name() != temp {
		t.Fail()
	}
}
