package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/neketsky/Timetracker/dbrepository"
	"log"
	"net/http"
)

var RepSQL dbrepository.TaskRepositorySQL

func ConnectWithHandlers(dbase *sql.DB) {
	RepSQL.DB = dbase
}

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
