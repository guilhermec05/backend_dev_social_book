package main

import (
	"dev_book_backend/configs"
	"dev_book_backend/internal/http/router"
	database "dev_book_backend/pkg/DataBase"

	"fmt"
	"log"
	"net/http"
)

func main() {
	configs.LoadConfig()

	db, err := database.Conn()

	if err != nil {
		log.Fatal(err)

	}

	defer db.Close()

	r := router.ConfigRouter(db)
	log.Print("servindo a porta 3001")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.PORT), r))

}
