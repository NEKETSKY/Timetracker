package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReturnError(w http.ResponseWriter, data interface{}, code int) {
	message, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(message); err != nil {
		log.Fatal(err)
	}
}
