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

/*
description: adds a source destination file pair
arguments:
  - sourceFile: the source file path to add
  - destFile: the dest file path to add

return: no return
*/
func (copier *Copier) Add(sourceFile, destFile string) {
	copier.SourceFiles = append(copier.SourceFiles, sourceFile)
	copier.DestFiles = append(copier.DestFiles, destFile)
}
