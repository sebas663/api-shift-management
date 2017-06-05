package main

import (
	m "./models"
	r "./routes"
)

var routes = m.Routes{}

// init Give us some seed data
func init() {
	routes = append(routes, r.RootRoutes...)
	routes = append(routes, r.ApointmentRoutes...)
}
