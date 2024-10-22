package main

import "fmt"

func countVowels (s string) int {
	counter := 0
	for _, symbol := range s {
		switch symbol {
		case 'a','o','i','u','y','e','A','E','Y','U','I','O','У','Е','Ё','Э','О','А','Ы','Я','И','Ю','у','е','ё','э','о','а','ы','я','и','ю':
			counter ++
		}
	}
	return counter
}

func main() {
	s := "Приветики-пистолетики"
	dig := countVowels(s)
	fmt.Println(dig)
}