package application

import (
	domain "dev_book_backend/internal/domain/user"
)

func (u *Publish) GetLikePublish(id_publish uint64) ([]domain.User, error) {

	return u.repo.GetLikePublish(id_publish)
}
