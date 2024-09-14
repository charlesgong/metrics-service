package main

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Metrics struct {
	CPUUsage    float64 `json:CPUUsage`
	MemoryUsage float64 `json:MemoryUsage`
	DiskIO      float64 `json:DiskIO`
	NetworkIO   float64 `json:NetworkIO`
	Time        int64
}

// Metrics Collection
func collectMetrics() Metrics {
	cpuUsage, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.IOCounters()
	netInfo, _ := net.IOCounters(false)

	return Metrics{
		CPUUsage:    cpuUsage[0],
		MemoryUsage: float64(memInfo.Used / memInfo.Total),
		DiskIO:      float64(diskInfo["sda"].WriteBytes),
		NetworkIO:   float64(netInfo[0].BytesSent),
		Time:        time.Now().UTC().Unix(),
	}
}
