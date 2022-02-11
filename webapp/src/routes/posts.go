package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesPosts = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequireAuth: true,
	},
}
