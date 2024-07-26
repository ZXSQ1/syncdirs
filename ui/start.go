package ui

import (
	"fmt"
	"os"

	"github.com/ZXSQ1/syncdirs/app"
	"github.com/ZXSQ1/syncdirs/channels"
	"github.com/ZXSQ1/syncdirs/utils"
)

/*
description: synchronizes 2 directories
arguments:
  - sourceDir: the string path of the source directory
  - destDir: the string path of the destination directory
  - sourceFile: the source file string channel
  - destFile: the destination file string channel
  - err: the error message string channel
  - progress: the float32 channel to contain the progress

return: no return
*/
func Synchronize(sourceDir, destDir string, sourceFile, destFile, err chan string, progress chan float32) {
	lister := app.NewLister([]string{sourceDir, destDir})
	lister.List()

	differer := app.NewPathDiffererAB(sourceDir, destDir, lister.Get(sourceDir), lister.Get(destDir))
	differer.Differ()

	intProgress := make(chan int)
	copier := app.NewCopier(differer.GetFound(), differer.GetMissing())
	copier.Copy(sourceFile, destFile, err, intProgress)

	progressVal := float32((channels.Unfeed(intProgress)).(int) * 100 / len(lister.Get(sourceDir)))
	channels.Feed(progress, progressVal)
	channels.Close(progress)
}

func Start() {
	dirs := Handle(os.Args)

	sourceFile, destFile, err, progress := make(chan string), make(chan string), make(chan string), make(chan float32)

	go func() {
		for {
			sourcePath := channels.Unfeed(sourceFile).(string)
			destPath := channels.Unfeed(destFile).(string)
			progressIndication := channels.Unfeed(progress).(float32)

			fmt.Printf("%s -> %s (%.2f%%)\n", sourcePath, destPath, progressIndication)
		}
	}()

	go func() {
		for {
			errMsg := channels.Unfeed(err).(string)
			utils.PrintError(errMsg)
		}
	}()

	switch len(dirs) {
	case 0, 1:
		utils.PrintError("not enough directories provided")
		Help()
	case 2:
		Synchronize(dirs[0], dirs[1], sourceFile, destFile, err, progress)
	default:
		for i := 0; i < 2; i++ {
			centralDir := dirs[0]

			for _, syncDir := range dirs[1:] {
				Synchronize(centralDir, syncDir, sourceFile, destFile, err, progress)
			}
		}
	}

}
