package main

import "fmt"

var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{}

var Concat func(xs ...interface{}) interface{}

var Sum func(xs ...interface{}) interface{}

func operation(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i)
}

func intSum(slice ...interface{}) interface{} {
	var intSum int

	for _, numbers := range slice {
		for _, n := range numbers {
			switch t := n.(type) {
			case int:
				intSum += t
			default:
				continue
			}
		}
	}
	return intSum
}

// func floatSum(slice ...interface{}) interface{} {

// }

// func Concatenation (stroka ...interface{}) interface{} {
// }

func main() {
	Operate = operation
	summa := Operate(intSum, 7, 3, 7, 4, 212)
	fmt.Println(summa)
}
