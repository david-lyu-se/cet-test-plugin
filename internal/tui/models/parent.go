package models

import (
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ParentModel struct {
	Mode         variables.Mode
	List         list.Model
	hasTryDelete bool
	err          error
	Quitting     bool
}

// Todo not sure if this is needed
type UpdateListMsg struct{}

func (parentModel ParentModel) Init() tea.Cmd {
	//future move operations file create and read here
	return nil
}

func (parentModel ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	// tea.msgs default or we can create
	case tea.WindowSizeMsg:
		variables.WindowSize = msg
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	// custom msg for updating Item List after creatihng new Environment
	case UpdateListMsg:
	// tea.msg for keypress
	case tea.KeyMsg:
		if key.Matches(msg, variables.Keymap.Enter) {
			//go to plugins model init
		} else if key.Matches(msg, variables.Keymap.Create) {
			var appModel tea.Model
			appModel, cmd = InitAppModel(&parentModel)
			return appModel, cmd
			// mm := teaModel.(AppModel)
		} else if key.Matches(msg, variables.Keymap.Delete) {
			//delete environment tell them to do it manually
			parentModel.hasTryDelete = true
		} else if key.Matches(msg, variables.Keymap.Quit) {
			parentModel.Quitting = true
			//Work around for graceful termination...
			// variables.ParentProgram.ReleaseTerminal()
			cmd = tea.Quit
		} else {
			parentModel.List, cmd = parentModel.List.Update(msg)
		}
		//cant fallthrough
	default:
		parentModel.List, cmd = parentModel.List.Update(msg)
	}
	if cmd != nil {
		cmds = append(cmds, cmd)
		cmd = nil
	}
	return parentModel, tea.Batch(cmds...)
}

func (parentModel ParentModel) View() string {
	if parentModel.Quitting {
		return ""
	}

	if parentModel.hasTryDelete {
		parentModel.hasTryDelete = false
		return variables.DocStyle.Render("\nCannot delete items, please go to json file and delete manually" + parentModel.List.View())
	}

	return variables.DocStyle.Render(parentModel.List.View() + "\n")
}

func InitParent(apps *structures.Applications) (tea.Model, tea.Cmd) {

	var items = make([]list.Item, len(*apps))
	for i, proj := range *apps {
		items[i] = list.Item(proj)
	}

	var parentModel = ParentModel{
		Mode:     variables.Nav,
		List:     list.New(items, list.NewDefaultDelegate(), 8, 8),
		Quitting: false,
	}

	if variables.WindowSize.Height != 0 {
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(variables.WindowSize.Width-left-right, variables.WindowSize.Height-top-bottom)
	}

	parentModel.List.Title = "Applications"
	parentModel.List.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			variables.Keymap.Create,
			variables.Keymap.Rename,
			variables.Keymap.Delete,
			variables.Keymap.Back,
		}
	}

	return parentModel, func() tea.Msg { return "error" }
}
