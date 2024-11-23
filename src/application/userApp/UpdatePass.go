package application

import (
	"dev_book_backend/src/model"
	"dev_book_backend/src/utils"
	"errors"
)

func (u *User) UpdatePass(user_id uint64, passObj model.UpdatePass) error {

	user, err := u.repo.GetUsersWithPassByid(user_id)

	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(passObj.Old_pass, user.Password) {
		return errors.New("A senha não correponde")
	}

	new_pass, err := utils.HashPassword(passObj.New_pass)

	if err != nil {
		return err
	}

	_, err = u.repo.UpdateUsersPassById(user_id, new_pass)

	if err != nil {
		return err
	}

	return nil
}
