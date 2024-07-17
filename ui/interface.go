package ui

import (
	"github.com/ZXSQ1/syncdirs/sync"
	tea "github.com/charmbracelet/bubbletea"
)

var dirs []string

type Model struct {
	sourceDir   string
	destDir     string
	currentFile string
	progress    float32
}

func (m Model) Init() tea.Cmd {
	var syncData = make(chan *sync.SyncData)
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
