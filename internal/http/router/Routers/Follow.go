package router

import (
	"database/sql"
	application "dev_book_backend/internal/application/followApp"
	"dev_book_backend/internal/http/controllers"
	"dev_book_backend/internal/repository"

	"net/http"
)

func NewFollowRotes(db *sql.DB) []Router {

	follow := repository.NewRepositoryFollow(db)

	appFollow := application.NewAppFollow(follow)

	controllerFollow := controllers.NewFollowController(appFollow)

	return []Router{
		{
			URI:         "/users/{user_id}/follow",
			Method:      http.MethodPost,
			Function:    controllerFollow.SetFollow,
			RequireAuth: true,
		},
		{
			URI:         "/users/{user_id}/unfollow",
			Method:      http.MethodPost,
			Function:    controllerFollow.UnFollow,
			RequireAuth: true,
		},
		{
			URI:         "/users/follow_me",
			Method:      http.MethodGet,
			Function:    controllerFollow.GetFollowMe,
			RequireAuth: true,
		},

		{
			URI:         "/users/follow",
			Method:      http.MethodGet,
			Function:    controllerFollow.GetFollow,
			RequireAuth: true,
		},
	}
}
