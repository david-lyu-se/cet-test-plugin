package variables

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	Create key.Binding
	Enter  key.Binding
	Edit   key.Binding
	Delete key.Binding
	Back   key.Binding
	Quit   key.Binding
	Down   key.Binding
	Up     key.Binding
	Toggle key.Binding
	Pick   key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("j", "down"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("k", "up"),
	),
	Toggle: key.NewBinding(
		key.WithKeys("toggle", "t"),
		key.WithHelp("t", "toggle"),
	),
	Pick: key.NewBinding(
		key.WithKeys("pick","p"),
		key.WithHelp("p","pick")
	),
}

func FilePickerKeyHelper(extra string = "") string {
	var s strings.Builder

	s.WriteString("Keymaps: \n")
	s.WriteString("Enter - Pick File; Left/Right - Move in/out file; Up/Down; ")
	s.WriteString(extra)
	return s.String()
}
