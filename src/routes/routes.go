package routes

import (
	"devbook-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	userRoutes := usersRoutes
	loginRoute := loginRoute

	for _, route := range userRoutes {

		if route.RequireAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	r.HandleFunc(loginRoute.URI, loginRoute.Function).Methods(loginRoute.Method)

	return r
}
