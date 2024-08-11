package ui

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ZXSQ1/syncdirs/utils"
)

/*
description: prints the help message and exits
*/
func Help() {
	helpMsg := `usage: ` + os.Args[0] + ` <directories>
ex: syncdirs ~/Downloads ~/Desktop ~/Temp
options:

-h, --help
	for help
-j, --jobs <uint>
	sets the number of jobs for copying (with 0 meaning infinite)
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
	args := fullArgs[1:]

	if len(fullArgs) < 2 {
		utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
		os.Exit(1)
	}

	filteredArgs := []string{}
	for index, option := range args {
		if !strings.HasPrefix(option, "-") {
			continue
		}

		optionIndices := []int{}

		switch option {
		case "-j", "--jobs":
			if index+1 == len(args) {
				utils.PrintError("must specify uint after option\n")
				utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
			}

			optionValue := args[index+1]
			num, err := strconv.Atoi(optionValue)

			if err != nil {
				utils.PrintError("must specify uint after option\n")
				utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
			} else if num < 0 {
				utils.PrintError("must specify uint after option\n")
				utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
			}

			Jobs = uint(num)
			optionIndices = append(optionIndices, []int{index, index + 1}...)

		case "-h", "--help":
			Help()
		default:
			utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
		}

		for index, value := range args {
			if slices.Contains(optionIndices, index) {
				continue
			}

			filteredArgs = append(filteredArgs, value)
		}
	}

	args = filteredArgs

	if len(args) < 2 {
		utils.PrintError("must specify at least 2 directories after\n")
		utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
	}

	for index, dir := range args {
		if dir, ok := utils.ValidateDir(dir); ok {
			args[index] = dir
		} else {
			utils.PrintError("invalid directory path\n")
			utils.PrintInfo("use `" + fullArgs[0] + " -h` for help\n")
		}
	}

	return args
}
