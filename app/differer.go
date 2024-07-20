package app

type MissingPath string
type FoundPath string

type PathDiffererAB struct {
	DirAName   string
	DirBName   string
	Difference map[MissingPath]FoundPath
}
