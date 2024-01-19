package main

import (
	"fmt"
	"os"
)

func index(s []rune, toFind rune) int {
	for i := 0; i < len(s); i++ {
		if s[i] == toFind {
			return i
		}
	}
	return -1
}

func isIFlag(s string) bool {
	i := index([]rune(s), '=')
	return i > 0 && (s[:i] == "-i" || s[:i] == "--insert")
}

func isOFlag(s string) bool {
	return s == "-o" || s == "--order"
}

func isSorted(s []rune) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func printInOrder(s []rune) {
	for !isSorted(s) {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				tmp := s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
			}
		}
	}
	fmt.Println(string(s))
}

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("--insert")
		fmt.Println("  " + "-i")
		fmt.Println("\t " + "This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  " + "-o")
		fmt.Println("\t " + "This flag will behave like a boolean, if it is called it will order the argument.")
	} else {
		var str string
		var insertStr string
		isOrder := false
		for i := 1; i < len(os.Args); i++ {
			if isIFlag(os.Args[i]) {
				indexEq := index([]rune(os.Args[i]), '=')
				if indexEq < len([]rune(os.Args[i]))-1 {
					insertStr += os.Args[i][indexEq+1:]
				}
			} else if isOFlag(os.Args[i]) {
				isOrder = true
			} else {
				str += os.Args[i]
			}
		}
		str += insertStr
		if isOrder {
			printInOrder([]rune(str))
		} else {
			fmt.Println(str)
		}
	}
}
