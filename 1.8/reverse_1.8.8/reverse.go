package main

import "fmt"

func reverseString(s string) string {
	if s == "" {
		return "В функцию передана пустая строка"
	}
	runes := []rune(s)
	runes = reverseRunes(runes)

	result := runesToString(runes)
	return result
}

func reverseRunes(runes []rune) []rune {
	lenght := len(runes) / 2

	for i := 0; i < lenght; i++ {
		runes[i], runes[(len(runes)-1-i)] = runes[(len(runes)-1-i)], runes[i]
	}
	return runes
}

func runesToString(runes []rune) string {
	var s string
	for _, runa := range runes {
		s += string(runa)
	}
	return s
}

func main() {
	s := "Armen nemnojko molodec"
	fmt.Println(reverseString(s))
	// reverseString(s)
}
