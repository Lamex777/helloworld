package main

import (
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

type Animal struct {
	Type   string
	Animal string
	Age    int
}

func main() {
	animals := getAnimals()
	result := preparePrint(animals)
	fmt.Println(result)
	fmt.Println(animals)
}

func getAnimals() []Animal {
	var a []Animal
	for i := 'a'; i < 'd'; i++ {
		var i Animal
		i.Type = gofakeit.Name()
		i.Animal = gofakeit.FirstName()
		i.Age = gofakeit.Number(0, 20)
		a = append(a, i)
	}
	return a
}

func preparePrint(a []Animal) string {
	var s string
	for i, val := range a {
		s += "Животное " + strconv.Itoa(i+1) + ": " + val.Type + "\nИмя: " + val.Animal + "\nВозраст: " + strconv.Itoa(val.Age) + "\n"
	}
	return s
}
