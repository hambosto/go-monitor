package tui

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hambosto/go-monitor/internal/system"
	"github.com/hambosto/go-monitor/internal/tui/theme"
)

func (m Model) Init() tea.Cmd {
	return tickEvery()
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMessage(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.ProcessTable.Focused() {
				m.TableStyle.Selected = m.BaseStyle
				m.ProcessTable.SetStyles(m.TableStyle)
				m.ProcessTable.Blur()
			} else {
				m.TableStyle.Selected = m.TableStyle.Selected.Background(theme.Colors.Highlight)
				m.ProcessTable.SetStyles(m.TableStyle)
				m.ProcessTable.Focus()
			}
		case "up", "k":
			if m.ProcessTable.Focused() {
				m.ProcessTable.MoveUp(1)
			}
		case "down", "j":
			if m.ProcessTable.Focused() {
				m.ProcessTable.MoveDown(1)
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case TickMessage:
		m.LastUpdate = time.Time(msg)

		cpuStats, err := system.GetCPUStats()
		if err != nil {
			slog.Error("Could not get CPU info", "error", err)
		} else {
			m.CpuUsage = cpuStats
		}

		memStats, err := system.GetMEMStats()
		if err != nil {
			slog.Error("Could not get memory info", "error", err)
		} else {
			m.MemUsage = memStats
		}

		netStats, err := system.GetNetworkStats(&m.NetworkStats)
		if err != nil {
			slog.Error("Could not get network info", "error", err)
		} else {
			m.NetworkStats = netStats
		}

		procs, err := system.GetProcesses(100)
		if err != nil {
			slog.Error("Could not get processes", "error", err)
		} else {
			rows := []table.Row{}
			for _, p := range procs {
				memString, memUnit := system.ConvertBytes(p.Memory)
				rows = append(rows, table.Row{
					fmt.Sprintf("%d", p.PID),
					p.Name,
					fmt.Sprintf("%.2f%%", p.CPUPercent),
					fmt.Sprintf("%s %s", memString, memUnit),
					p.Username,
					p.RunningTime,
				})
			}
			m.ProcessTable.SetRows(rows)
		}

		return m, tickEvery()
	}

	return m, nil
}
