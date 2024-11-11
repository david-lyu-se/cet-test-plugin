package application

import (
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type updateListMsg struct {
	Item structures.Application
}

/* ----------------- Init ------------------ */

func InitMenu(parent tea.Model) (tea.Model, tea.Cmd) {
	aModel := application{
		Items:  variables.Conf.Apps,
		title:  "Application List",
		parent: parent,
	}

	return aModel, nil
}

/* ----------------- Model ----------------- */
type application struct {
	Items  []structures.Application
	title  string
	cursor int
	choice structures.Application
	// Ref to parent menu
	parent tea.Model
	// non list
	fp           *tea.Model
	hasTryDelete bool
	hasTryEdit   bool
	err          error
}

func (aModel application) Init() tea.Cmd {
	return nil
}

func (aModel application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case updateListMsg:
		aModel.Items = append(aModel.Items, msg.Item)
	case tea.KeyMsg:
		return aModel.handleKeyInputs(msg)
	}

	cmds = append(cmds, cmd)

	return aModel, tea.Batch(cmds...)
}

func (aModel application) View() string {
	s := strings.Builder{}
	count := len(aModel.Items)

	//Create style Title
	s.WriteString(variables.TitleStyle(aModel.title))
	s.WriteString((string)(count))

	//Move this to its own function for styling create style Body
	s.WriteString("\n")
	s.WriteString(aModel.formatBody())

	return s.String()
}

/* --------------- Helpers ---------------- */

func (aModel application) handleKeyInputs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	//clears out no delete/edit msg in view
	if !aModel.hasTryDelete {
		aModel.hasTryDelete = false
	}
	if !aModel.hasTryEdit {
		aModel.hasTryEdit = false
	}

	switch true {
	case key.Matches(msg, variables.Keymap.Create):
		// var fileModel tea.Model
		return InitFileModel(&aModel)
		// aModel.fp = &fileModel
		// return fileModel, cmd
	case key.Matches(msg, variables.Keymap.Delete):
		aModel.hasTryDelete = true
	case key.Matches(msg, variables.Keymap.Edit):
		aModel.hasTryEdit = true
	// Goes back to primary menu with Application information
	case key.Matches(msg, variables.Keymap.Enter):
		cmd = func() tea.Msg {
			return variables.UpdateAppChosen{
				Application: aModel.Items[aModel.cursor],
			}
		}
		fallthrough
	case key.Matches(msg, variables.Keymap.Back):
		fallthrough
	case key.Matches(msg, variables.Keymap.Quit):
		return aModel.parent, cmd
	case key.Matches(msg, variables.Keymap.Down):
		aModel.cursor++
		if aModel.cursor >= len(aModel.Items) {
			aModel.cursor = 0
		}
	case key.Matches(msg, variables.Keymap.Up):
		aModel.cursor--
		if aModel.cursor < 0 {
			aModel.cursor = len(aModel.Items) - 1
		}
	}
	return aModel, cmd
}

func (aModel application) formatBody() string {
	s := strings.Builder{}

	for i := 0; i < len(aModel.Items); i++ {
		if i == aModel.cursor {
			// color this with lipgloss
			s.WriteString("\n")
			s.WriteString(variables.ModelSelectStyle(string(aModel.Items[i].Name) + "\n" + string(aModel.Items[i].Path)))
		} else {
			s.WriteString("\n")
			s.WriteString(variables.ModelChoiceStyle(string(aModel.Items[i].Name)) + "\n")
			s.WriteString(variables.ModelChoiceStyle(string(aModel.Items[i].Path)))
		}

	}
	return variables.ModelChoiceContainerStyle(s.String())
}
