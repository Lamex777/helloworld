package main

import "fmt"

func main() {
	s := []int{3,6,3}
	// p := &s
	appendInt(s, 64, 54)
	fmt.Println(s)
}

func appendInt(xs []int, x ...int) []int{
	xs = append(xs, x...)
	return xs
}
