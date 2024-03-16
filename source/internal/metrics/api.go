package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	reg *prometheus.Registry
)

func PromMetricsHandler() http.Handler {
	// Create non-global registry.
	reg = prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
}
