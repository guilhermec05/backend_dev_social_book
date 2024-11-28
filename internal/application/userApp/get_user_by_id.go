package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *User) GetUserById(id uint64) (*domain.User, error) {
	return u.repo.GetUsersById(id)
}
