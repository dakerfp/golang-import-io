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
	CountBytesInFile("async.go")
	CountBytesInFile("censor.go")
	CountBytesInFile("hello.go")
	CountBytesInFile("io.go")
}

// END OMIT
