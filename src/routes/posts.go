package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Function:    controllers.SearchPosts,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodGet,
		Function:    controllers.SearchPost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}/posts",
		Method:      http.MethodGet,
		Function:    controllers.SearchByUsers,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}/like",
		Method:      http.MethodPost,
		Function:    controllers.Like,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}/unlike",
		Method:      http.MethodPost,
		Function:    controllers.Unlike,
		RequireAuth: true,
	},
}
