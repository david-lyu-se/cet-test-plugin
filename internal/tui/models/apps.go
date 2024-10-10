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
	Quiting      bool
	err          error
	/* Name of App */
	input textinput.Model
}

func (AppModel) init() tea.Cmd {
	return nil
}

func (AppModel) update() {
	// ask user for name of app/environment
	//
	// double check
	//
	// ask for file path of environment

}

func (AppModel) view() {

}

func InitAppModel(apps structures.Applications) {

}
