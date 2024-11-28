package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *User) GetUsers() ([]domain.User, error) {
	return u.repo.GetAllUsers()
}
