package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *Follow) GetFollowMe(user_id uint64) ([]domain.User, error) {

	return u.repo.GetFollowMe(user_id)
}
