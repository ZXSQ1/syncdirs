package utils

import (
	"testing"
)

func TestError(t *testing.T) {
	errMessage := "error!"

	if Error(errMessage) != errCol.Sprint("E: ") + errMessage {
		t.Fail()
	}
}
