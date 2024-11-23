package main

import (
	database "dev_book_backend/src/DataBase"
	routers "dev_book_backend/src/router"
	"dev_book_backend/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.Carregar()

	db, err := database.Conn()

	if err != nil {
		log.Fatal(err)

	}

	defer db.Close()

	r := routers.ConfigRouter(db)
	log.Print("servindo a porta 3001")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", utils.PORT), r))

}
