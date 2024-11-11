package sync

import (
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Not the model struct
type result struct {
	//daemon-combo
	duration time.Duration
	emoji    string
}

/* ----------------- Init ------------------ */

func InitSync(path string, app structures.Application, primary tea.Model) (tea.Model, tea.Cmd) {
	pModel := sync{
		parent: primary,
		target: path,
		source: app.Path,
	}

	return pModel, nil
}

/* ----------------- Model ----------------- */
type sync struct {
	parent     tea.Model
	target     string
	source     string
	hasStarted bool
	error      error

	spinner spinner.Model
	results []result
}

func (pModel sync) Init() tea.Cmd {
	return nil
}

func (pModel sync) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, variables.Keymap.Enter):
			//start deamon
		case key.Matches(msg, variables.Keymap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	cmds = append(cmds, cmd)

	return pModel, tea.Batch(cmds...)
}

func (pModel sync) View() string {
	s := strings.Builder{}

	if pModel.hasStarted {
		s.WriteString("Press enter to sync")
	}

	return s.String()
}

/* --------------- Helpers ---------------- */
