package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Fatal(err)
	}
}

func ResposeError(w http.ResponseWriter, code int, err error) {
	ResponseJson(w, code, struct {
		Error string `json:erro`
	}{Error: err.Error()})
}
