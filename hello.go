package main

import (
	"io"
	"os"
)

// START OMIT

func main() {
	buffer := make([]byte, 64) // HL
	for {
		n, err := os.Stdin.Read(buffer) // HL
		switch err {
		case nil:
			// proceed
		case io.EOF: // HL
			return
		default:
			panic(err)
		}

		n, err = os.Stdout.Write(buffer[:n]) // HL
		if err != nil {
			panic(err)
		}
	}
}

// END OMIT
