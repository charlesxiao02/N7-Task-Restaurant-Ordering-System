package main

import (
	"fmt"
)

type Dish struct {
	name  string
	price float32
}

type Menu struct {
	dishes []Dish
}

type Order struct {
	dishes []Dish
}

type Kitchen struct {
	orders []Order
}

func main() {
	fmt.Printf("test print")
}
