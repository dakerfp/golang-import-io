package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// START OMIT

func CountBytesInFile(wg *sync.WaitGroup, filename string) { // HL
	defer wg.Done() // HL

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := io.Copy(ioutil.Discard, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(filename, n)
}

func main() {
	var wg sync.WaitGroup // HL
	wg.Add(4)             // HL
	go CountBytesInFile(&wg, "async.go")
	go CountBytesInFile(&wg, "censor.go")
	go CountBytesInFile(&wg, "hello.go")
	go CountBytesInFile(&wg, "io.go")
	wg.Wait() // HL
}

// END OMIT
