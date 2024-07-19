package utils

import (
	"slices"
	"strings"
	"testing"
)

func TestDiffer(t *testing.T) {
	entriesTableA := []string{
		"something/good",
		"something/bad",
		"something/to_eat/pizza",
		"something/to_eat/tomato",
		"something/to_drink/coffee",
		"rust",
	}

	entriesTableB := []string{
		"something/good",
		"something/bad",
		"something/to_eat/burger",
		"something/to_eat/mango",
		"something/to_eat/tomato",
		"something/to_drink/tea",
		"go",
	}

	differenceSlice := Differ(entriesTableA, entriesTableB)

	for _, entry := range differenceSlice {
		if strings.HasPrefix(entry, "a:\t") {
			entry = entry[len("a:\t"):]

			if !slices.Contains([]string{
				"something/to_eat/burger",
				"something/to_eat/mango",
				"something/to_drink/tea",
				"go",
			}, entry) {
				t.Fail()
			}
		} else {
			entry = entry[len("b:\t"):]

			if !slices.Contains([]string{
				"something/to_eat/pizza",
				"something/to_drink/coffee",
				"rust",
			}, entry) {
				t.Fail()
			}
		}
	}
}
