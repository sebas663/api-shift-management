package routes

import (
	h "../handlers"
	m "../models"
)

//RootRoutes index route.
var RootRoutes = m.Routes{
	m.Route{
		"Index",
		"GET",
		"/",
		h.Index,
	},
}
