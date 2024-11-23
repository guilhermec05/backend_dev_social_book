package application

import interfaces "dev_book_backend/src/interfaces/publish"

type Publish struct {
	repo interfaces.Publish
}

func NewAppPublish(repo interfaces.Publish) *Publish {
	return &Publish{repo}
}
