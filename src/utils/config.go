package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Connection = ""
	PORT       = 0
)

func Carregar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	Connection = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	PORT, err = strconv.Atoi(os.Getenv("PORT_APP"))

	if err != nil {
		log.Fatal("You are var PORT_APP")
	}
}
