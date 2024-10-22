package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func main() {
	order := Order{}
	dish1 := Dish{"Pizza", 10.99}
	dish2 := Dish{"Burger", 5.99}

	order.addDish(dish2)
	order.addDish(dish1)
	fmt.Println(order.Dishes)

	order.removeDish(dish1)
	fmt.Println(order.Dishes)

	order.calculatTotal()
	fmt.Println(order)
}

func (o *Order) addDish(d Dish) {
	o.Dishes = append(o.Dishes, d)
}

func (o *Order) removeDish(d Dish) {

	for i, val := range o.Dishes {
		if val == d {
			if i != (len(o.Dishes) - 1) { // не последний элемент
				o.Dishes = append(o.Dishes[:i], o.Dishes[i+1:]...)
				break
			} else {
				o.Dishes = (o.Dishes)[:i]
				break
			}
		}
	}
}

func (o *Order) calculatTotal () {
	o.Total = 0
	for _, val := range o.Dishes {
		o.Total += val.Price
	}
}
