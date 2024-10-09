package variables

import "github.com/charmbracelet/lipgloss"

/* Styling */

var DocStyle = lipgloss.NewStyle().Margin(0, 2)

var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render
