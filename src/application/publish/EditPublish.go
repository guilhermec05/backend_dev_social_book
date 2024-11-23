package application

import "dev_book_backend/src/model"

func (u *Publish) EditPublish(publish model.Publish, id_publish, id_author uint64) error {

	if err := publish.Validation(); err != nil {

		return err
	}

	return u.repo.EditPublish(publish, id_publish, id_author)
}
