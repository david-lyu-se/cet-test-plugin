package plugins

import (
	"errors"
	"os"
	"strings"
	"test-cet-wp-plugin/internal/tui/variables"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type fileModel struct {
	file         filepicker.Model
	SelectedFile string
	err          error
	/* Parent Model */
	primary tea.Model
}

type clearErrorMsg struct{}

type initFileModelMsg struct{}

func InitFileModel(p tea.Model) (tea.Model, tea.Cmd) {

	fp := filepicker.New()

	fp.AllowedTypes = []string{".txt", ".md", ".sh", ".json"}
	if variables.Conf.WorkingDir == "" {
		fp.CurrentDirectory, _ = os.UserHomeDir()
	} else {
		fp.CurrentDirectory = variables.Conf.WorkingDir
	}
	fp.ShowHidden = true
	fp.AutoHeight = true

	var appModel = fileModel{
		file:    fp,
		primary: p,
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
		var model tea.Model
		model, cmd = fm.filePickerKeyPress(msg)
		if model != nil {
			return model, cmd
		}
	case clearErrorMsg:
		fm.err = nil
	case initFileModelMsg:
		cmds = append(cmds, fm.file.Init())
		cmds = append(cmds, func() tea.Msg { return tea.WindowSizeMsg{Width: 100, Height: 24} })
		return fm, tea.Batch(cmds...)
	}

	// var cmd tea.Cmd
	fm.file, cmd = fm.file.Update(msg)

	// Did the user select a file?
	if didSelect, path := fm.file.DidSelectFile(msg); didSelect {
		fm.SelectedFile = path
		return fm.primary, func() tea.Msg {
			return variables.UpdatePluginRepo{
				Path: fm.file.CurrentDirectory,
			}
		}
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

func (fm fileModel) filePickerKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, variables.Keymap.Quit):
		return fm.primary, func() tea.Msg { return nil }
	}
	return nil, nil
}
