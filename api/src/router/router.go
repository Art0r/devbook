package router

import (
	"devbook-api/src/routes"

	"github.com/gorilla/mux"
)

// Return a router with all the configured routes
func Routes() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
