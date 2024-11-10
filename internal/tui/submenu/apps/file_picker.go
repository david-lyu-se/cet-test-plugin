package application

import (
	"errors"
	"os"
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/operations"
	"test-cet-wp-plugin/internal/tui/variables"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type fileModel struct {
	file         filepicker.Model
	SelectedFile string
	err          error
	/* Name of App */
	AppName string
	input   textinput.Model
	isFocus bool
	/* Parent Model */
	appModel *application
}

type clearErrorMsg struct{}

type initFileModelMsg struct{}

func InitFileModel(p *application) (tea.Model, tea.Cmd) {

	if p.fp != nil {
		return *variables.AppModel, func() tea.Msg {
			return initFileModelMsg{}
		}
	}

	fp := filepicker.New()

	fp.AllowedTypes = []string{".txt", ".md", ".sh", ".json"}
	if variables.Conf.WorkingDir == "" {
		fp.CurrentDirectory, _ = os.UserHomeDir()
	} else {
		fp.CurrentDirectory = variables.Conf.WorkingDir
	}
	fp.ShowHidden = true
	fp.AutoHeight = true

	ti := textinput.New()
	ti.Placeholder = "Please give app name"
	ti.CharLimit = 50
	ti.Width = 20

	var appModel = fileModel{
		file:     fp,
		appModel: p,
		input:    ti,
		AppName:  "",
		isFocus:  false,
	}

	return appModel, func() tea.Msg {
		return initFileModelMsg{}
	}
}

func (fm fileModel) Init() tea.Cmd {
	return nil
}

func (fm fileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ask user for name of app/environment
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if fm.isFocus {
			return fm.inputFocusKeyPress(msg)
		} else {
			var model tea.Model
			model, cmd = fm.filePickerKeyPress(msg)
			if model != nil {
				return model, cmd
			}

		}
	// case updateMsg:
	case clearErrorMsg:
		fm.err = nil
	case initFileModelMsg:
		cmds = append(cmds, fm.file.Init())
		cmds = append(cmds, func() tea.Msg { return tea.WindowSizeMsg{Width: 100, Height: 24} })
		if fm.AppName != "" {
			fm.AppName = ""
		}
		return fm, tea.Batch(cmds...)
	}

	// var cmd tea.Cmd
	fm.file, cmd = fm.file.Update(msg)

	// Did the user select a file?
	if didSelect, path := fm.file.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		fm.SelectedFile = path
		fm.isFocus = true
		fm.input.Focus()
		cmd = func() tea.Msg { return tea.KeyMsg{} }
		cmds = append(cmds, cmd)
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := fm.file.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		fm.err = errors.New(path + " is not valid.")
		fm.SelectedFile = ""
		return fm, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return fm, cmd
}

func (fm fileModel) View() string {
	var s strings.Builder

	if fm.isFocus {
		s.WriteString("Please enter you application alias")
		s.WriteString(variables.DocStyle.Render(fm.input.View()))
		// s.WriteString(fm.input.)
	} else {
		s.WriteString(fm.file.CurrentDirectory)
		s.WriteString("\n")
		if fm.err != nil {
			s.WriteString(fm.file.Styles.DisabledFile.Render(fm.err.Error()))
		} else if fm.SelectedFile == "" {
			s.WriteString("Pick the composer.json file in root of application")
		} else {
			s.WriteString("Selected file: " + fm.file.Styles.Selected.Render(fm.SelectedFile))
		}
		s.WriteString("\n\n" + fm.file.View() + "\n")

	}

	return s.String()
}

/**** ---------------- ****/
/**** Helper Functions ****/

// Move to variables or varialbes/functions
func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (fm fileModel) inputFocusKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch true {
	case key.Matches(msg, variables.Keymap.Enter):
		fm.AppName = fm.input.Value()
		if fm.AppName != "" {

			// do nothing
			fm.isFocus = false
			fm.input.Blur()

			var app = structures.Application{
				Name: fm.AppName,
				Path: fm.file.CurrentDirectory,
			}

			variables.Conf.Apps = append(variables.Conf.Apps, app)
			operations.WriteFile(variables.File, variables.Conf)

			//go back to Application Menu
			return fm.appModel, func() tea.Msg {
				return updateListMsg{
					Item: app,
				}
			}
		}
	case key.Matches(msg, variables.Keymap.Back):
		fm.isFocus = false
		fm.AppName = ""
	}
	fm.input, cmd = fm.input.Update(msg)
	return fm, cmd
}

func (fm fileModel) filePickerKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, variables.Keymap.Quit):
		return fm.appModel, func() tea.Msg { return variables.InitAppModel{} }
	}
	return nil, nil
}
