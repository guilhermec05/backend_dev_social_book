package domain

import (
	"dev_book_backend/pkg/utils"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type UpdatePass struct {
	Old_pass string `json:"old_pass,omitempty"`
	New_pass string `json:"new_pass,omitempty"`
}

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"pass,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(when string) error {
	if err := u.Validation(when); err != nil {
		return err
	}

	return u.format(when)

}

func (u *User) Validation(when string) error {
	var Err error = nil
	if u.Name == "" {
		Err = errors.New("name não existe")
	}

	if u.Nick == "" {
		Err = errors.New("nick não existe")
	}

	if u.Email == "" {
		Err = errors.New("email não existe")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		Err = errors.New("email invalido")
	}

	if u.Password == "" && when == "inserted" {
		Err = errors.New("pass não existe")
	}

	return Err
}

func (u *User) format(when string) error {
	u.Email = strings.TrimSpace(u.Email)

	if when == "inserted" {

		hash_pass, err := utils.HashPassword(strings.TrimSpace(u.Password))
		if err != nil {
			return err
		}
		u.Password = hash_pass
	}
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	return nil
}

func (u *User) ValidateLogin() error {

	if u.Email == "" {
		return errors.New("email não existe")
	}

	if u.Password == "" {
		return errors.New("pass não existe")
	}

	if err := u.format("inserted"); err != nil {
		return err
	}

	return nil

}
