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
			err := files.Copy(sourcePath, destPath)

			if err != nil {
				utils.PrintError("copy operation file: %s failed", sourcePath)
			}
			
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}


/*
description: synchronizes multiple directories
arguments:
	- dirs: the string slice containing the directories
return: no return
*/
func SynchronizeMultiple(dirs []string) {
	switch len(dirs) {
	case 0, 1:
		return
	case 2:
		Synchronize(dirs[0], dirs[1])
		return
	}

	centralDir := dirs[0]
	waitGroup := &sync.WaitGroup{}

	for _, dir := range dirs[1:] {
		waitGroup.Add(1)

		go func() {
			Synchronize(centralDir, dir)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
