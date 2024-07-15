package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	currentDirectory string
	currentFile string
	progress float32
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {

	}
}
