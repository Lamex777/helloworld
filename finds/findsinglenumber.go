package main

import "fmt"

func bitwiseXOR(n, res int) int {
	res ^= n
	return res
}

func findSingleNumber(numbers []int) int {
	res := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j == i {
				continue
			}
			equal := bitwiseXOR(numbers[i], numbers[j])
			if equal != 0 {
				res = numbers[i]
			} else {
				res = 0
				break
				// придумать переменную булевую которая отвечает за наличие вообще такого числа
			}
		}
		if res != 0 {
			break
		}
	}
	return res
}

func main() {
	numbers := []int{1, 2, 3, 4, 1, 5, 4, 3, 2}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}
