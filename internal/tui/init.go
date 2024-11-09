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

	// //Test
	// app := structures.Application{Name: "Test", Path: "/", PluginPath: "/"}
	// conf.Apps = append(conf.Apps, app)
	// app = structures.Application{Name: "Test1", Path: "/", PluginPath: "/"}
	// conf.Apps = append(conf.Apps, app)

	model, _ := models.InitParent()

	variables.ParentModel = &model
	variables.ParentProgram = tea.NewProgram(model, tea.WithAltScreen())

	if _, err := variables.ParentProgram.Run(); err != nil {
		utilities.HandleFatalError(err, true, "Error running bubble tea")
	}
}
