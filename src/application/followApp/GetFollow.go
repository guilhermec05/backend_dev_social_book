package application

import (
	"dev_book_backend/src/model"
)

func (u *Follow) GetFollow(user_id uint64) ([]model.User, error) {

	return u.repo.GetFollows(user_id)
}
