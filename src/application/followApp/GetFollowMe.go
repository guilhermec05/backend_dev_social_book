package application

import (
	"dev_book_backend/src/model"
)

func (u *Follow) GetFollowMe(user_id uint64) ([]model.User, error) {

	return u.repo.GetFollowMe(user_id)
}
