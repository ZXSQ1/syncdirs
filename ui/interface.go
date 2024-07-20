package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var dirs []string

type Model struct {
	sourceFile chan string
	destFile   chan string
	sourceDir  chan string
	destDir    chan string
	progress   chan float32
}

func (m Model) Init() tea.Cmd {

	return func() tea.Msg {
		return nil
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	return ""
}
