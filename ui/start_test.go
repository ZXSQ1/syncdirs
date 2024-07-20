package ui

import (
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/syncdirs/files"
)

func TestSynchronize(t *testing.T) {
	dirA, dirB, dirC := "temp1", "temp2", "temp3"

	entriesDirA := []string{"go", "rust", "c", "cpp"}
	entriesDirB := []string{"python", "ruby", "bash"}
	entriesDirC := []string{"haskell", "ruby", "rust", "bash"}

	t.Cleanup(func() {
		for _, dir := range []string{dirA, dirB, dirC} {
			os.RemoveAll(dir)
		}
	})

	for dir, entries := range map[string][]string{
		dirA: entriesDirA, dirB: entriesDirB, dirC: entriesDirC,
	} {
		for _, path := range entries {
			path = dir + "/" + path
			os.MkdirAll(path, files.DirPerm)
		}
	}

	sourceDir, destDir, sourceFile, destFile, progress := make(chan string), make(chan string), make(chan string), make(chan string), make(chan float32)
	Synchronize([]string{dirA, dirB, dirC}, sourceFile, destFile, sourceDir, destDir, progress)

	go func() {
		for {
			fmt.Printf("%s %s %s %s %f", <-sourceFile, <-sourceDir, <-destFile, <-destDir, <-progress)

			if _, ok := <-progress; !ok {
				break
			}
		}
	}()

	contentsDirA, _ := files.ListDir(dirA, true)
	contentsDirB, _ := files.ListDir(dirB, true)
	contentsDirC, _ := files.ListDir(dirC, true)

	for _, contents := range map[string][]string{
		dirA: contentsDirA, dirB: contentsDirB,
	} {
		if !slices.Equal(contents, contentsDirC) {
			t.Fail()
		}
	}
}
