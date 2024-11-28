package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) EditPublish(publish domain.Publish, id_publish, id_author uint64) error {

	if err := publish.Validation(); err != nil {

		return err
	}

	return u.repo.EditPublish(publish, id_publish, id_author)
}
