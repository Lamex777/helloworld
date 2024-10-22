package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	result := make(map[string]int)
	lowertext := strings.ToLower(txt)
	slice := strings.Split(lowertext, ".")
	clearText := strings.Join(slice, " ")
	slice = strings.Split(clearText, ",")
	moreClearText := strings.Join(slice, " ")
	newSlice := strings.Fields(moreClearText)
	for _, search := range words {
		for _, word := range newSlice {
			fmt.Println("search - ", search, "word - ", word)
			if word == search {
				result[search]++
			}
		}
	}
	return result
}
func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}
	result := CountWordsInText(txt, words)
	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
}
