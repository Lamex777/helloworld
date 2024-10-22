package main

// import "fmt"

// var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{}

// var Concat func(xs ...interface{}) interface{}

// var Sum func(xs ...interface{}) interface{}

// func operation(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
// 	return f(i...)
// }

// func intSum(slice ...interface{}) (interface{}, interface{}) {
// 	var intSum int
// 	var floatSum float64
// 	for _, numbers := range slice {
// 		switch t := numbers.(type) {
// 		case int:
// 			intSum += t
// 		case float64: 
// 			floatSum += t
// 		default:
// 			continue
// 		}
// 	}
// 	return intSum, floatSum
// }

// // func floatSum(slice ...interface{}) interface{} {

// // }

// // func Concatenation (stroka ...interface{}) interface{} {
// // }

// func main() {
// 	Operate = operation
// 	intSumma, floatSumma := Operate(intSum, 7, 3, 7, 4, 212)
// 	fmt.Println()
// }
