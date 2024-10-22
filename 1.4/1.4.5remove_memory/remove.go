package main

import "fmt"

func main() {
	slice := make([]int, 20,25)
	fmt.Println(cap(slice))
	slice = removeExtraMemory(slice)
	fmt.Println(slice, cap(slice))
}

func removeExtraMemory (xs []int) []int {
	xs = xs[:len(xs):len(xs)]
	return xs
}