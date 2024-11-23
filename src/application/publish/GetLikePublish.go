package application

import "dev_book_backend/src/model"

func (u *Publish) GetLikePublish(id_publish uint64) ([]model.User, error) {

	return u.repo.GetLikePublish(id_publish)
}
