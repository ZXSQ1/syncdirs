package sync

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
)

/*
description: synchronizes 2 directories
arguments:
	- dirA: the string path to the first directory
	- dirB: the string path to the second directory
return: no return
*/
func Synchronize(dirA, dirB string) {
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
			files.Copy(sourcePath, destPath)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
