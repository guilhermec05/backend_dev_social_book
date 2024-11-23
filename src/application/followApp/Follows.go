package application

import (
	interfaces "dev_book_backend/src/interfaces/followImp"
)

type Follow struct {
	repo interfaces.Follow
}

func NewAppFollow(repo interfaces.Follow) *Follow {
	return &Follow{repo}
}
