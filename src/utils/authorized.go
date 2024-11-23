package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func CreatedToken(user_id uint) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["user_id"] = user_id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	return token.SignedString(secret)

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {

		return 0, errors.New("Missing authorization header")
	}

	tokenString = tokenString[len("Bearer "):]

	token, err := VerifyToken(tokenString)

	if err != nil {

		return 0, err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id, err := strconv.ParseUint(fmt.Sprintf("%0.f", permission["user_id"]), 10, 64)

		if err != nil {

			return 0, err
		}

		return user_id, nil
	}

	return 0, fmt.Errorf("invalid token")

}
