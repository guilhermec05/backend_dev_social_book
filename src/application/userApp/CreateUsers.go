package application

import (
	"dev_book_backend/src/model"
)

func (u *User) CreateUsers(Users model.User) (uint64, error) {

	if err := Users.Prepare("inserted"); err != nil {

		return 0, err
	}

	return u.repo.CreateNewUser(Users)
}
