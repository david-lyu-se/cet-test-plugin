package tui

import (
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/menu"
	"test-cet-wp-plugin/internal/tui/variables"
	"test-cet-wp-plugin/internal/utilities"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTea(conf *structures.ConfFile) {
	variables.Conf = conf

	model, _ := menu.InitMenu()
	variables.ParentModel = &model
	variables.ParentProgram = tea.NewProgram(model, tea.WithAltScreen())

	if _, err := variables.ParentProgram.Run(); err != nil {
		utilities.HandleFatalError(err, true, "Error running bubble tea")
	}
}
