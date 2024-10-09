package bubbletea

import (
	"test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectMsg struct {
	ActiveProjectID uint
}

type ParentModel struct {
	Mode     variables.Mode
	List     list.Model
	Input    textinput.Model
	Quitting bool
}

func (parentModel ParentModel) Init() tea.Cmd {
	//future move operations file create and read here
	return tea.Cmd
}

func (parentModel ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		variables.WindowSize = msg
		top, right, bottom, left := variables.DocStyle.GetMargin()
		parentModel.List.SetSize(msg.Width-left-right, msg.Height-top-bottom)

	case tea.KeyMsg:
		if key.Matches(msg, variables.Keymap.Quit) {
			parentModel.Quitting = true
			cmds = append(cmds, tea.Quit)
		}

	}
	return parentModel, tea.Batch(cmds...)
}

func (parentModel ParentModel) View() string {
	if parentModel.Quitting {
		return "Bye"
	}
	if parentModel.Input.Focused() {
		return variables.DocStyle.Render(parentModel.List.View() + "\n" + parentModel.Input.View())
	}
	return variables.DocStyle.Render(parentModel.List.View() + "\n")
}

func InitParent(apps *structs.Environments) (tea.Model, tea.Cmd) {

	input := textinput.New()
	input.Prompt = "$ "
	input.Placeholder = "Project name"
	input.CharLimit = 250
	input.Width = 50

	var items = make([]list.Item, len(*apps))
	for i, proj := range *apps {
		items[i] = list.Item(proj)
	}

	var parentModel = ParentModel{
		Mode:     variables.Nav,
		List:     list.New(items, list.NewDefaultDelegate(), 8, 8),
		Input:    input,
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
