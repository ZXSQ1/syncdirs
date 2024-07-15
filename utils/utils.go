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
	- format: the format string passed to Printf
	- objs: the objects to print in the format
return: no return
*/
func PrintError(format string , objs ...any) {
	fmt.Println(Error(fmt.Sprintf(format, objs...)))
}
