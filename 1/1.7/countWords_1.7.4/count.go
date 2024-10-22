package main

import (
	"fmt"
	"strings"
)

func countWordOccurrences(text string) map[string]int {
	mapa := make(map[string]int)
	slice := strings.Split(text, " ")
	for _, val := range slice {
		mapa[val]++
	}
	return mapa
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)
	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
