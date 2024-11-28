package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *Follow) GetFollow(user_id uint64) ([]domain.User, error) {

	return u.repo.GetFollows(user_id)
}
