package main

import "fmt"

func main() {
	slice := []int{2,5,7,2,-3,3,4,5,6}
	newSlice := insertToStart (slice, 83,-95,52,06,42,73,15)
	fmt.Println(newSlice)
}

func insertToStart (xs []int, x... int) []int {
	newXS := make ([]int,0)
	for _, val := range x {
		newXS = append(newXS, val)
	}
	xs = append(newXS, xs...)
	return xs
}