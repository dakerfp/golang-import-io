package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// START OMIT

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

// END OMIT

func main() {
	var wg sync.WaitGroup // HL
	wg.Add(4)             // HL
	go CountBytesInFile(&wg, "async.go")
	go CountBytesInFile(&wg, "censor.go")
	go CountBytesInFile(&wg, "hello.go")
	go CountBytesInFile(&wg, "io.go")
	wg.Wait() // HL
}
