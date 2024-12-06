package sync

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
	"strings"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"
	"time"
)

// Not the model struct
type result struct {
	//daemon-combo
	duration time.Duration
	emoji    string
}

type complete struct{}

/* ----------------- Init ------------------ */

func InitSync(path string, app structures.Application, primary tea.Model) (tea.Model, tea.Cmd) {
	pModel := sync{
		parent:    primary,
		source:    path,
		target:    app.Path,
		isTheme:   strings.Contains(path, "theme"),
		toggleNPM: false,
	}
	// need to check target and source see if it is valid destination

	return pModel, nil
}

/* ----------------- Model ----------------- */
type sync struct {
	parent     tea.Model
	target     string
	source     string
	isTheme    bool
	hasStarted bool
	error      error
	toggleNPM  bool

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
func (pModel sync) themeStart() {
	// Get theme
	// check for vendor or lib
	// if it doesn't exist compser install.
	// are we npm installing here? Nope should happen already
	target := pModel.target + "/themes"
	source := pModel.source
	rsync(source, target)
}

func (pModel sync) pluginStart() {
	// Get plugins
	// check for vendor or lib
	// if it doesn't exist composer install.
	// go to block lib
	// grab all the plugins
	// npm install? eahc plugin
	target := pModel.target + "/plugins"
	source := pModel.source
	rsync(source, target)

}

func rsync(source string, destination string) tea.Cmd {
	// run rsync to application
	cmd := exec.Command("rsync -av --exclude 'node_modules' ", source, destination)
	err := cmd.Run()

	if err != nil {
		//handle error
	}
	return func() tea.Msg {
		return complete{}
	}
}
