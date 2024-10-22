package main

import "fmt"

func concatStrings (xs ...string ) (result string) {
	fmt.Println(xs)
	for _,s := range xs {
		result += s
	}
	return
}

func main () {
	s1 := "sdv"
	s2 := "gtrbh"
	s3 := " Пируем!!!!"

	result := concatStrings(s1,s2,s3)
	fmt.Println(result)
}