package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFirstLine(filename string) string {
	f, _ := os.Open(filename)
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	return strings.Split(string(content), "\n")[0]
}

func main() {
	first := GetFirstLine("compose.go")
	fmt.Println(first)
}
