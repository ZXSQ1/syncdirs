package utils

import (
	"fmt"
	"path"

	"github.com/ZXSQ1/syncdirs/files"
	"github.com/fatih/color"
)

var errCol = color.New(color.Bold, color.FgRed)
var infoCol = color.New(color.Bold, color.FgYellow)

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
description: formats info message
arguments:
  - infoMsg: the information message to format

return: no return
*/
func Info(infoMsg string) string {
	return infoCol.Sprint("W: ") + infoMsg
}

/*
description: prints the formatted error
arguments:
  - format: the format string passed to Printf
  - objs: the objects to print in the format

return: no return
*/
func PrintError(format string, objs ...any) {
	fmt.Println(Error(fmt.Sprintf(format, objs...)))
}

/*
description: prints the formatted information
arguments:
  - format: the format string passed to Printf
  - objs: the objects to print

return: no return
*/
func PrintInfo(format string, objs ...any) {
	fmt.Println(Info(fmt.Sprintf(format, objs...)))
}

/*
description: validates the directory path
arguments:
  - dir: the string type path to the directory

return:
  - true if the everything is ok
  - false if something wasn't valid
*/
func ValidateDir(dir string) (string, bool) {
	dir = path.Clean(dir)

	if !files.IsExist(dir) {
		return dir, false
	}

	if isDir, _ := files.IsDir(dir); !isDir {
		return dir, false
	}

	return dir, true
}
