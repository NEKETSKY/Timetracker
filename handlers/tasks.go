package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"log"
	"net/http"
	"strconv"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := dbrepository.GetTasks()
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to get group", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println(err)
	}

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task dbrepository.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
		ReturnError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	task, err = dbrepository.CreateTask(task)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(task); err != nil {
		log.Println(err)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task dbrepository.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Invalid request", http.StatusBadRequest)
	}
	defer r.Body.Close()
	task.ID, err = strconv.Atoi(params["id"])
	if err != nil {
		ReturnError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	task, err = dbrepository.UpdateTask(task)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(task); err != nil {
		log.Println(err)
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ReturnError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err = dbrepository.DeleteTask(id); err !=nil {
		log.Println(err)
		ReturnError(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
