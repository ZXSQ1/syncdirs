package copier

type Copier struct {
	SourceFiles []string
	DestFiles   []string
}

/*
description: creates a new Copier instance
arguments: no arguments
return: the Copier instance
*/
func NewCopier() Copier {
	return Copier{}
}
