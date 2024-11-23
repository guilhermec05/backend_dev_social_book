package application

import (
	"errors"
)

func (u *User) DeleteUser(id uint64) (int, error) {

	_, err := u.repo.GetUsersById(id)

	if err != nil {
		return 0, errors.New("User não existe")
	}

	return u.repo.DeleteUser(id)
}
