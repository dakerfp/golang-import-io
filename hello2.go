package main

import (
	"io"
	"os"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin) // HL
	if err != nil {
		panic(err)
	}
}
