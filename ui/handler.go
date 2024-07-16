package ui

import (
	"fmt"
	"os"
)

/*
description: prints the help message and exits
*/
func Help() {
	helpMsg := `	usage: syncdirs <directories>
	ex: syncdirs ~/Downloads ~/Desktop ~/Temp
	`

	fmt.Print(helpMsg)
	os.Exit(1)
}
