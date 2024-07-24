package app

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
)

type Copier struct {
	SourceFiles []string
	DestFiles   []string
}

/*
description: creates a new Copier instance
arguments: no arguments
return: the Copier instance
*/
func NewCopier(sourceFiles, destFiles []string) Copier {
	if len(sourceFiles) != len(destFiles) {
		return Copier{}
	}

	return Copier{sourceFiles, destFiles}
}

/*
description: adds a source destination file pair
arguments:
  - sourceFiles: the source file path to add
  - destFiles: the dest file path to add

return: no return
*/
func (copier *Copier) Add(sourceFiles, destFiles []string) {
	if len(sourceFiles) != len(destFiles) {
		return
	}

	copier.SourceFiles = append(copier.SourceFiles, sourceFiles...)
	copier.DestFiles = append(copier.DestFiles, destFiles...)
}

/*
description: copies the sources to their destinations
arguments:
  - sourceFile: the string channel to carry the source file
  - destFile: the string channel to carry the destination file
  - err: the string channel to carry the error file
  - progress: the int channel to carry the current progress
*/
func (copier *Copier) Copy(sourceFile, destFile, err chan string, progress chan int) {
	var waitGroup = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	defer func() {
		close(sourceFile)
		close(destFile)
		close(err)
		close(progress)
	}()

	progress <- 0

	for index, _ := range copier.SourceFiles {
		sourcePath := copier.SourceFiles[index]
		destPath := copier.DestFiles[index]

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			errVal := files.Copy(sourcePath, destPath)

			if errVal != nil {
				err <- errVal.Error()
			} else {
				mutex.Lock()
				progress <- <-progress + 1
				mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()
}
