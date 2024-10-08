package tui

import (
	"test-cet-wp-plugin/internal/model/structs"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	InitalProgramRef *tea.Program
	EnvironmentsRef  *structs.Environments
	WindowSize       tea.WindowSizeMsg
)
