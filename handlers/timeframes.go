package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"log"
	"net/http"
	"strconv"
)

func CreateTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var timeframe dbrepository.Timeframe
	err := json.NewDecoder(r.Body).Decode(&timeframe)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	timeframe, err = RepSQL.CreateTimeframe(timeframe)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to create timeframe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(timeframe); err != nil {
		log.Println(err)
	}
}

func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		ReturnError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err = RepSQL.DeleteTimeframe(id); err != nil {
		log.Println(err)
		ReturnError(w, "Failed to delete timeframe", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
