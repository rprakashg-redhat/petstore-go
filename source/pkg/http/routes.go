package http

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// Route defines an http route
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc routing.Handler
}

// Routes is a collection of http routes supported by this web application
type Routes []Route
