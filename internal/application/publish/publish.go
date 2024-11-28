package application

import (
	domain "dev_book_backend/internal/domain/publish"
)

type Publish struct {
	repo domain.PublishRpository
}

func NewAppPublish(repo domain.PublishRpository) *Publish {
	return &Publish{repo}
}
