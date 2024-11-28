package domain

import domain "dev_book_backend/internal/domain/user"

type PublishRpository interface {
	PublicPublish(publish Publish) error
	EditPublish(publish Publish, id_publish, id_author uint64) error
	DeletePublish(id_publish, id_author uint64) error
	MyPublish(id uint64) ([]Publish, error)
	PublishFollow(id uint64) ([]Publish, error)
	MyPublishById(id_publish, id_author uint64) (*Publish, error)
	PublishById(id_publish uint64) (*Publish, error)
	LikePublish(id_publish, id_author uint64) error
	GetLikePublish(id_publish uint64) ([]domain.User, error)
	UnLikePublish(id_publish, id_author uint64) error
}
