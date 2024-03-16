package http

import (
	health "github.com/rprakashg-redhat/petstore-go/internal/healthcheck"
	"github.com/rprakashg-redhat/petstore-go/internal/metrics"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func newRouter(routes Routes) *routing.Router {
	router := routing.New()
	rg := router.Group("/v1")

	for _, route := range routes {
		rg.To(route.Method, route.Path, route.HandlerFunc)
	}

	//Add health check
	router.To("GET", "/health", health.HealthCheckHandler())

	//add metrics
	router.To("GET", "/metrics", routing.HTTPHandler(metrics.PromMetricsHandler()))

	return router
}
