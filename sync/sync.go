package sync

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
	"github.com/ZXSQ1/syncdirs/utils"
)

type SyncData struct {
	sourceFile string // the source file path
	destFile   string // the destination file path
	err        string // the error string if there is
}

/*
description: synchronizes 2 directories
arguments:
  - dirA: the string path to the first directory
  - dirB: the string path to the second directory
  - syncData: a channel of the *SyncData type

return: no return
*/
func Synchronize(dirA, dirB string, syncData chan *SyncData) {
	var waitGroup = &sync.WaitGroup{}

	source := make(chan string)
	dest := make(chan string)

	go func() {
		DifferDirToCopy(dirA, dirB, source, dest)
	}()

	for {
		sourcePath, sourceOk := <-source
		destPath, destOk := <-dest

		if !sourceOk || !destOk {
			break
		}

		waitGroup.Add(1)

		go func() {
			data := SyncData{}
			err := files.Copy(sourcePath, destPath)

			data.sourceFile = sourcePath
			data.destFile = destPath

			if err != nil {
				data.err = utils.Error("copy operation failed")
			}

			syncData <- &data
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}

/*
description: synchronizes multiple directories
arguments:
  - dirs: the string slice containing the directories
  - syncData: the channel of the *SyncData type

return: no return
*/
func SynchronizeMultiple(dirs []string, syncData chan *SyncData) {
	switch len(dirs) {
	case 0, 1:
		return
	case 2:
		Synchronize(dirs[0], dirs[1], syncData)
		return
	}

	centralDir := dirs[0]
	waitGroup := &sync.WaitGroup{}

	for _, dir := range dirs[1:] {
		waitGroup.Add(1)

		go func() {
			Synchronize(centralDir, dir, syncData)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
