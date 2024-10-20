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
	ParentProgram *tea.Program
	AppProgram    *tea.Program
	PluginProgram *tea.Program
	Conf          *structures.ConfFile
	WindowSize    tea.WindowSizeMsg
	File          *os.File
)
