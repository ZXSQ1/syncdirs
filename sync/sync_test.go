package sync

import (
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestSynchronize(t *testing.T) {
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
		os.MkdirAll(path, 0644)
	}

	for _, path := range entriesDirB {
		path = pathDirB + "/" + path
		os.MkdirAll(path, 0644)
	}

	data := make(chan *SyncData)
	Synchronize(pathDirA, pathDirB, data)

	go func() {
		for {
			if syncData, ok := <-data; ok {
				fmt.Println(syncData.sourceFile)
				fmt.Println(syncData.destFile)
				fmt.Println(syncData.err)
			} else {
				break
			}
		}
	}()

	contentsDirA, _ := files.ListDir(pathDirA, true)
	contentsDirB, _ := files.ListDir(pathDirB, true)

	for _, path := range contentsDirA {
		if !slices.Contains(contentsDirB, path) {
			t.Fail()
		}
	}
}

func TestSynchronizeMultiple(t *testing.T) {
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

	pathDirC := "temp3"
	entriesDirC := []string{
		"something/good",
		"something/bad",
		"something/horrible",
		"something/terrible",
		"something/to_do/code",
		"something/to_eat/mango",
		"something/to_do/write_tests",
		"python",
	}

	t.Cleanup(func() {
		for _, pathDir := range []string{pathDirA, pathDirB, pathDirC} {
			os.RemoveAll(pathDir)
		}
	})

	for dirPath, entries := range map[string][]string{
		pathDirA: entriesDirA,
		pathDirB: entriesDirB,
		pathDirC: entriesDirC} {

		for _, path := range entries {
			path = dirPath + "/" + path
			os.MkdirAll(path, 0644)
		}
	}

	data := make(chan *SyncData)
	SynchronizeMultiple([]string{pathDirA, pathDirB, pathDirC}, data)

	go func() {
		for {
			if syncData, ok := <-data; ok {
				fmt.Println(syncData.sourceFile)
				fmt.Println(syncData.destFile)
				fmt.Println(syncData.err)
			} else {
				break
			}
		}
	}()

	contentsDirA, _ := files.ListDir(pathDirA, true)
	contentsDirB, _ := files.ListDir(pathDirB, true)
	contentsDirC, _ := files.ListDir(pathDirC, true)

	for _, path := range contentsDirA {
		if !slices.Contains(contentsDirB, path) {
			t.Fail()
		}
	}

	for _, path := range contentsDirA {
		if !slices.Contains(contentsDirC, path) {
			t.Fail()
		}
	}
}
