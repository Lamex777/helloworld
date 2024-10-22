package main

import (
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}

func getUsers() []User {
	slice := make([]User, 0)
	for i := 'a'; i < 'k'; i++ {
		var i User
		i.Name = gofakeit.Name()
		i.Age = gofakeit.Number(18,65)
		slice = append(slice, i)
	}
	return slice
}

func preparePrint(u []User) string {
	var s string
	for i, val := range u {
		s += "Имя " + strconv.Itoa(i+1) + ": " + val.Name + "\nВозраст: " + strconv.Itoa(val.Age) + "\n"
	}
	return s
}
