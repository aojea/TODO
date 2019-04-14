package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aojea/TODO/internal/handlers"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the Database
	DB, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	// Initialize the API server
	r := mux.NewRouter()
	// API version 1
	api := r.PathPrefix("/api/v1").Subrouter()
	// TODO lists handlers
	api.HandleFunc("/user/me/lists", handlers.ListsHandler)
	api.HandleFunc("/user/me/lists/{listId}", handlers.ListIDHandler)
	// Tasks handlers
	api.HandleFunc("/lists/{listId}/tasks", handlers.TasksHandler)
	api.HandleFunc("/lists/{listId}/tasks/{taskId}", handlers.TaskIDHandler)
	// Run the app
	http.ListenAndServe(":8080", r)
}
