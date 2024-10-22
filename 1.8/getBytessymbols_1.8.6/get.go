package main

import (
	"fmt"
)

func getBytes(s string) []byte {
	bytes := []byte(s)
	return bytes
}

func getSymbols(s string) []rune {
	symbols := []rune(s)
	return symbols
}

func main() {
	text := "привет дженкинс!"
	bytes := getBytes(text)
	symbols := getSymbols(text)
	fmt.Println(bytes)
	fmt.Println(symbols)
}
