package application

import "dev_book_backend/src/model"

func (u *Publish) GetMyPublishById(id_publish, id_author uint64) (*model.Publish, error) {

	return u.repo.MyPublishById(id_publish, id_author)
}
