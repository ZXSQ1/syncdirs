package sync

import (
	"slices"
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

	tableA := &DifferenceTable{
		Name: "table A",
		Entries: entriesTableA,
	}

	tableB := &DifferenceTable{
		Name: "table B",
		Entries: entriesTableB,
	}

	Differ(tableA, tableB)

	for _, entry := range tableA.Missing {
		if !slices.Contains([]string{
			"something/to_eat/burger",
			"something/to_eat/mango",
			"something/to_drink/tea",
			"go",
		}, entry) {
			t.Fail()
		}
	}
	
	for _, entry := range tableB.Missing {
		if !slices.Contains([]string{
			"something/to_eat/pizza",
			"something/to_drink/coffee",
			"rust",
		}, entry) {
			t.Fail()
		}
	}
}
