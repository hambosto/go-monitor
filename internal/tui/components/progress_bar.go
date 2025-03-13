package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/hambosto/go-monitor/internal/tui/theme"
)

func RenderProgressBar(percentage float64, baseStyle lipgloss.Style) string {
	totalBars := 20
	fillBars := int(percentage / 100 * float64(totalBars))

	if fillBars > totalBars {
		fillBars = totalBars
	} else if fillBars < 0 {
		fillBars = 0
	}

	filled := baseStyle.
		Foreground(theme.Colors.Green).
		Render(strings.Repeat("|", fillBars))

	empty := baseStyle.
		Foreground(theme.Colors.Secondary).
		Render(strings.Repeat("|", totalBars-fillBars))

	return baseStyle.Render(fmt.Sprintf("%s%s%s%s", "[", filled, empty, "]"))
}
