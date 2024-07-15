package utils

import (
	"testing"

	"github.com/fatih/color"
)

func TestError(t *testing.T) {
	errMessage := "error!"

	if Error(errMessage) != color.RedString("E: " + errMessage) {
		t.Fail()
	}
}
