package main

import "fmt"

func Swap(a *int, b *int) {
	lazar := *a
	*a = *b
	*b = lazar
}
func SwapG(slice []int) []int {
    // Create a new slice with the same length as the input slice
    result := []int{}

    // Set all elements in the new slice to 0
    for range slice {
        result = append(result, 0)
    }

    return result
}

func main() {
	a := 0
	b := 1
    Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
    fmt.Println(SwapG([]int{1, 2, 3, 4, 5, 6, 7}))
    fmt.Println(SwapG([]int{}))
}
