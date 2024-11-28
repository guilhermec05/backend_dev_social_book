package domain

type UserRepository interface {
	CreateNewUser(user User) (uint64, error)
	GetAllUsers() ([]User, error)
	GetUsersById(id uint64) (*User, error)
	GetUsersByEmail(email string) (*User, error)
	GetUsersWithPassByid(id uint64) (*User, error)
	UpdateUsersById(id uint64, user *User) (int, error)
	UpdateUsersPassById(id uint64, new_pass string) (int, error)
	DeleteUser(id uint64) (int, error)
}
