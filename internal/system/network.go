package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/net"
)

type NetworkStats struct {
	BytesSent    uint64
	BytesRecv    uint64
	SendRate     float64
	ReceiveRate  float64
	PreviousSent uint64
	PreviousRecv uint64
	LastMeasured time.Time
}

func GetNetworkStats(previousStats *NetworkStats) (NetworkStats, error) {
	counters, err := net.IOCounters(false)
	if err != nil {
		return NetworkStats{}, err
	}

	if len(counters) == 0 {
		return NetworkStats{}, fmt.Errorf("no network interfaces found")
	}

	totalIO := counters[0]

	now := time.Now()
	newStats := NetworkStats{
		BytesSent:    totalIO.BytesSent,
		BytesRecv:    totalIO.BytesRecv,
		LastMeasured: now,
	}

	if previousStats != nil && !previousStats.LastMeasured.IsZero() {
		timeDiff := now.Sub(previousStats.LastMeasured).Seconds()
		if timeDiff > 0 {
			bytesSentDiff := totalIO.BytesSent - previousStats.BytesSent
			bytesRecvDiff := totalIO.BytesRecv - previousStats.BytesRecv

			newStats.SendRate = float64(bytesSentDiff) / timeDiff
			newStats.ReceiveRate = float64(bytesRecvDiff) / timeDiff
		}
	}

	return newStats, nil
}
