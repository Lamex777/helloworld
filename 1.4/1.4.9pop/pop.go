package main

import "fmt"

func main() {
	slice := []int{2,54,2,3,6,8,6,9,70}
	fmt.Println(Pop(slice))
}

func Pop (xs []int) (int, []int) {
	num := xs[0]
	xs = append(xs[:0], xs[1:]...)
	return num,xs
}