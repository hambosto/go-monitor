package system

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

// CPUStats returns CPU usage statistics as percentages
func GetCPUStats() (cpu.TimesStat, error) {
	stats, err := cpu.Times(false)
	if err != nil {
		return cpu.TimesStat{}, err
	}
	if len(stats) == 0 {
		return cpu.TimesStat{}, nil
	}

	currentStats := stats[0]

	total := currentStats.User + currentStats.System + currentStats.Idle + currentStats.Nice +
		currentStats.Iowait + currentStats.Irq + currentStats.Softirq + currentStats.Steal +
		currentStats.Guest

	if total == 0 {
		return cpu.TimesStat{}, nil
	}

	// Convert raw values to percentages
	currentStats.User = (currentStats.User / total) * 100
	currentStats.System = (currentStats.System / total) * 100
	currentStats.Idle = (currentStats.Idle / total) * 100
	currentStats.Nice = (currentStats.Nice / total) * 100
	currentStats.Iowait = (currentStats.Iowait / total) * 100
	currentStats.Irq = (currentStats.Irq / total) * 100
	currentStats.Softirq = (currentStats.Softirq / total) * 100
	currentStats.Steal = (currentStats.Steal / total) * 100
	currentStats.Guest = (currentStats.Guest / total) * 100

	return currentStats, nil
}
