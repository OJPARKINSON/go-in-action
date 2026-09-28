package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide a filename")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("failed to read the file")
	}

	var wordCount int
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Found", wordCount, "words")
}
