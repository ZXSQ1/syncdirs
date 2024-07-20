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

const (
	EventDataChanged = 1982
)

func (m Model) Init() tea.Cmd {
	Synchronize(dirs, m.sourceFile, m.destFile, m.sourceDir, m.destDir, m.progress)

	return func() tea.Msg {
		if _, ok := <-m.sourceFile; ok {
			m.sourceFile <- <-m.sourceFile

			return EventDataChanged
		} else {
			return nil
		}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	return ""
}
