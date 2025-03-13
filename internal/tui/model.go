package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/hambosto/go-monitor/internal/system"
	"github.com/hambosto/go-monitor/internal/tui/components"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Model struct {
	Width        int
	Height       int
	LastUpdate   time.Time
	ProcessTable table.Model
	TableStyle   table.Styles
	BaseStyle    lipgloss.Style
	ViewStyle    lipgloss.Style
	CpuUsage     cpu.TimesStat
	MemUsage     mem.VirtualMemoryStat
	NetworkStats system.NetworkStats
}

func NewModel(tableStyle table.Styles) Model {
	processTable := components.NewProcessTable(tableStyle)

	return Model{
		ProcessTable: processTable,
		TableStyle:   tableStyle,
		BaseStyle:    lipgloss.NewStyle(),
		ViewStyle:    lipgloss.NewStyle(),
		LastUpdate:   time.Now(),
	}
}

type TickMessage time.Time
