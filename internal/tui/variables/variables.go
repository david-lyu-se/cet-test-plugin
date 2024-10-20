package variables

import (
	"os"
	structures "test-cet-wp-plugin/internal/model/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	Nav Mode = iota
	Edit
	Create
)

var (
	//os variables
	Conf *structures.ConfFile
	File *os.File
	//Model variables
	ParentProgram *tea.Program
	ParentModel   *tea.Model
	AppModel      *tea.Model
	PluginModel   *tea.Model
	//Application variables
	AppInfo    structures.Application
	PluginPath string
	//Window Variables
	WindowSize tea.WindowSizeMsg
)

/* Tea.Msg */
// Used to like inner tea.WindowSizeMsg. Bug with overlap with ParentModel
type InitAppModel struct{}
