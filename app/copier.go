package app

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
)

type Copier struct {
	SourceFiles []string
	DestFiles   []string
}

type CopierData struct {
	SourceFile  string
	DestFile    string
	CopiedFiles int
	Err         error
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
  - infoFn: the function to print the info and takes a CopierData sturcture.
  - jobs: the uint value that specifies the number of jobs

return: no return
*/
func (copier *Copier) Copy(infoFn func(CopierData), jobs uint) {
	var waitGroup = &sync.WaitGroup{}
	var waitGroupCounter uint
	var progressMutex = &sync.Mutex{}

	var progress int

	for index := range copier.SourceFiles {
		sourcePath := copier.SourceFiles[index]
		destPath := copier.DestFiles[index]

		waitGroup.Add(1)
		waitGroupCounter++

		go func(src, dst string) {
			defer waitGroup.Done()

			errVal := files.Copy(src, dst)
			progressMutex.Lock()
			progress += 1

			infoFn(CopierData{
				src,
				dst,
				progress,
				errVal,
			})

			progressMutex.Unlock()
		}(sourcePath, destPath)

		if jobs != 0 {
			if waitGroupCounter == jobs {
				waitGroupCounter = 0
				waitGroup.Wait()
			}
		}
	}

	waitGroup.Wait()
}
