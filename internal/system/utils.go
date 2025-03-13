package system

import (
	"fmt"
)

func ConvertBytes(bytes uint64) (string, string) {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f", float64(bytes)/float64(GB)), "GB"
	case bytes >= MB:
		return fmt.Sprintf("%.2f", float64(bytes)/float64(MB)), "MB"
	case bytes >= KB:
		return fmt.Sprintf("%.2f", float64(bytes)/float64(KB)), "KB"
	default:
		return fmt.Sprintf("%d", bytes), "B"
	}
}

func ConvertByteRate(bytesPerSecond float64) (string, string) {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	var value float64
	var unit string

	switch {
	case bytesPerSecond >= TB:
		value = bytesPerSecond / TB
		unit = "TB"
	case bytesPerSecond >= GB:
		value = bytesPerSecond / GB
		unit = "GB"
	case bytesPerSecond >= MB:
		value = bytesPerSecond / MB
		unit = "MB"
	case bytesPerSecond >= KB:
		value = bytesPerSecond / KB
		unit = "KB"
	default:
		value = bytesPerSecond
		unit = "B"
	}

	return fmt.Sprintf("%.2f", value), unit
}
