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

/*
description: adds a directory path to the slice of directories
arguments:
  - dirName: the directory path to add

return: no return
*/
func (lister *Lister) Add(dirName string) {
	lister.DirNames = append(lister.DirNames, dirName)
}
