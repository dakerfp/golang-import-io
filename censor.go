package main

import (
	"bytes"
	"io"
	"os"
)

// START OMIT

type Censor struct {
	R         io.Reader
	Blacklist []string
	Replace   string
}

func (c *Censor) Read(p []byte) (n int, err error) {
	n, err = c.R.Read(p)
	if err != nil && err != io.EOF {
		return
	}
	replaced := p
	for _, word := range c.Blacklist {
		replaced = bytes.Replace(replaced, []byte(word), []byte(c.Replace), -1)
	}

	if len(p) >= len(replaced) {
		n = copy(p, replaced)
	} else {
		err = io.ErrShortBuffer
	}

	return
}

// END OMIT
// START MAIN

func main() {
	censor := &Censor{
		os.Stdin,
		[]string{"fuck", "shit"},
		"@#!",
	}
	_, _ = io.Copy(os.Stdout, censor)
}

// END MAIN
