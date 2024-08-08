package ui

import (
	"fmt"
	"os"
	"sync"

	"github.com/ZXSQ1/syncdirs/app"
	"github.com/ZXSQ1/syncdirs/utils"
)

func Synchronize(sourceDir, destDir string) {
	lister := app.NewLister([]string{sourceDir, destDir})
	lister.List()

	differer := app.NewPathDiffererAB(sourceDir, destDir, lister.Get(sourceDir), lister.Get(destDir))
	differer.Differ()

	copier := app.NewCopier([]string{}, []string{})

	for missing, found := range differer.Difference {
		copier.Add([]string{string(found)}, []string{string(missing)})
	}

	copier.Copy(func(cd app.CopierData) {
		if cd.Err != nil {
			utils.PrintError("%s. skipping...\n", cd.Err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "%-70s -> %-90s (%d left)\n",
				cd.SourceFile, cd.DestFile, len(differer.Difference)-cd.CopiedFiles)
			os.Stdout.Sync()
		}
	})
}

func SynchronizeMultiple(dirs []string) error {
	switch len(dirs) {
	case 0, 1:
		return fmt.Errorf("not enough directories provided")
	case 2:
		Synchronize(dirs[0], dirs[1])
		return nil
	}

	var waitGroup = &sync.WaitGroup{}
	centralDir := dirs[0]

	for _, dir := range dirs[1:] {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			Synchronize(centralDir, dir)
		}()
	}

	waitGroup.Wait()

	return nil
}

func Start() {
	dirs := Handle(os.Args)
	SynchronizeMultiple(dirs)
}
