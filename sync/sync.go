package sync

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
	"github.com/ZXSQ1/syncdirs/utils"
)

/*
description: synchronizes 2 directories
arguments:
	- dirA: the string path to the first directory
	- dirB: the string path to the second directory
	- currentFile: the string channel where the currentFile is signaled
	- errChan: the string channel to signal error texts
return: no return
*/
func Synchronize(dirA, dirB string, currentFile, errChan chan string) {
	var waitGroup = &sync.WaitGroup{}

	source := make(chan string)
	dest := make(chan string)

	go func() {
		DifferDirToCopy(dirA, dirB, source, dest)
	}()
	
	for {
		sourcePath, sourceOk := <- source
		destPath, destOk := <- dest

		if !sourceOk || !destOk {
			break
		}
		
		waitGroup.Add(1)

		go func() { 
			err := files.Copy(sourcePath, destPath)

			if err != nil {
				errChan<-utils.Error("copy operation failed")
			}

			currentFile<-sourcePath
			
			close(currentFile)
			close(errChan)

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}


/*
description: synchronizes multiple directories
arguments:
	- dirs: the string slice containing the directories
	- currentFile: the string channel to carry the currentFile
	- errChan: the string channel that holds the error text
return: no return
*/
func SynchronizeMultiple(dirs []string, currentFile, errChan chan string) {
	switch len(dirs) {
	case 0, 1:
		return
	case 2:
		Synchronize(dirs[0], dirs[1], currentFile, errChan)
		return
	}

	centralDir := dirs[0]
	waitGroup := &sync.WaitGroup{}

	for _, dir := range dirs[1:] {
		waitGroup.Add(1)

		go func() {
			Synchronize(centralDir, dir, currentFile, errChan)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
