package routes

import "net/http"

type Route struct {
	Name        string
	Method      string
	Queries     string
	Hadle       string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route
