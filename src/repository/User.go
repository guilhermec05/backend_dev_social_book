package repository

import (
	"database/sql"
	"dev_book_backend/src/model"
	"errors"
	"time"
)

type User struct {
	sql *sql.DB
}

func NewRepositoryUser(sql *sql.DB) *User {
	return &User{sql: sql}
}

func (u *User) CreateNewUser(user model.User) (uint64, error) {

	stmt, err := u.sql.Prepare("insert into users(name, nick, email, password ) value (?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return uint64(id), nil
}

func (u *User) GetAllUsers() ([]model.User, error) {
	db, err := u.sql.Query(" SELECT id,name,nick,email,created_at from users as s")

	if err != nil {
		return nil, err
	}

	var Users []model.User

	for db.Next() {
		var user model.User
		var createdAtRaw []uint8

		if err := db.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &createdAtRaw); err != nil {
			return nil, err
		}

		user.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

		if err != nil {
			return nil, err
		}

		Users = append(Users, user)
	}

	return Users, nil
}

func (u *User) GetUsersById(id uint64) (*model.User, error) {
	var err error

	db := u.sql.QueryRow(" SELECT id,name,nick,email,created_at from users  where id = ?", id)

	var user model.User
	var createdAtRaw []uint8

	if err := db.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &createdAtRaw); err != nil {

		err = errors.New("Não foi encontrado nenhum usuário")
		return nil, err
	}

	user.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) UpdateUsersById(id uint64, user *model.User) (int, error) {
	var err error

	stmt, err := u.sql.Prepare("update users set name = ?, nick = ?, email = ?  where  id = ?")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, id)

	if err != nil {
		return 0, err
	}

	last_id, _ := result.LastInsertId()

	return int(last_id), nil
}

func (u *User) DeleteUser(id uint64) (int, error) {
	var err error

	stmt, err := u.sql.Prepare("delete from users  where  id = ?")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return 0, err
	}

	affected_row, _ := result.RowsAffected()

	return int(affected_row), nil
}

func (u *User) GetUsersByEmail(email string) (*model.User, error) {

	db := u.sql.QueryRow(" SELECT id,email,password from users  where email = ?", email)

	var user model.User

	if err := db.Scan(&user.ID, &user.Email, &user.Password); err != nil {

		err = errors.New("Não foi encontrado nenhum usuário")
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUsersWithPassByid(id uint64) (*model.User, error) {
	var err error

	db := u.sql.QueryRow(" SELECT id , password from users  where id = ?", id)

	var user model.User

	if err := db.Scan(&user.ID, &user.Password); err != nil {

		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) UpdateUsersPassById(id uint64, new_pass string) (int, error) {

	var err error

	stmt, err := u.sql.Prepare("update users set password = ? where  id = ?")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(new_pass, id)

	if err != nil {
		return 0, err
	}

	last_id, _ := result.LastInsertId()

	return int(last_id), nil

}
