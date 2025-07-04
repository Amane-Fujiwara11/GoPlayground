package main

import (
	"backend/constants"
	"backend/db"
	"backend/models"
	"backend/response"
	"backend/validation"
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

var dbConn *sql.DB

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, err error, message string) {
	response.RespondError(w, status, err, message)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks(dbConn)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err, "DB取得エラー")
		return
	}
	respondJSON(w, http.StatusOK, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := validation.ParseAndValidateJSON(r, &task, func(v interface{}) error {
		return v.(*models.Task).Validate()
	})
	if err != nil {
		respondError(w, http.StatusBadRequest, err, "バリデーションエラー")
		return
	}
	if err := models.CreateTask(dbConn, &task); err != nil {
		respondError(w, http.StatusInternalServerError, err, "DB登録エラー")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := models.DeleteTask(dbConn, id); err != nil {
		respondError(w, http.StatusInternalServerError, err, "DB削除エラー")
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
	if err := json.NewDecoder(r.Body).Decode(&updatedStatus); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if updatedStatus.Status != constants.StatusRegistered && updatedStatus.Status != constants.StatusDoing && updatedStatus.Status != constants.StatusCompleted {
		http.Error(w, "Invalid status value", http.StatusBadRequest)
		return
	}

	if err := models.UpdateTaskStatus(dbConn, id, updatedStatus.Status); err != nil {
		respondError(w, http.StatusInternalServerError, err, "DBステータス更新エラー")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	var err error
	dbConn, err = db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	fmt.Println("Connected to MySQL!")

	r := NewRouter()
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
