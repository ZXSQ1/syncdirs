package utils

/*
description: gets the difference between 2 string slices
arguments:
  - a: the first slice to compare.
  - b: the second slice to compare.

return: a string slice of the missing strings prefixed by:
  - the string literal "a\t" if the string is missing from slice a
  - the string literal "b\t" if the string is missing from slice b
*/
func Differ(a []string, b []string) []string {
	const (
		both       = 0
		onlyTableA = 1
		onlyTableB = 2
	)

	var fullTable = map[string]int{}
	var fullSlice = []string{}

	for _, entry := range a {
		fullTable[entry] = onlyTableA
	}

	for _, entry := range b {
		if _, ok := fullTable[entry]; ok {
			fullTable[entry] = both
		} else {
			fullTable[entry] = onlyTableB
		}
	}

	for key, val := range fullTable {
		if val == onlyTableA {
			fullSlice = append(fullSlice, "b\t"+key)
		} else if val == onlyTableB {
			fullSlice = append(fullSlice, "a\t"+key)
		}
	}

	return fullSlice
}
