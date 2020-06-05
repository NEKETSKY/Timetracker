package handlers

import (
	"encoding/json"
	//	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"net/http"
)


func GetGroups(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	groups := dbrepository.GetGroups()
	json.NewEncoder(w).Encode(groups)
}

/*
func CreateGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")


	json.NewEncoder(w).Encode(???)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//дальше брать через params["id"]
	_ = json.NewDecoder(r.Body).Decode(params)

	json.NewEncoder(w).Encode(???)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//дальше брать через params["id"]
}

*/
