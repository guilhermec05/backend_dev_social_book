package database

import (
	"database/sql"
	"dev_book_backend/configs"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", configs.Connection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
