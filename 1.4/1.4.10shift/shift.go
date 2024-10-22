package main

import (
	"errors"
	"fmt"
)

func main() {
	slice := make([]int, 0)
	firstElement, shiftedSlice, err := Shift(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(firstElement, shiftedSlice, slice)
	}
}

func Shift(xs []int) (int, []int, error) {
	if len(xs) == 0 {
		return 0, nil, errors.New("Slice is empty")
	} else {
		firstElement := xs[0]
		newXS := make([]int, len(xs), len(xs))
		for i := 1; i < len(xs); i++ {
			if i == (len(xs) - 1) {
				newXS[i] = xs[i-1]
				newXS[0] = xs[(len(xs) - 1)]
				break
			}
			newXS[i] = xs[i-1]
		}
		return firstElement, newXS, nil
	}
}
