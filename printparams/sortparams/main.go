package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printStr prints each rune in the given slice
func printStr(s []rune) {
	n := 0
	for n < len(s) {
		z01.PrintRune(s[n])
		n++
	}
}

// endIndex finds the last index of the specified rune in the slice
func endIndex(s []rune, toFind rune) int {
	ret := 0
	for nextIndex := 0; nextIndex < len(s); nextIndex++ {
		if s[nextIndex] == toFind {
			ret = nextIndex + 1
		}
	}
	return ret
}

// isSorted checks if the given slice of strings is sorted
func isSorted(args []string) bool {
	for i := 1; i < len(args)-1; i++ {
		if args[i] > args[i+1] {
			return false
		}
	}
	return true
}

func main() {
	// Get the command-line arguments
	args := os.Args

	// Continue sorting until the arguments are sorted
	for !isSorted(args) {
		// Iterate through the arguments
		for i := 1; i < len(args)-1; i++ {
			// If the current argument is greater than the next one, swap them
			if args[i] > args[i+1] {
				tmp := args[i]
				args[i] = args[i+1]
				args[i+1] = tmp
			}
		}
	}

	// Iterate through the sorted arguments
	for i := 1; i < len(args); i++ {
		// Convert the argument to a rune slice
		arg := []rune(args[i])

		// Find the index of the last '/' in the argument
		lastSlashIndex := endIndex(arg, '/')

		// Print the substring starting from the last '/'
		printStr(arg[lastSlashIndex:])

		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
