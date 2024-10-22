package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	mapa := make(map[string]int)
	slice := strings.Split(text, " ")
	result := ""
	for _, val := range slice {
		mapa[val]++
	}

	for key, val := range mapa {
		if val != 1 {
			delete(mapa, key)
		}
	}

	for i := 0; i < len(slice); i++ {
		for key := range mapa {
			if slice[i] == key {
				result += key + " "
			}
		}
	}
	return result
}

func main() {
	text := "hello world my name is Yanis and i am a doctor i like my job"
	uniqueText := createUniqueText(text)
	fmt.Println(uniqueText)
}
