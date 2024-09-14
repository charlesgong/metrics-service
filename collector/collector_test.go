package main

import (
	"fmt"
	"testing"
)

var TEN_BILLION = float64(9999999999)

func TestCollectMetrics(t *testing.T) {
	metrics := collectMetrics()
	if metrics.CPUUsage > 100.0 || metrics.CPUUsage < 0.0 {
		t.Error("CPU usage error")
	}

	if metrics.MemoryUsage > 100.0 || metrics.CPUUsage < 0.0 {
		t.Error("memory usage error")
	}

	if metrics.DiskIO > TEN_BILLION {
		t.Error("disk usage error")
	}

	if metrics.NetworkIO > TEN_BILLION {
		t.Error("network usage error")
	}

	fmt.Printf("CPU: %.2f%%, Memory: %.2f%%, Disk I/O: %f, Network I/O: %f\n",
		metrics.CPUUsage, metrics.MemoryUsage, metrics.DiskIO, metrics.NetworkIO)

}
