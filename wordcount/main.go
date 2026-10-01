package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide at least one filename")
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("failed to read the file %s", fileName)
			file.Close()
			continue
		}

		var wordCount int
		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wordCount++
		}

		if scanner.Err() != nil {
			log.Printf("scan error for %s, %v", fileName, scanner.Err())
			file.Close()
			continue
		}

		file.Close()
		fmt.Printf("%s: %d words\n", fileName, wordCount)
	}
}
