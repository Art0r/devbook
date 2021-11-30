package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		RequireAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		RequireAuth: false,
	},
}
