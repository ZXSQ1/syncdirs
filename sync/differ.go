package sync

import "github.com/ZXSQ1/syncdirs/files"

type DifferenceTable struct {
	Name string
	Entries []string
	Missing []string
}

/*
description: gets the difference between 2 string slices
arguments:
	- tableA: the first table to compare (of type *DifferenceTable)
	- tableB: the second table to compare (of type *DifferenceTable)
return: no return
*/
func Differ(tableA, tableB *DifferenceTable) {
	const (
		both = 0
		onlyTableA = 1
		onlyTableB = 2
	)

	var fullTable = map[string]int{}

	for _, entry := range tableA.Entries {
		fullTable[entry] = onlyTableA
	}

	for _, entry := range tableB.Entries {
		if _, ok := fullTable[entry]; ok {
			fullTable[entry] = both
		} else {
			fullTable[entry] = onlyTableB
		}
	}

	for key, val := range fullTable {
		if val == onlyTableA {
			tableB.Missing = append(tableB.Missing, key)
		} else if val == onlyTableB {
			tableA.Missing = append(tableA.Missing, key)
		}
	}
}

/*
description: gets the difference in the contents of directories
arguments:
	- dirA: the path to the first directory
	- dirB: the path to the second directory
	- source: a channel to transfer the source to be copied
	- dest: a channel to transfer the dest to be copied to
return: no return
*/
func DifferDirToCopy(dirA, dirB string, source, dest chan string) {
	entriesDirA, _ := files.ListDir(dirA, true)
	tableDirA := &DifferenceTable{
		Name: dirA,
		Entries: entriesDirA,
	}

	entriesDirB, _ := files.ListDir(dirB, true)
	tableDirB := &DifferenceTable{
		Name: dirB,
		Entries: entriesDirB,
	}

	Differ(tableDirA, tableDirB)

	defer close(source)
	defer close(dest)

	for _, missingPath := range tableDirA.Missing {
		sourcePath := tableDirB.Name + "/" + missingPath
		destPath := tableDirA.Name + "/" + missingPath

		source<-sourcePath
		dest<-destPath
	}

	for _, missingPath := range tableDirB.Missing {
		sourcePath := tableDirA.Name + "/" + missingPath
		destPath := tableDirB.Name + "/" + missingPath

		source<-sourcePath
		dest<-destPath
	}
}
