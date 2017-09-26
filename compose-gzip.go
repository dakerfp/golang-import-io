package main

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Priority int
	Text     string
}

var todoList []TodoItem

// START OMIT

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Encoding", "gzip") // HL
		w.Header().Set("Content-Type", "application/json")
		gw := gzip.NewWriter(w)    // HL
		enc := json.NewEncoder(gw) // HL
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
