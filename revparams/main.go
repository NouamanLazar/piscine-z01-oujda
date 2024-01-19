package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s []rune) {
	n := 0
	for n < len(s) {
		z01.PrintRune(s[n])
		n++
	}
}

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
	for i := len(os.Args) - 1; i > 0; i-- {
		arg := []rune(os.Args[i])
		printStr(arg[endIndex(arg, '/'):])
		z01.PrintRune('\n')
	}
}
