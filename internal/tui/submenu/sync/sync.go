package sync

import (
	"os"
	"os/exec"
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
			pModel.Init()
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
func (pModel sync) init() {

	if pModel.isTheme {
		pModel.rsync("vendor")
	} else {
		pModel.rsync("lib")
	}
}

func (pModel sync) rsync(vendorName string) tea.Cmd {
	// run rsync to application
	vendorExclude := pModel.checkComposer(vendorName)
	cmd := exec.Command(
		"rsync -av --exclude 'node_modules"+vendorExclude+"' ",
		pModel.source,
		pModel.source)
	err := cmd.Run()

	if err != nil {
		//handle error
	}
	return func() tea.Msg {
		return complete{}
	}
}

func (pModel sync) checkComposer(vendorName string) string {
	s := strings.Builder{}
	overWriteVendorIncluded := false
	//check to see if user wants to include composer in ecosystem, false
	//check composer exists in wp ecosystem
	vendorPath := pModel.source + "/" + vendorName
	_, err := os.Stat(vendorPath)
	if err != nil {
		runComposerCmd := exec.Command("composer install --working-dir=" + vendorPath)
		err = runComposerCmd.Run()

		if err != nil {
			//handle no composer error;
			return ""
		}
		overWriteVendorIncluded = true
	}
	//check if vendor exists in plugin
	if !pModel.isVendorIncluded || overWriteVendorIncluded {
		s.WriteString(" vendor lib")
	}

	return s.String()
}
