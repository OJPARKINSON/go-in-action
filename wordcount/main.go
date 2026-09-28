package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		panic("Please provide a filename")
	}

	fileContents, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("failed to read the file")
	}

	words := strings.Fields(string(fileContents))

	fmt.Println("Found", len(words), "words")
}
