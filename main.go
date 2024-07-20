package main

import (
	"os"

	"github.com/ZXSQ1/syncdirs/ui"
	"github.com/ZXSQ1/syncdirs/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ui.Dirs = ui.Handle(os.Args)
	program := tea.NewProgram(ui.Model{})

	if _, err := program.Run(); err != nil {
		utils.PrintError("error on running")
	}
}
