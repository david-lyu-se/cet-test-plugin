package models

import (
	"os"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PluginModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (PluginModel) Init() tea.Cmd {
	return nil
}

func (PluginModel) Update() (tea.Model, tea.Cmd) {
	return nil, nil
}

func (PluginModel) View() string {
	return ""
}

func InitPlugin() (tea.Model, tea.Cmd) {
	if variables.PluginModel != nil {
		return *variables.PluginModel, func() tea.Msg {
			return nil
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
