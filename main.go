package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func extractWords(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	words := [][]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, strings.Fields(strings.ToLower(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning: %v", err)
	}
	return words, nil
}

func hash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum[0:4])
}

func main() {
	filename := os.Args[1]
	words, err := extractWords(filename)
	if err != nil {
		fmt.Printf("error reading file %q: %v\n", filename, err)
		os.Exit(1)
	}

	for _, line := range words {
		for _, word := range line {
			fmt.Printf("%s ", hash(word))
		}
		fmt.Println()
	}
}
