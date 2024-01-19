package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	isNeg := false
	result := 0
	index := 0
	len := len([]rune(s))

	if index < len && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			isNeg = true
		}
		index++
	}
	for index < len && s[index] == '0' {
		index++
	}
	for index < len && s[index] >= '0' && s[index] <= '9' {
		result *= 10
		result += int(s[index] - '0')
		index++
	}
	if isNeg {
		result *= -1
	}
	if index < len {
		return 0
	}
	return result
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--upper" {
		for i := 2; i < len(os.Args); i++ {
			let := atoi(os.Args[i])
			if let >= 1 && let <= 26 {
				z01.PrintRune(rune(let - 1 + 'A'))
			} else {
				z01.PrintRune(' ')
			}
		}
	} else {
		for i := 1; i < len(os.Args); i++ {
			let := atoi(os.Args[i])
			if let >= 1 && let <= 26 {
				z01.PrintRune(rune(let - 1 + 'a'))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	if len(os.Args) > 1 {
		z01.PrintRune('\n')
	}
}
