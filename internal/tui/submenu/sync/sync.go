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
		parent:           primary,
		source:           path,
		target:           app.Path,
		isTheme:          strings.Contains(path, "theme"),
		isVendorIncluded: false,
	}
	// need to check target and source see if it is valid destination

	return pModel, nil
}

/* ----------------- Model ----------------- */
type sync struct {
	parent           tea.Model
	target           string
	source           string
	isTheme          bool
	hasStarted       bool
	error            error
	isVendorIncluded bool

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
		case key.Matches(msg, variables.Keymap.Toggle):
			pModel.isVendorIncluded = !pModel.isVendorIncluded
		case key.Matches(msg, variables.Keymap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	cmds = append(cmds, cmd)

	return pModel, tea.Batch(cmds...)
}

func (pModel sync) View() string {
	s := strings.Builder{}

	if pModel.isVendorIncluded {
		s.WriteString("Sync vendor is on")
	} else {
		s.WriteString("Sync vendor is off")
	}

	if pModel.hasStarted {
		s.WriteString("Press enter to sync")
	}

	//help text
	s.WriteString("\nKeymaps: \n")
	s.WriteString("Enter - start sync; t - toggle vendor sync; q - quit")
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
	pModel.rsync(source, target)
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
	pModel.rsync(source, target)

}

func (pModel sync) rsync(source string, destination string) tea.Cmd {
	// run rsync to application

	vendorExclude := pModel.checkComposer()
	cmd := exec.Command("rsync -av --exclude 'node_modules"+vendorExclude+"' ", source, destination)
	err := cmd.Run()

	if err != nil {
		//handle error
	}
	return func() tea.Msg {
		return complete{}
	}
}

func (pModel sync) checkComposer() string {
	s := strings.Builder{}

	//check to see if user wants to include composer in ecosystem, false
	//check composer exists in wp ecosystem

	//check if vendor exists in plugin
	if !pModel.isVendorIncluded {
		s.WriteString("vendor lib")
	}

	return s.String()
}
