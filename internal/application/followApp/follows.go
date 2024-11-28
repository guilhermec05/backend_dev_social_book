package application

import (
	domain "dev_book_backend/internal/domain/user"
)

type Follow struct {
	repo domain.FollowRepository
}

func NewAppFollow(repo domain.FollowRepository) *Follow {
	return &Follow{repo}
}
