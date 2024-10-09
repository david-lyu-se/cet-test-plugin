package bubbletea

import (
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type EnvironmentInputs struct {
	file         filepicker.Model
	selectedFile string
	Quiting      bool
	err          error
	/* Name of App */
}

func (EnvironmentInputs) init() tea.Cmd {
	return nil
}

func (EnvironmentInputs) update() {
	// ask user for name of app/environment
	//
	// double check
	//
	// ask for file path of environment

}

func (EnvironmentInputs) view() {

}

func InitEnvironmentTUI() {

}
