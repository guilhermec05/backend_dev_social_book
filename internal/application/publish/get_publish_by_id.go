package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) GetPublishById(id_publish uint64) (*domain.Publish, error) {

	return u.repo.PublishById(id_publish)
}
