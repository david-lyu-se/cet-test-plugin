package variables

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

/* Handle Tea textInput.Model Inputs */
func TextInputs(msg tea.KeyMsg, model *textinput.Model) tea.Cmd {
	if key.Matches(msg, Keymap.Quit) {
		model.SetValue("")
		model.Blur()
	}
	if key.Matches(msg, Keymap.Back) {
		model.SetValue("")
		model.Blur()
	}
	var _, cmd = model.Update(msg)
	return cmd
}
