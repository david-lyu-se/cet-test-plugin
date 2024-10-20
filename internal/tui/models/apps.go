package models

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

type AppModel struct {
	file         filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
	/* Name of App */
	AppName string
	input   textinput.Model
	isFocus bool
	/* Parent Model */
	parentModel *ParentModel
	//Test
	cmd tea.Cmd
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
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ask user for name of app/environment
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if a.isFocus {
			switch {
			case key.Matches(msg, variables.Keymap.Enter):
				a.AppName = a.input.Value()
				if a.AppName != "" {
					// do nothing
					a.isFocus = false
					a.input.Blur()

					var app = structures.Application{
						Name: a.AppName,
						Path: a.file.CurrentDirectory,
					}

					variables.Conf.Apps = append(variables.Conf.Apps, app)
					operations.WriteFile(variables.File, variables.Conf)

					return a.parentModel, func() tea.Msg {
						return UpdateListMsg{
							Item: app,
						}
					}
				}
				// cmd = variables.TextInputs(msg, &a.input)
				// cmd = func()tea.Msg {return }
			case key.Matches(msg, variables.Keymap.Back):
				a.isFocus = false
				a.AppName = ""

				// default:
			}
			a.input, cmd = a.input.Update(msg)
			return a, cmd
		} else {
			switch {
			case key.Matches(msg, variables.Keymap.Quit):
				a.quitting = true
				return a.parentModel, func() tea.Msg { return nil }
			}
		}
	// case updateMsg:
	case clearErrorMsg:
		a.err = nil
	case initMsg:
		cmds = append(cmds, a.file.Init())
		cmds = append(cmds, func() tea.Msg { return tea.WindowSizeMsg{Width: 100, Height: 24} })
		if a.AppName != "" {
			a.AppName = ""
		}
		return a, tea.Batch(cmds...)
	}

	// var cmd tea.Cmd
	a.file, cmd = a.file.Update(msg)

	// Did the user select a file?
	if didSelect, path := a.file.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		a.SelectedFile = path
		a.isFocus = true
		a.input.Focus()
		cmd = func() tea.Msg { return tea.KeyMsg{} }
		cmds = append(cmds, cmd)
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := a.file.DidSelectDisabledFile(msg); didSelect {
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

	if a.isFocus {
		s.WriteString("Please enter you application alias")
		s.WriteString(variables.DocStyle.Render(a.input.View()))
		// s.WriteString(a.input.)
	} else {
		s.WriteString(a.file.CurrentDirectory)
		s.WriteString("\n")
		if a.err != nil {
			s.WriteString(a.file.Styles.DisabledFile.Render(a.err.Error()))
		} else if a.SelectedFile == "" {
			s.WriteString("Pick the composer.json file in root of application")
		} else {
			s.WriteString("Selected file: " + a.file.Styles.Selected.Render(a.SelectedFile))
		}
		s.WriteString("\n\n" + a.file.View() + "\n")

	}

	return s.String()
}

func InitAppModel(p *ParentModel) (tea.Model, tea.Cmd) {
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

	var appModel = AppModel{
		file:        fp,
		quitting:    false,
		parentModel: p,
		input:       ti,
		AppName:     "",
		isFocus:     false,
	}

	return appModel, func() tea.Msg {
		return initMsg{
			msg: "Initializing",
		}
	}
}

func addApplication(appName string, appDirectory string, apps structures.Applications) {
	app := structures.Application{
		Name: appName,
		Path: getAppDirectory(appDirectory),
	}
	newApps := append(apps, app)
	variables.Conf.Apps = newApps

	operations.WriteFile(variables.File, variables.Conf)
}

func getAppDirectory(filePath string) string {
	dirPath := ""
	if filePath != "" {
		var index = strings.LastIndex(filePath, "/")
		if index > -1 {
			dirPath = filePath[index : len(filePath)-1]
		}
	}

	return dirPath
}
