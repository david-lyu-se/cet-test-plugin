package tui

import (
	bubbleteaModel "test-cet-wp-plugin/internal/model/bubbletea"
	"test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"
	"test-cet-wp-plugin/internal/utilities"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(Environments *structs.Environments) {
	variables.EnvironmentsRef = Environments

	model, _ := bubbleteaModel.InitParent(variables.EnvironmentsRef)

	ParentProgram := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := ParentProgram.Run(); err != nil {
		utilities.HandleFatalError(err, true, "Error running bubble tea")
	}
}
