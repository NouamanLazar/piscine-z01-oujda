package main

import "fmt"

func BasicAtoi(s string) int {
	// kan3tiw int value
	var i int

	for _, atoi := range s {
		// kanbdlo 9imat i b i*10 li heya 1 * 10 we atoi-'1' heya 10 n7ayd mnha 0
		i = i*10 + int(atoi-'0')
	}
	return i
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
