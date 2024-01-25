package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {
	lazar := *a / *b
	lazarr := *a % *b
	*a = lazar
	*b = lazarr
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
