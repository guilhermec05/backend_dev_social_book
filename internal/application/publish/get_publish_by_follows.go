package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) GetPublishByFollows(id_publish uint64) ([]domain.Publish, error) {

	return u.repo.PublishFollow(id_publish)
}
