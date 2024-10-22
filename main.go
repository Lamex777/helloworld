package main

// import (
// 	"fmt"
// )

// func main() {
// 	intArr := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
// 	floatArr := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}
// 	sortedIntDesc := sortDescInt(intArr)
// 	sortedIntAsc := sortAscInt(intArr)
// 	sortedFloatDesc := sortDescFloat(floatArr)
// 	sortedFloatAsc := sortAscFloat(floatArr)
// 	fmt.Println("Sorted Int Array (Descending):", sortedIntDesc)
// 	fmt.Println("Sorted Int Array (Ascending):", sortedIntAsc)
// 	fmt.Println("Sorted Float Array (Descending):", sortedFloatDesc)
// 	fmt.Println("Sorted Float Array (Ascending):", sortedFloatAsc)
// }

// func sortDescInt(a [8]int) [8]int {
// 	for i := 0; i < (len(a) - 1); i++ {
// 		for j := (i + 1); j < len(a); j++ {
// 			if a[i] >= a[j] {
// 				continue
// 			} else {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return a
// }
// func sortAscInt(a [8]int) [8]int {
// 	for i := 0; i < (len(a) - 1); i++ {
// 		for j := (i + 1); j < len(a); j++ {
// 			if a[i] <= a[j] {
// 				continue
// 			} else {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return a
// }
// func sortDescFloat(a [8]float64) [8]float64 {
// 	for i := 0; i < (len(a) - 1); i++ {
// 		for j := (i + 1); j < len(a); j++ {
// 			if a[i] >= a[j] {
// 				continue
// 			} else {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return a
// }

// func sortAscFloat(a [8]float64) [8]float64 {
// 	for i := 0; i < (len(a) - 1); i++ {
// 		for j := (i + 1); j < len(a); j++ {
// 			if a[i] <= a[j] {
// 				continue
// 			} else {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return a
// }