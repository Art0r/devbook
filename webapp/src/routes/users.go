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
	{
		URI:         "/searchusers",
		Method:      http.MethodGet,
		Function:    controllers.LoadUserPage,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    controllers.LoadUserProfile,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/profile",
		Method:      http.MethodGet,
		Function:    controllers.LoadLoggedUserProfile,
		RequireAuth: true,
	},
	{
		URI:         "/edit-user",
		Method:      http.MethodGet,
		Function:    controllers.LoadLoggedUserProfileEdit,
		RequireAuth: true,
	},
	{
		URI:         "/edit-user",
		Method:      http.MethodPut,
		Function:    controllers.EditUser,
		RequireAuth: true,
	},
	{
		URI:         "/update-password",
		Method:      http.MethodGet,
		Function:    controllers.LoadUpdatePasswordPage,
		RequireAuth: true,
	},
	{
		URI:         "/update-password",
		Method:      http.MethodPost,
		Function:    controllers.UpdatePassword,
		RequireAuth: true,
	},
	{
		URI:         "/delete-user",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
}
