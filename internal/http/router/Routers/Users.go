package router

import (
	"database/sql"
	application "dev_book_backend/internal/application/userApp"
	"dev_book_backend/internal/http/controllers"
	"dev_book_backend/internal/repository"

	"net/http"
)

func NewControllerUser(db *sql.DB) []Router {

	User := repository.NewRepositoryUser(db)

	appUser := application.NewAppUser(User)

	controller := controllers.NewUserController(appUser)

	return []Router{
		{
			URI:         "/users",
			Method:      http.MethodGet,
			Function:    controller.GetUsers,
			RequireAuth: true,
		},

		{
			URI:         "/users/{id:[0-9]+}",
			Method:      http.MethodGet,
			Function:    controller.GetUsersById,
			RequireAuth: true,
		},
		{
			URI:         "/users",
			Method:      http.MethodPost,
			Function:    controller.CreateUser,
			RequireAuth: false,
		},
		{
			URI:         "/users/{id:[0-9]+}",
			Method:      http.MethodPut,
			Function:    controller.UpdateUser,
			RequireAuth: true,
		},
		{
			URI:         "/users/{id:[0-9]+}",
			Method:      http.MethodDelete,
			Function:    controller.DeleteUser,
			RequireAuth: true,
		},
		{
			URI:         "/users/update_pass",
			Method:      http.MethodPost,
			Function:    controller.UpdatePass,
			RequireAuth: true,
		},
		{
			URI:         "/login",
			Method:      http.MethodPost,
			Function:    controller.Login,
			RequireAuth: false,
		},
	}
}
