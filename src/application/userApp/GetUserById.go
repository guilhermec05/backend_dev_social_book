package application

import (
	"dev_book_backend/src/model"
)

func (u *User) GetUserById(id uint64) (*model.User, error) {
	return u.repo.GetUsersById(id)
}
