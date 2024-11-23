package application

import interfaces "dev_book_backend/src/interfaces/userImp"

type User struct {
	repo interfaces.User
}

func NewAppUser(repo interfaces.User) *User {
	return &User{repo}
}
