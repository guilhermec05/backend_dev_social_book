package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) GetMyPublish(id uint64) ([]domain.Publish, error) {

	return u.repo.MyPublish(id)
}