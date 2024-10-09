package variables

import (
	"test-cet-wp-plugin/internal/model/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	Nav Mode = iota
	Edit
	Create
)

var (
	ParentProgram   *tea.Program
	EnvironmentsRef *structs.Environments
	WindowSize      tea.WindowSizeMsg
)
