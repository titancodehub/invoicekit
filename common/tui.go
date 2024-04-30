package common

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func RegisterKey(m tea.Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		{
			return m, tea.Quit
		}
	}
	return m, nil
}

func InitFailure() string {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("31")).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("31"))
	return style.Render("failed to initialize command")
}
