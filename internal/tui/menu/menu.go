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
}

func (pModel primary) Init() tea.Cmd {
	return nil
}

func (pModel primary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case variables.UpdateAppChosen:
		pModel.appChosen = msg.Application
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, variables.Keymap.Enter):
			return pModel.goToSubMenu()
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

	//Create style Headers
	s.WriteString(variables.TitleStyle(pModel.title))
	s.WriteString((string)(count))

	//User choices here:
	s.WriteString("\n")
	s.WriteString(pModel.formatUserChoice())

	//Move this to its own function for styling create style Body
	s.WriteString(pModel.formatBody())
	// for i := 0; i < count; i++ {
	// 	s.WriteString("\n")
	// 	if i == pModel.cursor {
	// 		s.WriteString("[x] ")
	// 		// color this with lipgloss
	// 		s.WriteString(string(pModel.Items[i]))
	// 	} else {
	// 		s.WriteString("[ ] ")
	// 		s.WriteString(string(pModel.Items[i]))
	// 	}
	// }

	return s.String()
}

/* --------------- Helpers ---------------- */

func (pModel primary) goToSubMenu() (tea.Model, tea.Cmd) {
	choice := pModel.Items[pModel.cursor]

	switch choice {
	case enumApp:
		return application.InitMenu(pModel)
	case enumRepo:
	case enumPlugin:
	case enumFileSync:
	}

	return pModel, nil
}

func formatHeader() string {
	return ""
}

func (pModel primary) formatUserChoice() string {
	s := strings.Builder{}
	choices := []string{
		"App Chosen: " + pModel.appChosen.Name,
		"Monorepo path: " + pModel.repoChosen,
		"Plugin path: " + pModel.pluginChosen,
	}

	for index, element := range choices {
		if pModel.cursor == index {
			s.WriteString(variables.BlinkingStyle.Render(element))
		} else {
			s.WriteString(element)
		}
		if index < len(choices)-1 {
			s.WriteString("\n")
		}
	}

	return variables.UserChoiceStyle(s.String())
}

func (pModel primary) formatBody() string {
	s := strings.Builder{}

	for i := 0; i < len(pModel.Items); i++ {
		s.WriteString("\n")
		if i == pModel.cursor {
			// color this with lipgloss
			s.WriteString(variables.ModelSelectStyle(string(pModel.Items[i])))

		} else {
			s.WriteString(variables.ModelChoiceStyle(string(pModel.Items[i])))
		}
	}

	return variables.ModelChoiceContainerStyle(s.String())
}
