package ui

import (
	"github.com/ZXSQ1/syncdirs/app"
	"github.com/ZXSQ1/syncdirs/channels"
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
