package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

func CountBytes(r io.Reader) (n int64, err error) { // HL
	n, err = io.Copy(ioutil.Discard, r)
	return
}

func CountBytesInFile(wg *sync.WaitGroup, filename string) { // HL
	defer wg.Done() // HL

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := CountBytes(f) // HL
	if err != nil {
		panic(err)
	}
	fmt.Println(filename, n)
}

// START OMIT

// What will happen? // HL
func main() {
	go CountBytesInFile("async.go")
	go CountBytesInFile("censor.go")
	go CountBytesInFile("hello.go")
	go CountBytesInFile("io.go")
	for {
	}
}

// END OMIT
