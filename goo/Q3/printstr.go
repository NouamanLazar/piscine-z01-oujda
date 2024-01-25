package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, Print := range s {
		z01.PrintRune(Print)
	}
	z01.PrintRune('\n')
}

func main() {
	PrintStr("Hello World!")
}
