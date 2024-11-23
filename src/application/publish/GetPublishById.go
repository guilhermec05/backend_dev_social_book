package application

import "dev_book_backend/src/model"

func (u *Publish) GetPublishById(id_publish uint64) (*model.Publish, error) {

	return u.repo.PublishById(id_publish)
}
