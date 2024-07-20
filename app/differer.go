package app

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
	}
}
