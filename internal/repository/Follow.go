package repository

import (
	"database/sql"
	domain "dev_book_backend/internal/domain/user"
)

type Follow struct {
	sql *sql.DB
}

func NewRepositoryFollow(sql *sql.DB) *Follow {
	return &Follow{sql: sql}
}

func (u *Follow) FollowUser(user_id uint64, follow_id uint64) (int, error) {

	stmt, err := u.sql.Prepare("insert into follow(user_id, follow_id) value (?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user_id, follow_id)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return int(id), nil

}

func (u *Follow) UnFollowUser(user_id uint64, follow_id uint64) (int, error) {

	stmt, err := u.sql.Prepare("delete from follow where user_id = ? and follow_id = ?")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user_id, follow_id)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return int(id), nil

}

func (u *Follow) GetFollowMe(user_id uint64) ([]domain.User, error) {

	db, err := u.sql.Query(`SELECT u.name, u.nick  from follow f
 								join users u on f.user_id = u.id 
								where f.follow_id = ?`, user_id)

	if err != nil {
		return nil, err
	}
	var Users []domain.User

	for db.Next() {
		var user domain.User

		if err := db.Scan(&user.Name, &user.Nick); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		Users = append(Users, user)
	}

	return Users, nil

}

func (u *Follow) GetFollows(user_id uint64) ([]domain.User, error) {

	db, err := u.sql.Query(`SELECT u.name, u.nick  from follow f
 							join users u on f.follow_id  = u.id 
 							where f.user_id =  ?`, user_id)

	if err != nil {
		return nil, err
	}
	var Users []domain.User

	for db.Next() {
		var user domain.User

		if err := db.Scan(&user.Name, &user.Nick); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		Users = append(Users, user)
	}

	return Users, nil

}
