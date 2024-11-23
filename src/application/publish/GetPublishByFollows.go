package application

import "dev_book_backend/src/model"

func (u *Publish) GetPublishByFollows(id_publish uint64) ([]model.Publish, error) {

	return u.repo.PublishFollow(id_publish)
}
