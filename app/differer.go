package app

import (
	"github.com/ZXSQ1/syncdirs/utils"
)

type MissingPath string
type FoundPath string

type PathDiffererAB struct {
	DirAName    string
	DirBName    string
	DirAEntries []string
	DirBEntries []string
	Difference  map[MissingPath]FoundPath
}

/*
description: gets an instance of the PathDiffererAB structure
arguments:
  - dirA: the path of the first directory
  - dirB: the path of the second directory

return: a PathDiffererAB instance
*/
func NewPathDiffererAB(dirA, dirB string, dirAEntries, dirBEntries []string) PathDiffererAB {
	return PathDiffererAB{
		DirAName:    dirA,
		DirBName:    dirB,
		DirAEntries: dirAEntries,
		DirBEntries: dirBEntries,
		Difference:  map[MissingPath]FoundPath{},
	}
}

/*
description: finds the difference between 2 directories
arguments: no arguments
return: no return
*/
func (pathDiffererAB *PathDiffererAB) Differ() {
	differences := utils.Differ(pathDiffererAB.DirAEntries, pathDiffererAB.DirBEntries)

	for _, difference := range differences {
		if difference[:3] == "a:\t" {
			difference = difference[3:]

			missingPath := MissingPath(pathDiffererAB.DirAName + "/" + difference)
			foundPath := FoundPath(pathDiffererAB.DirBName + "/" + difference)

			pathDiffererAB.Difference[missingPath] = foundPath
		} else {
			difference = difference[3:]

			missingPath := MissingPath(pathDiffererAB.DirBName + "/" + difference)
			foundPath := FoundPath(pathDiffererAB.DirAName + "/" + difference)

			pathDiffererAB.Difference[missingPath] = foundPath
		}
	}
}
