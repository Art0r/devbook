package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:         "/createuser",
		Method:      http.MethodGet,
		Function:    controllers.LoadUserSigninPage,
		RequireAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
}
