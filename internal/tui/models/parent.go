package parentModel

import (
	"strings"
	applicationModel "test-cet-wp-plugin/internal/tui/models/apps"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/* Initial variables */
type item string

func (i item) FilterValue() string { return string(i) }
func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }

const (
	enumApp      item = "Go to Application settings"
	enumRepo     item = "Edit mono repo directory"
	enumPlugin   item = "Edit Plugins Parent Directory"
	enumFileSync item = "Execute file sync"
)

/* Start of Tea Model */
type ParentModel struct {
	List         list.Model
	hasTryDelete bool
	hasTryEdit   bool
	err          error
	Quitting     bool
}

func InitParent() (tea.Model, tea.Cmd) {

	var items = []list.Item{
		item(enumApp),
		item(enumRepo),
		item(enumPlugin),
		item(enumFileSync),
	}

	var parentModel = ParentModel{
		List:         list.New(items, list.NewDefaultDelegate(), 20, 24),
		hasTryDelete: false,
		hasTryEdit:   false,
		Quitting:     false,
	}

	if variables.WindowSize.Height != 0 {
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(variables.WindowSize.Width-left-right, variables.WindowSize.Height-top-bottom)
	}

	parentModel.List.SetSize(100, 100)
	parentModel.List.Title = "Start Up Menu"
	parentModel.List.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			variables.Keymap.Enter,
			variables.Keymap.Quit,
			variables.Keymap.Back,
		}
	}

	return parentModel, func() tea.Msg { return tea.WindowSizeMsg{} }
}

func (pModel ParentModel) Init() tea.Cmd {
	return nil
}

func (pModel ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// variables.WindowSize = msg
		top, right, bottom, left := variables.DocStyle.GetMargin()
		pModel.List.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		if key.Matches(msg, variables.Keymap.Enter) {
			var item = pModel.List.SelectedItem()
			return pModel.goToSubMenu(item)
		} else if (key.Matches(msg, variables.Keymap.Quit)) || key.Matches(msg, variables.Keymap.Quit) {
			pModel.Quitting = true
			cmd = tea.Quit
		} else {
			// Need this for default keys to work
			pModel.List, cmd = pModel.List.Update(msg)
		}
	default:
		pModel.List, cmd = pModel.List.Update(msg)
	}
	return pModel, cmd
}

func (pModel ParentModel) View() string {

	if pModel.Quitting {
		return ""
	}

	var s strings.Builder

	s.WriteString(variables.DocStyle.Render(pModel.List.View()) + "\n")

	return s.String()
}

/* Helper Functions */

func (pModel ParentModel) goToSubMenu(i list.Item) (tea.Model, tea.Cmd) {
	switch true {
	case enumApp == i:
		var appModel, cmd = applicationModel.InitAppModel(&variables.Conf.Apps)
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
}
