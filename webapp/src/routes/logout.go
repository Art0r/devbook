package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRoutes = Route{
	URI:         "/logout",
	Method:      http.MethodGet,
	Function:    controllers.Logout,
	RequireAuth: true,
}
