package controllers

import (
	application "dev_book_backend/src/application/publish"
	"dev_book_backend/src/model"
	"dev_book_backend/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PublishController struct {
	app *application.Publish
}

func NewPublishController(repo *application.Publish) *PublishController {
	return &PublishController{repo}
}

func (u *PublishController) PublicPublish(w http.ResponseWriter, r *http.Request) {

	var update_publish model.Publish
	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	BodyRequest, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(BodyRequest, &update_publish); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	update_publish.Author_id = user_id_token

	err = u.app.PublicPublish(update_publish)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"message": "Criada com sucesso"})
}

func (u *PublishController) MyPublish(w http.ResponseWriter, r *http.Request) {

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	MyPublish, err := u.app.GetMyPublish(user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, MyPublish)
}

func (u *PublishController) MyPublishById(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	MyPublish, err := u.app.GetMyPublishById(id_publish, user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, MyPublish)
}

func (u *PublishController) EditPublish(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)
	update_publish := model.Publish{}

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	BodyRequest, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(BodyRequest, &update_publish); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	err = u.app.EditPublish(update_publish, id_publish, user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "Alterado com sucesso"})
}

func (u *PublishController) DeletePublish(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	err = u.app.DeletePublish(id_publish, user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "Deletado com sucesso"})
}

func (u *PublishController) PublishById(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	MyPublish, err := u.app.GetPublishById(id_publish)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, MyPublish)
}

func (u *PublishController) PublishFollow(w http.ResponseWriter, r *http.Request) {

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	MyPublish, err := u.app.GetPublishByFollows(user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, MyPublish)
}

func (u *PublishController) LikePublish(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	err = u.app.LikePublish(id_publish, user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "Publicação curtida com sucesso"})
}

func (u *PublishController) GetLikePublish(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	users, err := u.app.GetLikePublish(id_publish)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, users)
}

func (u *PublishController) UnLikePublish(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id_publish, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return

	}

	user_id_token, err := utils.ExtractUserId(r)
	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	err = u.app.UnLikePublish(id_publish, user_id_token)

	if err != nil {
		utils.ResposeError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"message": "Publicação descurtida com sucesso"})
}
