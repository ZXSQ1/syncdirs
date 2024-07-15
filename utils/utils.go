package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var errCol = color.New(color.Bold, color.FgRed)


/*
description: formats error message
arguments:
	- errMsg: the string error message to format
return: the formatted error message string
*/
func Error(errMsg string) string {
	return errCol.Sprint("E: ") + errMsg
}

/*
description: prints the formatted error
arguments:
	- errMsg: the string error message
return: no return
*/
func PrintError(errMsg string) {
	fmt.Println(Error(errMsg))
}
