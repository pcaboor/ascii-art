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
	}
	// Convert data to string
	asciiFileBuffer := string(ascii)

	// Split result with double \n separator
	asciiFileDisplay := strings.Split(asciiFileBuffer, "\n\n")

	// After read the file, create a slice of string
	asciiIndex := []string{filePath}

	// Travel in each characters
	for _, char := range inputText {
		i := -1
		for y, value := range asciiIndex {
			if string(char) == value {
				i = y
				break
			}
		}
		// Print ascii char
		if i != -1 {
			fmt.Println(asciiFileDisplay[char-32])
		} else {
			fmt.Println(asciiFileDisplay[char-32])
		}
	}
}
