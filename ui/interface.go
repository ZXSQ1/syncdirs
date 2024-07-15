package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	sourceDir string
	destDir string
	currentFile string
	progress float32
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	return ""
}
