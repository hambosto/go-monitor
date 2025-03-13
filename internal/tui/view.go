package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/hambosto/go-monitor/internal/tui/components"
)

func (m Model) View() string {
	column := m.BaseStyle.Width(m.Width).Padding(1, 0, 0, 0).Render

	content := m.BaseStyle.
		Width(m.Width).
		Height(m.Height).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				column(m.renderHeader()),
				column(m.renderProcessTable()),
			),
		)

	return content
}

func (m Model) renderHeader() string {
	timeSinceUpdate := fmt.Sprintf("Last update: %d milliseconds ago\n",
		time.Since(m.LastUpdate).Microseconds())

	header := components.RenderHeader(
		m.BaseStyle,
		m.ViewStyle,
		m.CpuUsage,
		m.MemUsage,
		m.NetworkStats,
	)

	return m.ViewStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			timeSinceUpdate,
			header,
		),
	)
}

func (m Model) renderProcessTable() string {
	return m.ViewStyle.Render(m.ProcessTable.View())
}
