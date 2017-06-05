package models

import "net/http"

type (
	//Route the resources.
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
	//Routes array.
	Routes []Route
)
