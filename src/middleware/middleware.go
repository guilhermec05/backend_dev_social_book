package middleware

import (
	"dev_book_backend/src/utils"
	"errors"
	"fmt"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Autentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)

			utils.ResposeError(w, http.StatusUnauthorized, errors.New("Missing authorization header"))

			return
		}

		tokenString = tokenString[len("Bearer "):]

		_, err := utils.VerifyToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResposeError(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
