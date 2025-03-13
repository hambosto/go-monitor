package system

import (
	"sort"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

// ProcessInfo holds information about a running process
type ProcessInfo struct {
	PID         int32
	Name        string
	Username    string
	Memory      uint64
	CPUPercent  float64
	RunningTime string
}

// GetProcesses returns information about the top n processes sorted by CPU usage
func GetProcesses(n int) ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var processInfos []ProcessInfo
	for _, p := range procs {
		pid := p.Pid
		name, err := p.Name()
		if err != nil {
			name = "Unknown"
		}

		createTime, err := p.CreateTime()
		if err != nil {
			createTime = 0
		}

		startTime := time.Unix(createTime/1000, 0)
		runningTime := time.Since(startTime).Truncate(time.Second)

		username, err := p.Username()
		if err != nil {
			username = "Unknown"
		}

		memoryInfo, err := p.MemoryInfo()
		var memory uint64 = 0
		if err == nil && memoryInfo != nil {
			memory = memoryInfo.RSS
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		processInfos = append(processInfos, ProcessInfo{
			PID:         pid,
			Name:        name,
			RunningTime: runningTime.String(),
			Username:    username,
			Memory:      memory,
			CPUPercent:  cpuPercent,
		})
	}

	// Sort processes by CPU usage (descending)
	sort.Slice(processInfos, func(i, j int) bool {
		return processInfos[i].CPUPercent > processInfos[j].CPUPercent
	})

	// Limit to the top n processes
	if len(processInfos) > n {
		processInfos = processInfos[:n]
	}

	return processInfos, nil
}
