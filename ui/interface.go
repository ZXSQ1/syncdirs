package ui

import (
	"github.com/ZXSQ1/syncdirs/sync"
	tea "github.com/charmbracelet/bubbletea"
)

var dirs []string

type Model struct {
	sourceFile string
	destFile   string
	progress   float32
}

var syncData = make(chan *sync.SyncData)

func (m Model) Init() tea.Cmd {
	sync.SynchronizeMultiple(dirs, syncData)

	return func() tea.Msg {
		if data, ok := <-syncData; ok {
			return data
		}

		return nil
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	return ""
}
