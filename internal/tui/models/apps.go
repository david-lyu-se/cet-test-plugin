package models

import (
	"os"
	"strings"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	File         filepicker.Model
	SelectedFile string
	quiting      bool
	err          error
	/* Name of App */
	input   textinput.Model
	appName string
	/* Parent Model */
	parentModel *ParentModel
}

type updateMsg struct{}

func (a AppModel) Init() tea.Cmd {
	return a.File.Init()
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ask user for name of app/environment
	//
	var cmd tea.Cmd
	// var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if a.input.Focused() {
			cmd = variables.TextInputs(msg, &a.input)
		}
		if key.Matches(msg, variables.Keymap.Quit) {
			return a.parentModel, nil
			// return a, tea.QuitMsg{}
		}
	case updateMsg:

	default:
		cmd = nil
	}

	//
	// double check
	//
	// cmds = append(cmds, cmd)
	return a, cmd
	// ask for file path of environment
}

func (a AppModel) View() string {
	if a.quiting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n ")
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

	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	fp.CurrentDirectory, _ = os.UserHomeDir()

	var appModel = AppModel{
		File:        fp,
		quiting:     false,
		parentModel: p,
	}

	return appModel, func() tea.Msg { return UpdateListMsg{} }
}
