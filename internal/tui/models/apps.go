package models

import (
	structures "test-cet-wp-plugin/internal/model/structs"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	file         filepicker.Model
	selectedFile string
	quiting      bool
	err          error
	/* Name of App */
	input textinput.Model
}

func (a AppModel) Init() tea.Cmd {
	return a.file.Init()
}

func (a AppModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	// ask user for name of app/environment
	//
	// double check
	//
	// ask for file path of environment
	return a, nil
}

func (a AppModel) View() string {
	return ""
}

func InitAppModel(apps structures.Applications) (tea.Model, tea.Cmd) {
	fp := filepicker.New()
	// if()

	var appModel = AppModel{
		file:    fp,
		quiting: false,
	}

	return appModel, func() tea.Msg { return UpdateListMsg{} }
}
