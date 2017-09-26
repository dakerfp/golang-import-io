package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// START OMIT

func CountBytesInFile(filename string) {
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
	go CountBytesInFile("async.go")
	go CountBytesInFile("censor.go")
	go CountBytesInFile("hello.go")
	go CountBytesInFile("io.go")
	// XXX: will exit main // HL
}

// END OMIT
