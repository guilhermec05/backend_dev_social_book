package application

import (
	"dev_book_backend/src/model"
	"dev_book_backend/src/utils"
	"errors"
)

func (u *User) Login(Users model.User) (string, error) {

	pass := Users.Password

	if err := Users.ValidateLogin(); err != nil {

		return "", err
	}

	login, err := u.repo.GetUsersByEmail(Users.Email)

	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(pass, login.Password) {
		return "", errors.New("A senha não correponde")
	}
	return utils.CreatedToken(uint(login.ID))
}
