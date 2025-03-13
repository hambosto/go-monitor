package system

import (
	"github.com/shirou/gopsutil/v4/mem"
)

// GetMEMStats returns memory usage statistics
func GetMEMStats() (mem.VirtualMemoryStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return mem.VirtualMemoryStat{}, err
	}

	return mem.VirtualMemoryStat{
		Total:       v.Total,
		Used:        v.Used,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
		Available:   v.Available,
		Active:      v.Active,
		Buffers:     v.Buffers,
		Cached:      v.Cached,
	}, nil
}
