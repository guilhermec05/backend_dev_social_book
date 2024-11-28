package router

import (
	"database/sql"
	application "dev_book_backend/internal/application/publish"
	"dev_book_backend/internal/http/controllers"
	"dev_book_backend/internal/repository"

	"net/http"
)

func NewPublishRotes(db *sql.DB) []Router {

	publish := repository.NewRepositoryPublish(db)

	appPublish := application.NewAppPublish(publish)

	controllerPublish := controllers.NewPublishController(appPublish)

	return []Router{
		{
			URI:         "/MyPublish",
			Method:      http.MethodPost,
			Function:    controllerPublish.PublicPublish,
			RequireAuth: true,
		},
		{
			URI:         "/MyPublish",
			Method:      http.MethodGet,
			Function:    controllerPublish.MyPublish,
			RequireAuth: true,
		},
		{
			URI:         "/MyPublish/{id:[0-9]+}",
			Method:      http.MethodGet,
			Function:    controllerPublish.MyPublishById,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/{id:[0-9]+}",
			Method:      http.MethodGet,
			Function:    controllerPublish.PublishById,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/follow",
			Method:      http.MethodGet,
			Function:    controllerPublish.PublishFollow,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/{id}/Edit",
			Method:      http.MethodPut,
			Function:    controllerPublish.EditPublish,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/{id}/Delete",
			Method:      http.MethodDelete,
			Function:    controllerPublish.DeletePublish,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/{id}/post",
			Method:      http.MethodPost,
			Function:    controllerPublish.EditPublish,
			RequireAuth: true,
		},

		{
			URI:         "/Publish/{id:[0-9]+}/like",
			Method:      http.MethodPost,
			Function:    controllerPublish.LikePublish,
			RequireAuth: true,
		},

		{
			URI:         "/Publish/{id:[0-9]+}/like",
			Method:      http.MethodGet,
			Function:    controllerPublish.GetLikePublish,
			RequireAuth: true,
		},
		{
			URI:         "/Publish/{id:[0-9]+}/unlike",
			Method:      http.MethodDelete,
			Function:    controllerPublish.UnLikePublish,
			RequireAuth: true,
		},
	}
}
