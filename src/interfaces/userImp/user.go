package interfaces

import "dev_book_backend/src/model"

type User interface {
	CreateNewUser(user model.User) (uint64, error)
	GetAllUsers() ([]model.User, error)
	GetUsersById(id uint64) (*model.User, error)
	GetUsersByEmail(email string) (*model.User, error)
	GetUsersWithPassByid(id uint64) (*model.User, error)
	UpdateUsersById(id uint64, user *model.User) (int, error)
	UpdateUsersPassById(id uint64, new_pass string) (int, error)
	DeleteUser(id uint64) (int, error)
}
