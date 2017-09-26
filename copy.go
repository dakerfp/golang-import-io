package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, os.Stdin) // HL
	})
	http.ListenAndServe(":8080", nil)
}
