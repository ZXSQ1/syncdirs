package ui

import (
	"github.com/fatih/color"

	tea "github.com/charmbracelet/bubbletea"
)

var Dirs []string

type Model struct {
	sourceFile chan string
	destFile   chan string
	sourceDir  chan string
	destDir    chan string
	progress   chan float32
}

const (
	EventDataCreated = 1982
)

func (m Model) Init() tea.Cmd {
	Synchronize(Dirs, m.sourceFile, m.destFile, m.sourceDir, m.destDir, m.progress)

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	var displayString string
	var sourceDir, destDir, sourceFile, destFile string
	var progress float32

	sourceDir = <-m.sourceDir
	destDir = <-m.destDir
	sourceFile = <-m.sourceFile
	destFile = <-m.destFile
	progress = <-m.progress

	greenCol := color.New(color.Bold, color.FgGreen)
	yellowCol := color.New(color.Bold, color.FgYellow)

	progressText := yellowCol.Sprintf("%.1f\n\n", progress)

	if progress == 100.0 {
		progressText = greenCol.Sprint("finished\n\n")
	}

	displayString = sourceDir + " ->\n" + destDir + " " + progressText
	displayString += sourceFile + " ->\n" + destFile

	return displayString
}
