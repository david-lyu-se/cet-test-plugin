package menu

import (
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	application "test-cet-wp-plugin/internal/tui/submenu/apps"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type item string

const (
	enumApp      item = "Go to Application settings"
	enumRepo     item = "Edit mono repo directory"
	enumPlugin   item = "Edit Plugins Parent Directory"
	enumFileSync item = "Execute file sync"
)

/* ----------------- Init ------------------ */

func InitMenu() (tea.Model, tea.Cmd) {
	items := []item{enumApp, enumRepo, enumPlugin, enumFileSync}
	pModel := primary{
		Items: items,
		title: "Main Menu:",
	}

	return pModel, nil
}

/* ----------------- Model ----------------- */
type primary struct {
	Items  []item
	title  string
	cursor int
	choice item
	// application information
	appChosen    structures.Application
	repoChosen   string
	pluginChosen string
	isEnter      bool
}

func (pModel primary) Init() tea.Cmd {
	return nil
}

func (pModel primary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, variables.Keymap.Enter):
			pModel.goToSubMenu()
		case key.Matches(msg, variables.Keymap.Quit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, variables.Keymap.Down):
			pModel.cursor++
			if pModel.cursor >= len(pModel.Items) {
				pModel.cursor = 0
			}
		case key.Matches(msg, variables.Keymap.Up):
			pModel.cursor--
			if pModel.cursor < 0 {
				pModel.cursor = len(pModel.Items) - 1
			}
		}
	}

	cmds = append(cmds, cmd)

	return pModel, tea.Batch(cmds...)
}

func (pModel primary) View() string {
	s := strings.Builder{}
	count := len(pModel.Items)

	if pModel.isEnter {
		s.WriteString("Test")
	}

	//Create style Title
	s.WriteString(pModel.title)
	s.WriteString((string)(count))

	//Move this to its own function for styling create style Body
	for i := 0; i < count; i++ {
		s.WriteString("\n")
		if i == pModel.cursor {
			s.WriteString("[x] ")
			// color this with lipgloss
			s.WriteString(string(pModel.Items[i]))
		} else {
			s.WriteString("[ ] ")
			s.WriteString(string(pModel.Items[i]))
		}
	}

	return s.String()
}

/* --------------- Helpers ---------------- */

func (pModel primary) goToSubMenu() (tea.Model, tea.Cmd) {
	choice := pModel.Items[pModel.cursor]
	pModel.isEnter = true

	switch choice {
	case enumApp:
		return application.InitMenu()
	case enumRepo:
	case enumPlugin:
	case enumFileSync:
	}

	return pModel, nil
}

func formatHeader() string {
	return ""
}

func formatBody() string {
	return ""
}
