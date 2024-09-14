package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"github.com/juju/fslock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	cpuUsageGauge    = prometheus.NewGauge(prometheus.GaugeOpts{Name: "cpu_usage"})
	memoryUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "memory_usage"})
	diskIOGauge      = prometheus.NewGauge(prometheus.GaugeOpts{Name: "disk_io"})
	networkIOGauge   = prometheus.NewGauge(prometheus.GaugeOpts{Name: "network_io"})

	PUSH_INTERVAL = 10 * time.Second
	MAX_CACHE     = int64(500000)
)

func initialMetrics() {
	prometheus.MustRegister(cpuUsageGauge, memoryUsageGauge, diskIOGauge, networkIOGauge)
}

func pushMetrics() {

	metrics := collectMetrics()
	cpuUsageGauge.Set(metrics.CPUUsage)
	memoryUsageGauge.Set(metrics.MemoryUsage)
	diskIOGauge.Set(metrics.DiskIO)
	networkIOGauge.Set(metrics.NetworkIO)

	if err := push.New("http://192.168.68.53:9091/", "instance").
		Collector(cpuUsageGauge).
		Collector(memoryUsageGauge).
		Collector(diskIOGauge).
		Collector(networkIOGauge).
		Grouping("server", "performance").
		Push(); err != nil {
		slog.Error("Could not push to Pushgateway:", err)
		saveToLocalCache(metrics)
	} else {

		slog.Info("Metrics pushed")
	}
}

func saveToLocalCache(metrics Metrics) {
	f, err := os.OpenFile("backfill.lock", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		slog.Error("Error lock file: %v\n", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if info.Size() > MAX_CACHE {
		slog.Warn("Cache file reached max size ")
		return
	}

	json_metrics, err := json.Marshal(metrics)

	if err != nil {
		slog.Error("Error metrics format %v\n", err)
	}
	f.Write(json_metrics)
	f.WriteString("\n")
}

func recordMetrics() {
	go func() {
		for {
			pushMetrics()
			time.Sleep(PUSH_INTERVAL)
		}
	}()
}

func backfill() {
	var lockfile = fslock.New("backfill.lock")
	if err := lockfile.TryLock(); err != nil {
		slog.Error("Error lock file: %v\n", err)
		return
	}
	defer lockfile.Unlock()
	// TODO push backfill prometheus data after recovery
	/* promdump -url=http://localhost:9090 \
		 -metric='node_filesystem_free{job="node"}' \
	     -out=backfill.lock -batch=12h -batches_per_file=2 -period=8760h
	*/
}
