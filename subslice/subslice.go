package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subslice := getSubSlice(numbers, 1, 7)
	fmt.Println(subslice)
}

func getSubSlice (xs []int, start, end int) []int {
	subslice := xs[start:end]
	return subslice
}
