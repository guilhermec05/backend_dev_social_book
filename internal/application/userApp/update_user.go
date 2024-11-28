package application

import (
	domain "dev_book_backend/internal/domain/user"
	"errors"
)

func (u *User) UpdateUser(id uint64, users domain.User) (int, error) {

	_, err := u.repo.GetUsersById(id)

	if err != nil {
		return 0, errors.New("User não existe")
	}

	if err := users.Prepare(""); err != nil {

		return 0, err
	}

	return u.repo.UpdateUsersById(id, &users)
}
