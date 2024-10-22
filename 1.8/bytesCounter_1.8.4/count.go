package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytes(s string) int {
	counter := len(s)
	return counter
}

func countSymbols(s string) int {
	counter := utf8.RuneCountInString(s)
	return counter
}

func main() {
	bytes := countBytes("Привет, мир!")
	fmt.Println(bytes)
	symbols := countSymbols("Привет, мир!")
	fmt.Println(symbols)
}
