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
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

var db *sql.DB

func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, status FROM tasks")
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
	task.Status = "未完了"

	result, err := db.Exec("INSERT INTO tasks (title, content, status) VALUES (?, ?, ?)", task.Title, task.Content, task.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedStatus struct {
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&updatedStatus)

	_, err := db.Exec("UPDATE tasks SET status = ? WHERE id = ?", updatedStatus.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
	r.HandleFunc("/tasks/{id}/status", updateTaskStatus).Methods("PATCH")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r))
}
