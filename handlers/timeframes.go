package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"net/http"
	"strconv"
)

// ========================================= НЕ готово & НЕ протестировано =========================================
func CreateTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var timeframe dbrepository.Timeframe
	err := json.NewDecoder(r.Body).Decode(&timeframe)
	// ОБРАБОТАТЬ ОШИБКУ
	defer r.Body.Close()
	timeframe, err = dbrepository.CreateTimeframe(timeframe)
	// ОБРАБОТАТЬ ОШИБКУ
	err = json.NewEncoder(w).Encode(timeframe)
	// ОБРАБОТАТЬ ОШИБКУ
}

// ========================================= НЕ готово & НЕ протестировано =========================================
func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := dbrepository.DeleteTimeframe(strconv.Atoi(params["id"]))
	// ОБРАБОТАТЬ ОШИБКУ
}
