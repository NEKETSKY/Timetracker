package main

import (
	"github.com/gorilla/mux"
	"github.com/neketsky/Timetracker/dbrepository"
	"github.com/neketsky/Timetracker/handlers"
	"log"
	"net/http"
)

var taskRepositorySQL = dbrepository.TaskRepositorySQL{}

func main() {
	var err error
	taskRepositorySQL.DB, err = dbrepository.DBInit()
	if err != nil {
		log.Panic(err)
	}

	handlers.ConnectWithHandlers(taskRepositorySQL.DB)
	defer taskRepositorySQL.DB.Close()

	router := mux.NewRouter()
	GroupsRouter := router.PathPrefix("/groups").Subrouter()
	TasksRouter := router.PathPrefix("/tasks").Subrouter()
	TimeframesRouter := router.PathPrefix("/timeframes").Subrouter()

	GroupsRouter.HandleFunc("/", handlers.GetGroups).Methods(http.MethodGet)

	GroupsRouter.HandleFunc("/", handlers.CreateGroup).Methods(http.MethodPost)
	GroupsRouter.HandleFunc("/{id}", handlers.UpdateGroup).Methods(http.MethodPut)
	GroupsRouter.HandleFunc("/{id}", handlers.DeleteGroup).Methods(http.MethodDelete)

	TasksRouter.HandleFunc("/", handlers.GetTasks).Methods(http.MethodGet)
	TasksRouter.HandleFunc("/", handlers.CreateTask).Methods(http.MethodPost)
	TasksRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods(http.MethodPut)
	TasksRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods(http.MethodDelete)

	TimeframesRouter.HandleFunc("/", handlers.CreateTimeframe).Methods(http.MethodPost)
	TimeframesRouter.HandleFunc("/{id}", handlers.DeleteTimeframe).Methods(http.MethodDelete)

	log.Println("Server starting on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
