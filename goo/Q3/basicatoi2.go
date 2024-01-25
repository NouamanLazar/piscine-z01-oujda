package main

import "fmt"

func BasicAtoi2(s string) int {
	// kan3tiw int value
	var i int

	for _, atoi := range s {
		// ila kan string fih chi charachters mn ghir l2arm return 0
		if !(atoi >= '0' && atoi <= '9') {
			return 0
		}
		i = i*10 + int(atoi-'0')
	}
	return i
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
