package repository

import (
	"database/sql"
	"dev_book_backend/src/model"
	"errors"
	"time"
)

type Publish struct {
	sql *sql.DB
}

func NewRepositoryPublish(sql *sql.DB) *Follow {
	return &Follow{sql: sql}
}

func (u *Follow) PublicPublish(publish model.Publish) error {

	stmt, err := u.sql.Prepare("insert into publish(title, content, author_id ) value (?,?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(publish.Tile, publish.Content, publish.Author_id)

	if err != nil {
		return err
	}

	return nil

}

func (u *Follow) MyPublish(id uint64) ([]model.Publish, error) {
	db, err := u.sql.Query(`SELECT p.*, 
						u.nick,
						CASE when  count(lipSum.publish_id) > 0 then count(lipSum.publish_id) else 0 end as quantity_like
						from publish p 
						join users u on p.author_id  = u.id 
						left join (
							SELECT lp.publish_id   from like_publish lp
						) lipSum on p.id  = lipSum.publish_id
						where p.author_id = ?
						GROUP  by p.id ,p.title ,p.content ,p.author_id,p.created_at ,nick
						`, id)

	if err != nil {
		return nil, err
	}

	var publishies []model.Publish

	for db.Next() {
		var publish model.Publish
		var createdAtRaw []uint8

		if err := db.Scan(&publish.ID, &publish.Tile, &publish.Content, &publish.Author_id, &createdAtRaw, &publish.Author_nick, &publish.Like); err != nil {
			return nil, err
		}

		publish.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

		if err != nil {
			return nil, err
		}

		publishies = append(publishies, publish)
	}

	return publishies, nil
}

func (u *Follow) MyPublishById(id_publish, id_author uint64) (*model.Publish, error) {
	db := u.sql.QueryRow(`SELECT 
							p.*, 
							u.nick,
							COALESCE(count(lp.publish_id), 0) AS quantity_like
						FROM 
							publish p
						JOIN 
							users u ON p.author_id = u.id
						LEFT JOIN 
							like_publish lp ON p.id = lp.publish_id
						WHERE 
							p.author_id = ?
							AND p.id = ?
						GROUP BY 
							p.id, u.nick;

						`, id_author, id_publish)

	publish := &model.Publish{}
	var createdAtRaw []uint8

	publish.Like = 0

	err := db.Scan(&publish.ID, &publish.Tile, &publish.Content, &publish.Author_id, &createdAtRaw, &publish.Author_nick, &publish.Like)

	if err != nil {
		return nil, err
	}

	publish.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

	if err != nil {
		return nil, err
	}

	return publish, nil
}

func (u *Follow) EditPublish(publish model.Publish, id_publish, id_author uint64) error {

	stmt, err := u.sql.Prepare("UPDATE  publish set title = ? , content = ?  where author_id = ? and id =  ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(publish.Tile, publish.Content, id_author, id_publish)

	if err != nil {
		return err
	}

	nums_row_affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if nums_row_affected == 0 {
		return errors.New("não foi alterado nenhum dado")
	}
	return nil

}

func (u *Follow) DeletePublish(id_publish, id_author uint64) error {

	stmt, err := u.sql.Prepare("Delete from publish   where author_id = ? and id =  ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id_author, id_publish)

	if err != nil {
		return err
	}

	nums_row_affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if nums_row_affected == 0 {
		return errors.New("não foi possivel Deletar")
	}
	return nil

}

func (u *Follow) PublishById(id_publish uint64) (*model.Publish, error) {
	db := u.sql.QueryRow(`SELECT 
							p.*, 
							u.nick,
							COALESCE(count(lp.publish_id), 0) AS quantity_like
						FROM 
							publish p
						JOIN 
							users u ON p.author_id = u.id
						LEFT JOIN 
							like_publish lp ON p.id = lp.publish_id
						WHERE 
						 p.id = ?
						GROUP BY 
							p.id, u.nick;

						`, id_publish)

	publish := &model.Publish{}
	var createdAtRaw []uint8

	publish.Like = 0

	err := db.Scan(&publish.ID, &publish.Tile, &publish.Content, &publish.Author_id, &createdAtRaw, &publish.Author_nick, &publish.Like)

	if err != nil {
		return nil, err
	}

	publish.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

	if err != nil {
		return nil, err
	}

	return publish, nil
}

func (u *Follow) PublishFollow(id uint64) ([]model.Publish, error) {
	db, err := u.sql.Query(`SELECT 
								p.*, 
								u.nick,
								COALESCE(count(lp.publish_id), 0) AS quantity_like
							FROM 
								publish p
							JOIN 
								users u ON p.author_id = u.id
							JOIN  
								follow f on f.follow_id = u.id 
							LEFT JOIN 
								like_publish lp ON p.id = lp.publish_id
							WHERE 
								f.user_id = ?
							GROUP BY 
								p.id, u.nick;
						`, id)

	if err != nil {
		return nil, err
	}

	var publishies []model.Publish

	for db.Next() {
		publish := model.Publish{}
		var createdAtRaw []uint8

		if err := db.Scan(&publish.ID, &publish.Tile, &publish.Content, &publish.Author_id, &createdAtRaw, &publish.Author_nick, &publish.Like); err != nil {
			return nil, err
		}

		publish.Created_at, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))

		if err != nil {
			return nil, err
		}

		publishies = append(publishies, publish)
	}

	return publishies, nil
}

func (u *Follow) LikePublish(id_publish, id_author uint64) error {

	stmt, err := u.sql.Prepare("insert into like_publish(publish_id, user_id ) value (?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id_publish, id_author)

	if err != nil {
		return err
	}

	nums_row_affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if nums_row_affected == 0 {
		return errors.New("não foi alterado nenhum dado")
	}
	return nil

}

func (u *Follow) GetLikePublish(id_publish uint64) ([]model.User, error) {
	db, err := u.sql.Query(`SELECT u.id,u.name, u.nick,u.email  from users u 
							join like_publish lp on lp.user_id = u.id 
							where lp.publish_id = ?`, id_publish)

	if err != nil {
		return nil, err
	}

	var Users []model.User

	for db.Next() {
		user := model.User{}

		if err := db.Scan(&user.ID, &user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		Users = append(Users, user)
	}

	return Users, nil
}

func (u *Follow) UnLikePublish(id_publish, id_author uint64) error {

	stmt, err := u.sql.Prepare("delete from like_publish where publish_id = ? and user_id = ? ")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id_publish, id_author)

	if err != nil {
		return err
	}

	nums_row_affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if nums_row_affected == 0 {
		return errors.New("não foi deletado nenhum dado")
	}
	return nil
}
