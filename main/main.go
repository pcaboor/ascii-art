package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: program <arg1> <theme>")
		return
	}

	inputText := os.Args[1]
	theme := os.Args[2]

	var filePath string

	switch theme {
	case "standard":
		filePath = "./standard.txt"
	case "shadow":
		filePath = "./shadow.txt"
	case "thinkertoy":
		filePath = "./thinkertoy.txt"
	default:
		fmt.Println("Unknown theme:", theme)
		return
	}

	inputText = strings.ReplaceAll(inputText, "\\n", "\n")

	// Lire le fichier
	ascii, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed reading file: %s", err)
		return
	}
	// Convertir les données en chaîne de caractères
	asciiFileBuffer := string(ascii)

	// Séparer le contenu du fichier par des retours à la ligne
	asciiLines := strings.Split(asciiFileBuffer, "\n")

	// Créer une slice 2D pour stocker les lignes d'art ASCII
	asciiFileLines := make([][]string, (len(asciiLines)+8)/9)

	// Remplir la slice 2D avec les lignes du fichier ASCII
	for index, line := range asciiLines {
		partieEntiere := index / 9
		modulo := index % 9
		if modulo > 0 {
			if asciiFileLines[partieEntiere] == nil {
				asciiFileLines[partieEntiere] = make([]string, 8)
			}
			asciiFileLines[partieEntiere][modulo-1] = line
		}
	}

	// Créer l'index de fichier ASCII de manière dynamique
	asciiFileIndex := make([]string, len(asciiFileLines))
	for i := range asciiFileIndex {
		asciiFileIndex[i] = string(rune(i + 32))
	}

	// Préparer les lignes pour stocker la sortie d'art ASCII
	lines := make([]string, 8)

	for i := 0; i < len(inputText); i++ {
		char := inputText[i]
		if inputText[i] == '\n' {
			printAndResetLines(lines)
			if i+1 < len(inputText) && inputText[i+1] == '\n' {
				printEmptyLines(1)
				i++
			}
		} else {
			for j := 0; j < 8; j++ {
				lines[j] += asciiFileLines[char-32][j]
			}
		}
	}
	// Imprimer les lignes restantes
	printAndResetLines(lines)
}

func printAndResetLines(lines []string) {
	for i := 0; i < 8; i++ {
		if lines[i] != "" {
			fmt.Println(lines[i])
			lines[i] = ""
		}
	}
}

func printEmptyLines(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("")
	}
}
