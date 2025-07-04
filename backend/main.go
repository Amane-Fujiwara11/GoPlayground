package main

import (
	"backend/db"
	"backend/infrastructure/mysql"
	"backend/interface/handler"
	"backend/usecase"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {
	var err error
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	fmt.Println("Connected to MySQL!")

	repo := mysql.NewTaskRepository(dbConn)
	uc := usecase.TaskUsecase{Repo: repo}
	taskHandler := handler.NewTaskHandler(&uc)
	r := NewRouter(taskHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
