package apps

import (
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	cursor int
	Choice structures.Application
	fp     *tea.Model
}

type updateAppMsg struct {
	Item structures.Application
}

func InitApps() tea.Model {
	m := AppModel{
		cursor: 0,
		// Choice: "",
	}

	return m
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.Choice = variables.Conf.Apps[m.cursor]
			fallthrough
		case "ctrl+c", "q", "esc":
			return *variables.ParentModel, func() tea.Msg {
				return tea.WindowSizeMsg{
					Width:  variables.WindowSize.Width,
					Height: variables.WindowSize.Height,
				}
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(variables.Conf.Apps) {
				m.cursor = 0
			}
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(variables.Conf.Apps) - 1
			}
		case "c":
			return InitFileModel(&m)
		}
	}
	return m, nil
}

func (m AppModel) View() string {
	s := strings.Builder{}

	s.WriteString("Application List: \n")

	for i := 0; i < len(variables.Conf.Apps); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(variables.Conf.Apps[i].Name)
		s.WriteString(variables.Conf.Apps[i].Path)
		s.WriteString("\n")

	}
	return s.String()

	// return ""
}
