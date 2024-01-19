package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Get the command-line arguments (including the program name)
	result := os.Args

	// Iterate through the characters of the program name
	for j, i := range result[0] {
		// Skip the first two characters (program name and space)
		if j > 1 {
			// Print each character after the first two
			z01.PrintRune(i)
		}
	}

	// Print a newline character at the end
	z01.PrintRune(10)
}
