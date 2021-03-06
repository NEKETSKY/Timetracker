package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"log"
	"net/http"
	"strconv"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups, err := RepSQL.GetGroups()
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to get groups", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(groups); err != nil {
		log.Println(err)
	}
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var group dbrepository.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		ReturnError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	group, err = RepSQL.CreateGroup(group)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to create group", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(group); err != nil {
		log.Println(err)
	}
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var group dbrepository.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		ReturnError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	group.ID, err = strconv.Atoi(params["id"])
	if err != nil {
		ReturnError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := RepSQL.CheckGroupByID(group); err !=nil {
		log.Println(err)
		ReturnError(w, "Such an ID does not exist", http.StatusNotFound)
		return
	}

	group, err = RepSQL.UpdateGroup(group)
	if err != nil {
		log.Println(err)
		ReturnError(w, "Failed to update group", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		log.Println(err)
	}
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ReturnError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = RepSQL.DeleteGroup(id); err != nil {
		ReturnError(w, "Failed to delete group", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
