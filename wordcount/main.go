package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		panic("Please provide a filename")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("failed to read the file")
	}

	scanner := bufio.NewScanner(file)
	var wordCount int

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Found", wordCount, "words")
}
