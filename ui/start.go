package ui

import "github.com/ZXSQ1/syncdirs/app"

/*
description: synchronizes 2 directories
arguments:
  - sourceDir: the string path of the source directory
  - destDir: the string path of the destination directory
  - sourceFile: the source file string channel
  - destFile: the destination file string channel
  - err: the error message string channel
  - progress: the integer channel to contain the progress

return: no return
*/
func Synchronize(sourceDir, destDir string, sourceFile, destFile, err chan string, progress chan int) {
	lister := app.NewLister([]string{sourceDir, destDir})
	lister.List()

	differer := app.NewPathDiffererAB(sourceDir, destDir, lister.Get(sourceDir), lister.Get(destDir))
	differer.Differ()

	copier := app.NewCopier(differer.GetFound(), differer.GetMissing())
	copier.Copy(sourceFile, destFile, err, progress)
}
