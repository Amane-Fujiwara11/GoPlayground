package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var db *sql.DB

func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, completed FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = len(tasks) + 1
	task.Completed = false
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateTaskCompletion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedCompletion struct {
		Completed bool `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updatedCompletion); err != nil {
		fmt.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	fmt.Println("Updated completion received:", updatedCompletion.Completed)
	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Completed = updatedCompletion.Completed
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedStatus struct {
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&updatedStatus)

	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Completed = updatedStatus.Status == "完了"
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {

	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/taskdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL!")

	r := mux.NewRouter()
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}/complete", updateTaskCompletion).Methods("PATCH")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r))
}
