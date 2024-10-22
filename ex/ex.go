// package main

// import "fmt"

// func main() {
// 	slice1 := make([]int, 0, 3) // указываем тип, длину 0 и емкость 3

// 	// fmt.Println("адрес пустого первого среза - пустой")
// 	slice1 = append(slice1, 1, 2, 3)
// 	// a := &slice1[0]
// 	// fmt.Println("Исходный срез 1:", slice1, "adress 1 -", a) // Исходный срез #1: [1 2 3]
// 	// slice2 := append(slice1, 4, 5, 6)
// 	// b := &slice1[0]
// 	// c := &slice2[0]
// 	// fmt.Println("Cрез 1 после append для среза 2:", slice1, "adress 1 -", b) // Исходный срез #1: [1 2 3]
// 	// fmt.Println("Новый срез 2:", slice2, "adress 2 -", c)                    // Новый срез #2: [1 2 3 4 5 6]
// 	appendIntWrong(slice1, 7, 8, 9)
// 	// d := &slice1[0]
// 	// e := &slice2[0]
// 	// fmt.Println("Срез 1 после функции:", slice1, "adress 1 -", d) // Исходный срез: [1 2 3]
// 	// fmt.Println("Срез 2 после функции:", slice2, "adress 2 -", e) // Новый срез: [1 2 3 4 5 6]
// 	slice3 := make([]int, 0, 6)
// 	slice4 := append(slice3, slice1...)

// 	g := &slice4[0]
// 	fmt.Println("Исходный срез 4:", slice4, "adress 4 -", g) // Новый срез #4: [1 2 3]
// 	appendIntWrong(slice3, 7, 8, 9)
// 	h := &slice4[0]
// 	fmt.Println("Исходный срез 3:", slice3, "adress 3 - пустой") // Исходный срез #3: []
// 	fmt.Println("Новый срез 4:", slice4, "adress 4 -", h)        // Новый срез #4: [7 8 9] 
// }

// func appendIntWrong(xs []int, x ...int) {

// 	fmt.Println("xs before -", xs, "адреса нет? тк срез пустой")
// 	xs = append(xs, x...)
// 	fmt.Println("xs after -", xs, "adress xs aft -", &xs[0])
// }



package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 3) // указываем тип, длину 0 и емкость 3
	slice1 = append(slice1, 1, 2, 3)
	appendIntWrong(slice1, 7, 8, 9)
	slice3 := make([]int, 0, 2)
	slice4 := append(slice3, slice1...)
	// fmt.Println(&slice3[0])
	g := &slice4[0]
	fmt.Println("Исходный срез 4:", slice4, "adress 4 -", g) // Новый срез #4: [1 2 3]
	appendIntWrong(slice3, 7, 8, 9)
	h := &slice4[0]
	// fmt.Println("Исходный срез 3:", slice3, "adress 3 - ", &slice3[0]) // Исходный срез #3: []
	fmt.Println("Новый срез 4:", slice4, "adress 4 -", h)        // Новый срез #4: [7 8 9] 
}

func appendIntWrong(xs []int, x ...int) {

	fmt.Println("xs before -", xs, "адрес xs -")
	xs = append(xs, x...)
	fmt.Println("xs after -", xs, "adress xs aft -", &xs[0])
}

func appendIntRight(xs []int, x ...int) []int {
	return append(xs, x...)
	}
