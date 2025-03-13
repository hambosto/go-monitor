package components

import (
	"github.com/charmbracelet/bubbles/table"
)

func NewProcessTable(tableStyle table.Styles) table.Model {
	return table.New(
		table.WithColumns([]table.Column{
			{Title: "PID", Width: 10},
			{Title: "Name", Width: 25},
			{Title: "CPU", Width: 12},
			{Title: "MEM", Width: 12},
			{Title: "Username", Width: 12},
			{Title: "Time", Width: 12},
		}),

		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(20),
		table.WithStyles(tableStyle),
	)
}
