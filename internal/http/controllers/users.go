package controllers

import (
	application "dev_book_backend/internal/application/userApp"
	domain "dev_book_backend/internal/domain/user"
	"dev_book_backend/pkg/auth"
	"dev_book_backend/pkg/utils"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	app *application.User
}

func NewUserController(repo *application.User) *UserController {
	return &UserController{repo}
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := u.app.GetUsers()

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, users)
}

func (u *UserController) GetUsersById(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)

	id, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	users, err := u.app.GetUserById(id)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, users)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	BodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	var Users domain.User

	if err = json.Unmarshal(BodyRequest, &Users); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = u.app.CreateUsers(Users); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"message": "ok"})
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	BodyRequest, err := ioutil.ReadAll(r.Body)

	query := mux.Vars(r)

	id, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	user_id_toke, err := auth.ExtractUserId(r)

	if err != nil || user_id_toke != id {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	var Users domain.User

	if err = json.Unmarshal(BodyRequest, &Users); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = u.app.UpdateUser(id, Users); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"message": "ok"})

}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)

	id, err := strconv.ParseUint(query["id"], 10, 64)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	user_id_toke, err := auth.ExtractUserId(r)

	if err != nil || user_id_toke != id {
		utils.ResposeError(w, http.StatusUnauthorized, fmt.Errorf("não autorizado"))
		return
	}

	_, err = u.app.DeleteUser(id)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"message": "ok"})
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {

	var Users domain.User

	BodyRequest, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(BodyRequest, &Users); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := u.app.Login(Users)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"authorized": token})
}

func (u *UserController) UpdatePass(w http.ResponseWriter, r *http.Request) {

	var update_pass domain.UpdatePass
	user_id_token, err := auth.ExtractUserId(r)

	if err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	BodyRequest, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(BodyRequest, &update_pass); err != nil {
		utils.ResposeError(w, http.StatusInternalServerError, err)
		return
	}

	err = u.app.UpdatePass(user_id_token, update_pass)

	if err != nil {
		utils.ResposeError(w, http.StatusUnauthorized, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, map[string]interface{}{"message": "alterado com sucesso"})
}
