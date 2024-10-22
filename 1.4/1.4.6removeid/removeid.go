package main

import "fmt"

func main(){
	slice := []int{4,6,3,2,6}
	slice = removeIDX(slice, 3)
	fmt.Println(slice)
}

func removeIDX (xs[]int, id int) []int {
	xs = append(xs[:id], xs[(id+1):]...)
	return xs
}