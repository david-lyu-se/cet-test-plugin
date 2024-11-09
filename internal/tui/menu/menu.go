package menu

import (
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type list string

const (
	enumApp      list = "Go to Application settings"
	enumRepo     list = "Edit mono repo directory"
	enumPlugin   list = "Edit Plugins Parent Directory"
	enumFileSync list = "Execute file sync"
)

/* ----------------- Init ------------------ */

func InitMenu() (tea.Model, tea.Cmd) {
	items := []list{enumApp, enumRepo, enumPlugin, enumFileSync}
	pModel := primary{
		Items: items,
	}

	return pModel, nil
}

/* ----------------- Model ----------------- */
type primary struct {
	Items []list
}

func (pModel primary) Init() tea.Cmd {
	return nil
}

func (pModel primary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, variables.Keymap.Quit) {
			cmds = append(cmds, tea.Quit)
		}
	}

	cmds = append(cmds, cmd)

	return pModel, tea.Batch(cmds...)
}

func (pModel primary) View() string {
	return ""
}

/* --------------- Helpers ---------------- */

func (pModel ParentModel) goToSubMenu(i list.Item) (tea.Model, tea.Cmd) {
	switch true {
	case enumApp == i:
		var appModel, cmd = applicationModel.InitAppModel(&variables.Conf.Apps)
		variables.AppModel = &appModel
		return appModel, cmd
	case enumRepo == i:
		return nil, nil
	case enumPlugin == i:
		return nil, nil
	case enumFileSync == i:
		return nil, nil
	default:
		return pModel, nil
	}
