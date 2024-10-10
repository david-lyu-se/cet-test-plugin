package variables

import (
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
	ParentProgram     *tea.Program
	ApplicationsLists *structures.Applications
	WindowSize        tea.WindowSizeMsg
)
