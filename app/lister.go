package app

import (
	"sync"

	"github.com/ZXSQ1/syncdirs/files"
)

type Lister struct {
	DirNames   []string
	DirEntries map[string][]string
}

/*
description: returns a new lister instance
arguments:
  - dirNames: the string slice containing the directory names

return: a Lister structure
*/
func NewLister(dirNames []string) Lister {
	return Lister{
		DirNames: dirNames,
	}
}

/*
description: adds a directory path to the slice of directories
arguments:
  - dirName: the directory path to add

return: no return
*/
func (lister *Lister) Add(dirName string) {
	lister.DirNames = append(lister.DirNames, dirName)
}

/*
description: lists the contents of the directories
arguments: no arguments
return: no return
*/
func (lister *Lister) List() {
	var waitGroup = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}
	defer waitGroup.Wait()

	for _, dir := range lister.DirNames {
		if _, ok := lister.DirEntries[dir]; !ok {
			waitGroup.Add(1)

			go func() {
				defer waitGroup.Done()

				temp, _ := files.ListDir(dir, true)

				mutex.Lock()
				lister.DirEntries[dir] = temp
				mutex.Unlock()
			}()
		}
	}
}

/*
description: gets the directory entries given the name
arguments:
  - dirName: the string path referring to the directory path

return: the string slice of the paths
*/
func (lister *Lister) Get(dirName string) []string {
	if val, ok := lister.DirEntries[dirName]; ok {
		return val
	} else {
		return nil
	}
}
