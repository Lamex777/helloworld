package main

import "fmt"

func main() {
	slice := []int{2,56,34,8,3,1,0,7}
	newSlice := insertAfterIDX(slice,2,6,4,9,5)
	fmt.Println(newSlice)

}

func insertAfterIDX (xs []int, id int, x... int) []int {
	// xs1 := xs[:id]
	xs2 := make([]int, (len(xs)-(id+1)))
	copy(xs2, xs[(id+1):])
	xs = append(xs[:(id+1)], x...)
	xs = append(xs, xs2...)
	return xs
}