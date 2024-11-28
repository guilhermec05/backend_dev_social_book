package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

func (u *Publish) PublicPublish(publish domain.Publish) error {

	if err := publish.Validation(); err != nil {

		return err
	}

	return u.repo.PublicPublish(publish)
}
