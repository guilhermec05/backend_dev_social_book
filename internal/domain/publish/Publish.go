package domain

import (
	"errors"
	"strings"
	"time"
)

type Publish struct {
	ID          uint64    `json:"id,omitempty"`
	Tile        string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	Author_id   uint64    `json:"author_id,omitempty"`
	Author_nick string    `json:"author_nick,omitempty"`
	Like        uint64    `json:"like,omitempty"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

func (u *Publish) Prepare() error {
	if err := u.Validation(); err != nil {
		return err
	}

	return u.format()

}

func (p *Publish) Validation() error {
	var Err error = nil
	if p.Tile == "" {
		Err = errors.New("o titulo não preechido")
	}

	if p.Content == "" {
		Err = errors.New("conteudo não preechido")
	}

	return Err
}

func (u *Publish) format() error {

	u.Tile = strings.TrimSpace(u.Tile)
	u.Content = strings.TrimSpace(u.Content)

	return nil
}
