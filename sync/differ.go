package sync

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
