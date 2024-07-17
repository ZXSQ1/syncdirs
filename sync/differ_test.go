package sync

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
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
		Name:    "table A",
		Entries: entriesTableA,
	}

	tableB := &DifferenceTable{
		Name:    "table B",
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

func TestDifferDirToCopy(t *testing.T) {
	pathDirA := "temp1"
	entriesDirA := []string{
		"something/good",
		"something/bad",
		"something/to_eat/pizza",
		"something/to_eat/tomato",
		"something/to_drink/coffee",
		"rust",
	}

	pathDirB := "temp2"
	entriesDirB := []string{
		"something/good",
		"something/bad",
		"something/to_eat/burger",
		"something/to_eat/mango",
		"something/to_eat/tomato",
		"something/to_drink/tea",
		"go",
	}

	t.Cleanup(func() {
		os.RemoveAll(pathDirA)
		os.RemoveAll(pathDirB)
	})

	for _, path := range entriesDirA {
		path = pathDirA + "/" + path
		os.MkdirAll(path, files.DirPerm)
	}

	for _, path := range entriesDirB {
		path = pathDirB + "/" + path
		os.MkdirAll(path, files.DirPerm)
	}

	syncData := make(chan *SyncData)
	go DifferDirToCopy(pathDirA, pathDirB, syncData)

	for {
		var data *SyncData

		data, ok := <-syncData
		if !ok {
			break
		}

		files.Copy(data.sourceFile, data.destFile)
	}

	contentsDirA, _ := files.ListDir(pathDirA, true)
	contentsDirB, _ := files.ListDir(pathDirB, true)

	if !slices.Equal(contentsDirA, contentsDirB) {
		t.Fail()
	}
}
