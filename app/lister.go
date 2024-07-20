package app

type Lister struct {
	DirNames   []string
	DirEntries map[string]string
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
