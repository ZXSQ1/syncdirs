package ui

import "github.com/ZXSQ1/syncdirs/app"

func Synchronize(sourceDir, destDir string, sourceFile, destFile, err chan string, progress chan int) {
	lister := app.NewLister([]string{sourceDir, destDir})
	lister.List()

	differer := app.NewPathDiffererAB(sourceDir, destDir, lister.Get(sourceDir), lister.Get(destDir))
	differer.Differ()

	copier := app.NewCopier(differer.GetFound(), differer.GetMissing())
	copier.Copy(sourceFile, destFile, err, progress)
}
