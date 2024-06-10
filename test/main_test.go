// main_test.go
package main

import (
	"testing"
)

// TestAdd vérifie que la fonction Add fonctionne correctement.
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{-1, -1, -2},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func Add(a int, b int) int {
	return a + b
}

// Subtract soustrait le deuxième entier du premier et retourne le résultat.
func Subtract(a int, b int) int {
	return a - b
}
