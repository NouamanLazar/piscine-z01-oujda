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

func main() {
	// Iterate through command-line arguments starting from the second one
	for i := 1; i < len(os.Args); i++ {
		// Convert the argument to a rune slice
		arg := []rune(os.Args[i])

		// Find the index of the last '/' in the argument
		lastSlashIndex := endIndex(arg, '/')

		// Print the substring starting from the last '/'
		printStr(arg[lastSlashIndex:])

		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
