package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"net/http"
	"strconv"
)

// ========================================= НЕ готово & НЕ протестировано =========================================
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := dbrepository.GetTasks()
	// ОБРАБОТАТЬ ОШИБКУ
	err = json.NewEncoder(w).Encode(tasks)
	// ОБРАБОТАТЬ ОШИБКУ
}

// ========================================= НЕ готово & НЕ протестировано =========================================
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task dbrepository.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	// ОБРАБОТАТЬ ОШИБКУ
	defer r.Body.Close()
	task, err = dbrepository.CreateTask(task)
	// ОБРАБОТАТЬ ОШИБКУ
	err = json.NewEncoder(w).Encode(task)
	// ОБРАБОТАТЬ ОШИБКУ
}

// ========================================= НЕ готово & НЕ протестировано =========================================
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task dbrepository.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	// ОБРАБОТАТЬ ОШИБКУ
	defer r.Body.Close()
	task.ID, err = strconv.Atoi(params["id"])
	// ОБРАБОТАТЬ ОШИБКУ
	task, err = dbrepository.UpdateTask(task)
	// ОБРАБОТАТЬ ОШИБКУ
	err = json.NewEncoder(w).Encode(task)
	// ОБРАБОТАТЬ ОШИБКУ
}

// ========================================= НЕ готово & НЕ протестировано =========================================
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := dbrepository.DeleteTask(strconv.Atoi(params["id"]))
	// ОБРАБОТАТЬ ОШИБКУ
}
