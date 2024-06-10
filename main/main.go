package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// File to read
	filePath := "./standard.txt"

	inputText := os.Args[1]

	// Read the file
	ascii, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed reading file: %s", err)
		return
	}
	// Convert data to string
	asciiFileBuffer := string(ascii)

	asciiFilePaint := strings.Split(asciiFileBuffer, "\n\n")
	var asciiFileLines [][]string
	for _, block := range asciiFilePaint {
		asciiFileLines = append(asciiFileLines, strings.Split(block, "\n"))
	}

	// Create the ASCII file index dynamically
	asciiFileIndex := make([]string, len(asciiFileLines))
	for i := range asciiFileIndex {
		asciiFileIndex[i] = string(rune(i + 32))
	}

	// Prepare output lines
	lines := make([]string, 8)

	// Process input text
	for _, char := range inputText {

		for i := 0; i < 8; i++ {
			if char == '\n' {
				lines[i] += asciiFileLines[char-32][i]
			} else {
				lines[i] += asciiFileLines[char-32][i]
			}

		}
	}

	// Print the resulting ASCII art
	for i := 0; i < 8; i++ {
		fmt.Println(lines[i])
	}
}
