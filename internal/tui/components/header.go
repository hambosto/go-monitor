package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/hambosto/go-monitor/internal/system"
	"github.com/hambosto/go-monitor/internal/tui/theme"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func RenderHeader(baseStyle, viewStyle lipgloss.Style, cpuUsage cpu.TimesStat, memUsage mem.VirtualMemoryStat, netStats system.NetworkStats) string {
	list := baseStyle.
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(theme.Colors.Border).
		Height(4).
		Padding(0, 1)

	listHeader := baseStyle.Bold(true).Render

	return lipgloss.JoinHorizontal(lipgloss.Top,
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("% Usage"),
				renderListItem(baseStyle, "CPU", fmt.Sprintf("%s %.1f", RenderProgressBar(100-cpuUsage.Idle, baseStyle), 100-cpuUsage.Idle), "%"),
				renderListItem(baseStyle, "MEM", fmt.Sprintf("%s %.1f", RenderProgressBar(memUsage.UsedPercent, baseStyle), memUsage.UsedPercent), "%"),
			),
		),

		list.Border(lipgloss.NormalBorder(), false).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("CPU"),
				renderListItem(baseStyle, "user", fmt.Sprintf("%.1f", cpuUsage.User), "%"),
				renderListItem(baseStyle, "sys", fmt.Sprintf("%.1f", cpuUsage.System), "%"),
				renderListItem(baseStyle, "idle", fmt.Sprintf("%.1f", cpuUsage.Idle), "%"),
			),
		),
		list.Border(lipgloss.NormalBorder(), false).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader(""),
				renderListItem(baseStyle, "nice", fmt.Sprintf("%.1f", cpuUsage.Nice), "%"),
				renderListItem(baseStyle, "iowait", fmt.Sprintf("%.1f", cpuUsage.Iowait), "%"),
				renderListItem(baseStyle, "irq", fmt.Sprintf("%.1f", cpuUsage.Irq), "%"),
			),
		),
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader(""),
				renderListItem(baseStyle, "softirq", fmt.Sprintf("%.1f", cpuUsage.Softirq), "%"),
				renderListItem(baseStyle, "steal", fmt.Sprintf("%.1f", cpuUsage.Steal), "%"),
				renderListItem(baseStyle, "guest", fmt.Sprintf("%.1f", cpuUsage.Guest), "%"),
			),
		),

		list.Border(lipgloss.NormalBorder(), false).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("MEM"),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Total)
					return renderListItem(baseStyle, "total", value, unit)
				}(),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Used)
					return renderListItem(baseStyle, "used", value, unit)
				}(),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Available)
					return renderListItem(baseStyle, "free", value, unit)
				}(),
			),
		),
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader(""),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Active)
					return renderListItem(baseStyle, "active", value, unit)
				}(),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Buffers)
					return renderListItem(baseStyle, "buffers", value, unit)
				}(),
				func() string {
					value, unit := system.ConvertBytes(memUsage.Cached)
					return renderListItem(baseStyle, "cached", value, unit)
				}(),
			),
		),

		list.Border(lipgloss.NormalBorder(), false).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("NET"),
				func() string {
					value, unit := system.ConvertBytes(netStats.BytesSent)
					return renderListItem(baseStyle, "sent", value, unit)
				}(),
				func() string {
					value, unit := system.ConvertBytes(netStats.BytesRecv)
					return renderListItem(baseStyle, "recv", value, unit)
				}(),
			),
		),
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader(""),
				func() string {
					value, unit := system.ConvertByteRate(netStats.SendRate)
					return renderListItem(baseStyle, "tx", value, unit+"/s")
				}(),
				func() string {
					value, unit := system.ConvertByteRate(netStats.ReceiveRate)
					return renderListItem(baseStyle, "rx", value, unit+"/s")
				}(),
			),
		),
	)
}

func renderListItem(baseStyle lipgloss.Style, key, value string, suffix ...string) string {
	finalSuffix := ""
	if len(suffix) > 0 {
		finalSuffix = suffix[0]
	}

	listItemValue := baseStyle.Align(lipgloss.Right).Render(fmt.Sprintf("%s%s", value, finalSuffix))
	listItemKey := baseStyle.Render(key + ":")

	return fmt.Sprintf("%s %s", listItemKey, listItemValue)
}
