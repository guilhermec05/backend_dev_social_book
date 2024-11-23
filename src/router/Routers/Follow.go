package router

import (
	"database/sql"
	followapp "dev_book_backend/src/application/followApp"
	"dev_book_backend/src/controllers"
	"dev_book_backend/src/repository"
	"net/http"
)

func NewFollowRotes(db *sql.DB) []Router {

	follow := repository.NewRepositoryFollow(db)

	appFollow := followapp.NewAppFollow(follow)

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
			URI:         "/users/get_follow_me",
			Method:      http.MethodGet,
			Function:    controllerFollow.GetFollowMe,
			RequireAuth: true,
		},

		{
			URI:         "/users/get_follow",
			Method:      http.MethodGet,
			Function:    controllerFollow.GetFollow,
			RequireAuth: true,
		},
	}
}
