package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	mapa := make(map[User]int)
	var result []User
	for i, user := range users {
		if mapa[user] == 0 {
			mapa[user] = i + 1
			result = append(result, user)
		}
	}
	result = result[:len(result):len(result)]
	return result
}

func main() {
	var users []User
	for i := 0; i < 100; i++ {
		i := User{gofakeit.Name(), gofakeit.Number(18, 65), gofakeit.Email()}
		users = append(users, i)
	}
	res := getUniqueUsers(users)
	for i,val := range res {
		fmt.Println(i+1,"- ", val)
	}
}
