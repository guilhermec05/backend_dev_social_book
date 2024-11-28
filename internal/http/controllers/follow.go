package controllers

import (
	application "dev_book_backend/internal/application/followApp"
	"dev_book_backend/pkg/auth"
	"dev_book_backend/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FollowController struct {
	app *application.Follow
}

func NewFollowController(repo *application.Follow) *FollowController {
	return &FollowController{repo}
}

func (u *FollowController) SetFollow(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	user_id, err := strconv.ParseUint(query["user_id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	user_id_token, err := auth.ExtractUserId(r)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	_, err = u.app.FollowUser(user_id_token, user_id)

	if err != nil {
		utils.ResposeError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "ok"})
}

func (u *FollowController) UnFollow(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	user_id, err := strconv.ParseUint(query["user_id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	user_id_token, err := auth.ExtractUserId(r)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	_, err = u.app.UnFollowUser(user_id_token, user_id)

	if err != nil {
		utils.ResposeError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "ok"})
}

func (u *FollowController) GetFollowMe(w http.ResponseWriter, r *http.Request) {

	user_id_token, err := auth.ExtractUserId(r)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	users, err := u.app.GetFollowMe(user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, users)
}

func (u *FollowController) GetFollow(w http.ResponseWriter, r *http.Request) {

	user_id_token, err := auth.ExtractUserId(r)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	users, err := u.app.GetFollow(user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, users)
}
