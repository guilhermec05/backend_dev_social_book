package application

import "dev_book_backend/src/model"

func (u *Publish) PublicPublish(publish model.Publish) error {

	if err := publish.Validation(); err != nil {

		return err
	}

	return u.repo.PublicPublish(publish)
}
