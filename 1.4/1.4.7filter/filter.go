package main

import "fmt"

func main () {
	slice := []int{3,5,6,7,9,12,43,56}
	dividerSlice := filterDividers(slice, 7)
	fmt.Println(dividerSlice)
}

func filterDividers (xs []int, div int) []int {
	newXS := make([]int, 0,len(xs))
	for _, val := range xs {
		if val % div == 0 {
			newXS = append(newXS, val)
		}
	}
	return newXS
}