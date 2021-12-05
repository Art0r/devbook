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
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		RequireAuth: false,
	},
	{
		URI:         "/user/{id}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}/followers",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowers,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}/following",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowing,
		RequireAuth: true,
	},
	{
		URI:         "/user/{id}/changepassword",
		Method:      http.MethodPost,
		Function:    controllers.UpdatePassword,
		RequireAuth: true,
	},
}
