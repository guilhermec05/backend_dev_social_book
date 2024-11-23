package application

import (
	"dev_book_backend/src/model"
)

func (u *User) GetUsers() ([]model.User, error) {
	return u.repo.GetAllUsers()
}
