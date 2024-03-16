package healthcheck

import routing "github.com/go-ozzo/ozzo-routing/v2"

// HealthCheckHandler responds to a healthcheck request.
func HealthCheckHandler() routing.Handler {
	return func(c *routing.Context) error {
		return c.Write("OK")
	}
}
