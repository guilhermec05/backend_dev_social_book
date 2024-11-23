package interfaces

import "dev_book_backend/src/model"

type Follow interface {
	FollowUser(user_id uint64, follow_id uint64) (int, error)
	UnFollowUser(user_id uint64, follow_id uint64) (int, error)
	GetFollows(user_id uint64) ([]model.User, error)
	GetFollowMe(user_id uint64) ([]model.User, error)
}
