package application

import (
	domain "dev_book_backend/internal/domain/user"
)

type User struct {
	repo domain.UserRepository
}

func NewAppUser(repo domain.UserRepository) *User {
	return &User{repo}
}
