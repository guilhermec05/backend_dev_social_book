package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	Connection = ""
	Port       = 0
	Jwt_token  = ""
)

func LoadConfig() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	Jwt_token = os.Getenv("JWT_SECRET")

	Connection = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	port, err := strconv.Atoi(os.Getenv("PORT_APP"))

	if err != nil {
		log.Fatal("You are var PORT_APP")
	}

	Port = port
}
