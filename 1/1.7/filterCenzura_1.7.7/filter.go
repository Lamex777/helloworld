package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	text := strings.Split(sentence, " ")
	for i := 0; i < len(text); i++ {
		if filter[text[i]] {
			if i == (len(text) - 1) {
				text = text[:i]
				break
			}
			text = append(text[:i], text[i+1:]...)
			i--
		}
	}
	result := sliceToString(text)
	return result
}

func sliceToString (text []string) (result string) {
	for i:=0; i<len(text); i++ {
		if i == len(text) -1 {
			result += text[i]
			break
		}
		result += text[i] + " "
	}
	return
}
func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}
	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence, len(filteredSentence))
}
