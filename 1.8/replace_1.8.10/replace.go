package main

import "fmt"

func replaceSymbols(s string, old, new string) string {
	inter := []rune(s)
	old2 := []rune(old)
	new2 := []rune(new)

	for i, symbol := range inter {
		switch symbol {
		case old2[0]:
			inter[i] = new2[0]
		}
	}
	res := string(inter)
	return res
}

func main() {
	s := "Напиши функцию ReplaceSymbols, которая заменяет все вхождения символа old в строке str на символ new и возвращает полученную строку"
	fmt.Println(replaceSymbols(s, "и", "е"))
}
