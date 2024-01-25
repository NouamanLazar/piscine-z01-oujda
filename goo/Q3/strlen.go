package main

import (
	"fmt"
)

func StrLen(s string) int {
	// first we add variable
	var runed int
	// we give string a int valeu
	for range s {
		// ++ its mean the number of all the charcters in string
		runed++
	}
	return runed
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
