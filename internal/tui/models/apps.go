package models

import (
	"errors"
	"os"
	"strings"
	"test-cet-wp-plugin/internal/tui/variables"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	File         filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
	/* Name of App */
	appName string
	input   textinput.Model
	/* Parent Model */
	parentModel *ParentModel
}

type updateMsg struct{}

type clearErrorMsg struct{}

type initMsg struct {
	msg string
}

// Move to variables or varialbes/functions
func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (a AppModel) Init() tea.Cmd {
	return nil
	// return a.File.Init()
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ask user for name of app/environment
	var cmd tea.Cmd
	// var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// if a.input.Focused() {
		// 	cmd = variables.TextInputs(msg, &a.input)
		// }
		// if key.Matches(msg, variables.Keymap.Quit) {
		// 	return a.parentModel, nil
		// 	// return a, tea.QuitMsg{}
		// }
		// case updateMsg:

		// default:
		// 	cmd = nil
		switch msg.String() {
		case "ctrl+c", "q":
			a.quitting = true
			return a.parentModel, func() tea.Msg { return UpdateListMsg{} }
		}
	case clearErrorMsg:
		a.err = nil
	case initMsg:
		cmd = a.File.Init()
		return a, cmd
	}

	// var cmd tea.Cmd
	a.File, cmd = a.File.Update(msg)

	// Did the user select a file?
	if didSelect, path := a.File.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		a.SelectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := a.File.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		a.err = errors.New(path + " is not valid.")
		a.SelectedFile = ""
		return a, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return a, cmd
}

func (a AppModel) View() string {
	if a.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString(a.File.CurrentDirectory)
	s.WriteString("\n")
	if a.err != nil {
		s.WriteString(a.File.Styles.DisabledFile.Render(a.err.Error()))
	} else if a.SelectedFile == "" {
		s.WriteString("Pick the composer.json file in root of application")
	} else {
		s.WriteString("Selected file: " + a.File.Styles.Selected.Render(a.SelectedFile))
	}
	s.WriteString("\n\n" + a.File.View() + "\n")

	return s.String()
}

func InitAppModel(p *ParentModel) (tea.Model, tea.Cmd) {
	fp := filepicker.New()

	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md", ".sh"}

	if variables.Conf.WorkingDir != "" {
		fp.CurrentDirectory, _ = os.UserHomeDir()
	} else {
		fp.CurrentDirectory = variables.Conf.WorkingDir
	}

	var appModel = AppModel{
		File:        fp,
		quitting:    false,
		parentModel: p,
	}

	return appModel, func() tea.Msg {
		return initMsg{
			msg: "Initializing",
		}
	}
}
