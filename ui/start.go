package ui

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
	"github.com/ZXSQ1/syncdirs/utils"
)

type SyncData struct {
	centralDirEntries []string
	destDirEntries    []string
	missingCentralDir []string
	missingDestDir    []string
	differences       []string
	sourceFile        chan string
	sourceDir         chan string
	destFile          chan string
	destDir           chan string
	progress          chan int
}

func (data SyncData) ListDir(centralDir, destDir string) SyncData {
	var waitGroup = sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()

		data.destDirEntries, _ = files.ListDir(destDir, true)
	}()

	if slices.Equal(data.centralDirEntries, nil) {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			data.centralDirEntries, _ = files.ListDir(centralDir, true)
		}()
	}

	waitGroup.Wait()

	return data
}

func (data SyncData) Differ(centralDirName, destDirName string) SyncData {
	differences := utils.Differ(data.centralDirEntries, data.destDirEntries)

	for index, difference := range differences {
		if strings.HasPrefix(difference, "a:\t") {
			differences[index] = strings.ReplaceAll(difference, "a:\t", centralDirName+"//")
		} else {
			differences[index] = strings.ReplaceAll(difference, "b:\t", destDirName+"//")
		}
	}

	data.differences = differences

	return data
}

func (data SyncData) Copy(centralDirName, destDirName string) SyncData {
	var waitGroup = sync.WaitGroup{}
	var mutex = sync.Mutex{}

	for _, path := range data.missingDestDir {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			path := strings.Split(path, "//")[1]
			destDir := destDirName
			sourceDir := strings.Split(path, "//")[0]

			data.destDir <- destDir
			data.sourceDir <- sourceDir
			data.destFile <- destDir + "/" + path
			data.sourceDir <- sourceDir + "/" + path

			err := files.Copy(sourceDir+"/"+path, destDir+"/"+path)

			if err != nil {
				utils.PrintError(err.Error())
			} else {
				mutex.Lock()
				data.progress <- <-data.progress + 1
				mutex.Unlock()
			}
		}()
	}

	for _, path := range data.missingCentralDir {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			path := strings.Split(path, "//")[1]
			destDir := centralDirName
			sourceDir := strings.Split(path, "//")[0]

			data.destDir <- destDir
			data.sourceDir <- sourceDir
			data.destFile <- destDir + "/" + path
			data.sourceDir <- sourceDir + "/" + path

			err := files.Copy(sourceDir+"/"+path, destDir+"/"+path)

			if err != nil {
				utils.PrintError(err.Error())
			} else {
				mutex.Lock()
				data.progress <- <-data.progress + 1
				mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()

	return data
}

func Synchronize(directories []string, sourceFile, destFile, sourceDir, destDir chan string, progress chan float32) {
	var centralDir = directories[0]
	//	var mutex = sync.Mutex{}
	var waitGroup = sync.WaitGroup{}
	var data = SyncData{}

	defer func() {
		close(data.sourceFile)
		close(data.sourceDir)
		close(data.destFile)
		close(data.destDir)
		close(data.progress)

		close(sourceFile)
		close(sourceDir)
		close(destFile)
		close(destDir)
		close(progress)
	}()

	for i := 0; i < 2; i++ {
		for _, syncDir := range directories[1:] {
			// Listing Directory Contents
			data = data.ListDir(centralDir, syncDir)
			fmt.Println("done")

			// Differing Directories
			data = data.Differ(centralDir, syncDir)
			fmt.Println("done")

			// Copying Directories
			waitGroup.Add(1)

			go func() {
				defer waitGroup.Done()

				data = data.Copy(centralDir, syncDir)
			}()

			if val, ok := <-data.sourceFile; ok {
				sourceFile <- val
			}

			if val, ok := <-data.sourceDir; ok {
				sourceDir <- val
			}

			if val, ok := <-data.destFile; ok {
				sourceFile <- val
			}

			if val, ok := <-data.destDir; ok {
				destDir <- val
			}

			if val, ok := <-data.progress; ok {
				progress <- float32(val) * 100 / float32(len(data.differences))
			}

			waitGroup.Wait()
		}
	}
}
