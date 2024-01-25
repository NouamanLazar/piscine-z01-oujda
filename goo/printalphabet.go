package main

import "github.com/01-edu/z01"

func Fromatz() {
	for i := 'a'; i < 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func Fromzta() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	Fromatz()
	Fromzta()
}
