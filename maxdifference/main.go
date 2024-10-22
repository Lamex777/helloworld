package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{-2, 6, 1, 8, 3, 946679, 10, 245}
	diff := maxDifference(numbers)
	fmt.Println(diff)
}

func maxDifference(xs []int) int {
	var min float64 = math.Inf(4)
	var max float64 = math.Inf(-4)
	var diff int
	for i := 0; i < len(xs); i++ {
		if min > float64(xs[i]) {
			min = float64(xs[i])
		}
		if max < float64(xs[i]) {
			max = float64(xs[i])
		}
	}
	fmt.Println(max, min)
	diff = int(max) - int(min)
	return diff
}
