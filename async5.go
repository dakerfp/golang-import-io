package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CountBytes(r io.Reader) (n int64, err error) { // HL
	n, err = io.Copy(ioutil.Discard, r)
	return
}

// START OMIT

type count struct {
	filename string
	size     int64
}

func CountBytesInFile(filename string, counts chan count) { // HL
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := CountBytes(f) // HL
	if err != nil {
		panic(err)
	}

	counts <- count{filename, n}
}

// END OMIT
// START MAIN

func main() {
	counts := make(chan count) // HL
	go CountBytesInFile("async.go", counts)
	go CountBytesInFile("censor.go", counts)
	go CountBytesInFile("hello.go", counts)
	go CountBytesInFile("io.go", counts)
	for count := range counts { // XXX // HL
		fmt.Println(count.filename, count.size)
	}
}

// END MAIN
