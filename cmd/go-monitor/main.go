package main

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hambosto/go-monitor/internal/tui"
	"github.com/hambosto/go-monitor/internal/tui/theme"
)

func main() {
	tableStyle := table.DefaultStyles()
	tableStyle.Selected = lipgloss.NewStyle().Background(theme.Colors.Highlight)

	m := tui.NewModel(tableStyle)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
