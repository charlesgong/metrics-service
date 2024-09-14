package main

// works like a Prometheus pusher
// For assignment, it is '2. Data Storage'.
// will push data into Prometheus, wonder about it when large number of servers
import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	initialMetrics()

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9101", nil)
	if err != nil {
		slog.Error("metrics reponse error: %s", err)
	}
}
