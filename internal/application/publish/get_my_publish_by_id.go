package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) GetMyPublishById(id_publish, id_author uint64) (*domain.Publish, error) {

	return u.repo.MyPublishById(id_publish, id_author)
}
