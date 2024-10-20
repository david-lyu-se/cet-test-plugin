package applicationModel

import (
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	List         list.Model
	hasTryDelete bool
	hasTryEdit   bool
	err          error
	Quitting     bool
}

// Todo not sure if this is needed
type UpdateListMsg struct {
	Item structures.Application
}

func (parentModel AppModel) Init() tea.Cmd {
	//future move operations file create and read here
	return nil
}

func (parentModel AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// tea.msgs default or we can create
	case tea.WindowSizeMsg:
		variables.WindowSize = msg
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	// custom msg for updating Item List after creatihng new Application
	case UpdateListMsg:
		parentModel.List.InsertItem(len(variables.Conf.Apps)-1, msg.Item)
		parentModel.List, cmd = parentModel.List.Update(variables.Conf.Apps)
	// key inputs
	case tea.KeyMsg:
		//clears out error after key press
		return parentModel.handleKeyInputs(msg)
	//cant fallthrough
	default:
		parentModel.List, cmd = parentModel.List.Update(msg)
	}
	return parentModel, cmd
}

func (parentModel AppModel) View() string {
	if parentModel.Quitting {
		return ""
	}

	if parentModel.hasTryDelete {
		return variables.AlertStyle("Cannot delete please go into file and delete") + "\n" + variables.DocStyle.Render(parentModel.List.View())
	}

	return variables.DocStyle.Render(parentModel.List.View()) + "\n"
}

func InitAppModel(apps *structures.Applications) (tea.Model, tea.Cmd) {

	var items = make([]list.Item, len(*apps))
	for i, proj := range *apps {
		items[i] = list.Item(proj)
	}

	var parentModel = AppModel{
		// Mode:     variables.Nav,
		List:         list.New(items, list.NewDefaultDelegate(), 20, 24),
		hasTryDelete: false,
		Quitting:     false,
	}

	if variables.WindowSize.Height != 0 {
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(variables.WindowSize.Width-left-right, variables.WindowSize.Height-top-bottom)
	}

	parentModel.List.SetSize(100, 100)
	parentModel.List.Title = "Applications List"
	parentModel.List.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			variables.Keymap.Create,
			variables.Keymap.Rename,
			variables.Keymap.Delete,
			variables.Keymap.Back,
		}
	}

	return parentModel, func() tea.Msg { return tea.WindowSizeMsg{} }
}

/* Help Functions */

func (parentModel AppModel) handleKeyInputs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if !key.Matches(msg, variables.Keymap.Quit) {
		parentModel.hasTryDelete = false
	}
	if key.Matches(msg, variables.Keymap.Enter) {
		//go to plugins model init
	} else if key.Matches(msg, variables.Keymap.Create) {
		var appModel tea.Model
		// appModel, cmd = InitAppModel(&parentModel)
		return appModel, cmd
		// mm := teaModel.(AppModel)
	} else if key.Matches(msg, variables.Keymap.Delete) {
		//delete environment tell them to do it manually
		parentModel.hasTryDelete = true
	} else if key.Matches(msg, variables.Keymap.Quit) {
		parentModel.Quitting = true
		cmd = tea.Quit
	} else {
		parentModel.List, cmd = parentModel.List.Update(msg)
	}
	return parentModel, cmd
}
