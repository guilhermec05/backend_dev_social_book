package interfaces

import "dev_book_backend/src/model"

type Publish interface {
	PublicPublish(publish model.Publish) error
	EditPublish(publish model.Publish, id_publish, id_author uint64) error
	DeletePublish(id_publish, id_author uint64) error
	MyPublish(id uint64) ([]model.Publish, error)
	PublishFollow(id uint64) ([]model.Publish, error)
	MyPublishById(id_publish, id_author uint64) (*model.Publish, error)
	PublishById(id_publish uint64) (*model.Publish, error)
	LikePublish(id_publish, id_author uint64) error
	GetLikePublish(id_publish uint64) ([]model.User, error)
	UnLikePublish(id_publish, id_author uint64) error
}
