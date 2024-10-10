package tui

import (
	structures "test-cet-wp-plugin/internal/model/structs"
	models "test-cet-wp-plugin/internal/tui/models"
	"test-cet-wp-plugin/internal/tui/variables"
	"test-cet-wp-plugin/internal/utilities"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(apps *structures.Applications) {
	variables.ApplicationsLists = apps

	model, _ := models.InitParent(apps)

	ParentProgram := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := ParentProgram.Run(); err != nil {
		utilities.HandleFatalError(err, true, "Error running bubble tea")
	}
}
