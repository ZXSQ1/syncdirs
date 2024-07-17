package sync

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
)

type SyncDataFile struct {
	SourceFile string // the source file path
	DestFile   string // the destination file path
}

type SyncDataDir struct {
	SourceDir         string // the source directory file path
	DestDir           string // the destination directory file path
	SourceDirEntryLen int    // the number of files in the source directory
	DestDirEntryLen   int    // the number of files in the destination directory
}

type SyncData struct {
	FileData *SyncDataFile
	DirData  *SyncDataDir
	Err      string
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

	fileData := make(chan *SyncDataFile)
	dirData := make(chan *SyncDataDir)

	go func() {
		DifferDirToCopy(dirA, dirB, fileData, dirData)
	}()

	for {
		fileDataStruct, fileDataOk := <-fileData
		dirDataStruct, dirDataOk := <-dirData

		if !fileDataOk || !dirDataOk {
			break
		}

		waitGroup.Add(1)

		go func() {
			data := SyncData{}
			err := files.Copy(fileDataStruct.SourceFile, fileDataStruct.DestFile)

			data.FileData = fileDataStruct
			data.DirData = dirDataStruct

			if err != nil {
				data.Err = "copy operation failed"
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
