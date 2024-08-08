package ui

import (
	"fmt"
	"os"

	"github.com/ZXSQ1/syncdirs/utils"
)

/*
description: prints the help message and exits
*/
func Help() {
	helpMsg := `usage: ` + os.Args[0] + ` <directories>
ex: syncdirs ~/Downloads ~/Desktop ~/Temp
`

	fmt.Print(helpMsg)
	os.Exit(1)
}

/*
description: the function to handle the arguments and clean them
arguments:
  - fullArgs: the slice of the arguments that enter the program (including the program name)

return: the slice of cleaned strings
*/
func Handle(fullArgs []string) []string {
	args := fullArgs

	if len(args) < 3 {
		utils.PrintError("must specify at least 2 directories after\n")
		Help()
	}

	args = args[1:]

	for index, dir := range args {
		if dir, ok := utils.ValidateDir(dir); ok {
			args[index] = dir
		} else {
			utils.PrintError("invalid directory path\n")
			Help()
		}
	}

	return args
}
