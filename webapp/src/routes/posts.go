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
	{
		URI:         "/posts/{pid}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{pid}/dislike",
		Method:      http.MethodPost,
		Function:    controllers.DislikePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{pid}/edit",
		Method:      http.MethodGet,
		Function:    controllers.LoadEditPage,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{pid}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequireAuth: true,
	},
}
