package main

import (
	"os"

	"github.com/ZXSQ1/syncdirs/ui"
)

func main() {
	ui.Dirs := ui.Handle(os.Args)
}
