package main

import (
	"fmt"
	"strings"
	// "unicode"
)

type Word struct {
	Word string
	Pos  int
}

// 	func filterWords(text string, censorMap map[string]string) string {
// 		sentences := splitSentences(text)
// 		if len(sentences) > 1 {

// 			return strings.Join(sentences, " ")
// 		}
// 		strings.Fields(text)
// 	}

// 	func WordsToSentence(words []string) string {
// 		filtered := make([]string, 0, len(words))
// 		for _, word := range words {
// 			if word != "" {
// 				filtered = append(filtered, word)
// 			}
// 		}
// 		return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
// 	}

// 	func CheckUpper(old, new string) string {
// 		if len(old) == 0 || len(new) == 0 {
// 			return new
// 	}
// 	chars := []rune(old)
// 	if unicode.IsUpper(chars[0]) {
// 		runes := []rune(new)
// 		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
// 	}
// }

func splitSentences(message string) []string {
	originSentences := strings.Split(message, "!")
	// var orphan string
	var sentences []string
	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")
		for i:=0; i<len(words); i++ {
			if words[i] == " " {
				words = append(words[:i], words[(i+1)])
			}
		}
		originSentences[i] = strings.Join(words, " ")
		sentences = append(sentences, originSentences[i])
	}
	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независиым с помощью крипты! Крипта будущее финансового мира!"
	// censorMap := map[string]string{
	// 	"крипта":   "фрукты",
	// 	"крипту":   "фрукты",
	// 	"крипты":   "фруктов",
	// 	"биткоин":  "яблоки",
	// 	"лайткоин": "яблоки",
	// 	"эфир":     "яблоки",
	// }
	// filteredText := filterWords(text, censorMap)
	// fmt.Println(filteredText)
	res := splitSentences(text)
	for _, v := range res {
	fmt.Println(v)
	}
}
