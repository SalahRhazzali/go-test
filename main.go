package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo = []Todo{
	{ID: "1", Title: "Buy milk", Completed: false},
}

func main() {
	// Define routes
	http.HandleFunc("/todos", getTodos)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
