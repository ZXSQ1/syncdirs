package utils

import (
	"fmt"
	"path"

	"github.com/ZXSQ1/syncdirs/files"
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
func PrintError(format string, objs ...any) {
	fmt.Println(Error(fmt.Sprintf(format, objs...)))
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

	if isDir, err := files.IsDir(dir); !isDir {
		fmt.Println(err)

		return dir, false
	}

	return dir, true
}
