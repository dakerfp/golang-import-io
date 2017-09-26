package main

import (
	"fmt"
	"io"
	"os"
)

type BytesCounter struct {
	N int
}

func (bc *BytesCounter) Write(p []byte) (n int, err error) {
	bc.N += len(p)
	return
}

func main() {
	var bc BytesCounter
	_, _ = io.Copy(&bc, os.Stdin)
	fmt.Println("total = ", bc.N)
}
