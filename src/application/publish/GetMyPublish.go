package application

import "dev_book_backend/src/model"

func (u *Publish) GetMyPublish(id uint64) ([]model.Publish, error) {

	return u.repo.MyPublish(id)
}
