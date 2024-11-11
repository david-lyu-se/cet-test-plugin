package variables

import "github.com/charmbracelet/lipgloss"

/* Styling */

var DocStyle = lipgloss.NewStyle().Padding(0, 2)

// Foreground Purple, bold,
var TitleStyle = lipgloss.NewStyle().
	Bold(true).Foreground(lipgloss.Color("99")).Padding(1, 0, 1, 0).Render

var UserChoiceStyle = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder(), true).
	BorderForeground(lipgloss.Color("#228B22")).
	Padding(2).
	MarginBottom(1).
	Foreground(lipgloss.Color("#228B22")).
	Render

var ModelSelectStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.Color("99")).
	Foreground(lipgloss.Color("99")).
	PaddingLeft(1).
	Margin(1, 0).
	Render

var ModelChoiceStyle = lipgloss.NewStyle().
	Render

var ModelChoiceContainerStyle = lipgloss.NewStyle().
	Render

var BlinkingStyle = lipgloss.NewStyle().Blink(true)

var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render
