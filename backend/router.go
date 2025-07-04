package main

import (
	"backend/interface/handler"

	"github.com/gorilla/mux"
)

func NewRouter(taskHandler *handler.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	return r
}
