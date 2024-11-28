package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *User) CreateUsers(Users domain.User) (uint64, error) {

	if err := Users.Prepare("inserted"); err != nil {

		return 0, err
	}

	return u.repo.CreateNewUser(Users)
}
