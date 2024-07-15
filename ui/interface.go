package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	currentFile string
	progress float32
}

func (m Model) Init() tea.Cmd {
	
}
