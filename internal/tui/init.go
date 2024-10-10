package tui

import (
	structures "test-cet-wp-plugin/internal/model/structs"
	models "test-cet-wp-plugin/internal/tui/models"
	"test-cet-wp-plugin/internal/tui/variables"
	"test-cet-wp-plugin/internal/utilities"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(conf *structures.ConfFile) {
	variables.Conf = conf

	model, _ := models.InitParent(conf.Apps)

	ParentProgram := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := ParentProgram.Run(); err != nil {
		utilities.HandleFatalError(err, true, "Error running bubble tea")
	}
}
