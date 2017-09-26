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

// START OMIT
func CountBytesInFiles(counts chan count, filenames ...string) {
	var wg sync.WaitGroup
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go func(fn string) {
			CountBytesInFile(filename, count)
			wg.Done()
		}(fn)
	}
	wg.Wait()
	close(counts) // HL
}

func main() {
	counts := make(chan count)
	go CountBytesInFiles(counts, "async.go", "censor.go", "hello.go", "io.go") // HL
	for count := range counts {                                                // HL
		fmt.Println(count.filename, count.size)
	}
}

// END OMIT
