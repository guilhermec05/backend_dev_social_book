package application

import (
	domain "dev_book_backend/internal/domain/user"
	"dev_book_backend/pkg/auth"
	"dev_book_backend/pkg/utils"
	"errors"
)

func (u *User) Login(Users domain.User) (string, error) {

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
	return auth.CreatedToken(uint(login.ID))
}
