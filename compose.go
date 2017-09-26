package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START OMIT

type TodoItem struct {
	Priority int
	Text     string
}

var todoList []TodoItem

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w) // HL
		if err := enc.Encode(todoList); err != nil {
			log.Println(err)
		}
	}
}

// END OMIT

func main() {
	http.HandleFunc("/todo", handleTodo)
	http.ListenAndServe(":8080", nil)
}
