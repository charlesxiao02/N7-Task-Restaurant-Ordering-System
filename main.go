package main

import (
	"fmt"
)

// basic struct to represent a Dish, with Name and Price
type Dish struct {
	name  string
	price float32
}

// struct for a Menu, list of dishes
type Menu struct {
	dishes []Dish
}

// intermediary struct for Order, contains a Dish and relevant info
type OrderDishStruct struct {
	dish     Dish
	quantity int
	status   bool
}

// struct for an Order, has list of OrderDishes and status for overall order
// maybe consider a map[name]OrderDish?
type Order struct {
	dishes  []OrderDishStruct
	dishMap map[string]OrderDishStruct
	status  bool
}

// struct for the Kitchen, has a list of Orders
type Kitchen struct {
	orders []Order
}

// method to add a dish to a menu
func (m *Menu) AddDish(d Dish) {
	m.dishes = append(m.dishes, d)
}

// method to add a dish to an order
func (o *Order) OrderDish(d Dish) bool {

	return false
}

// method to submit order to kitchen
func (k *Kitchen) SubmitOrder(o Order) {
	k.orders = append(k.orders, o)
}

func main() {
	fmt.Printf("test print")

	menu := Menu{}

	mapoTofu := Dish{"Mapo Tofu", 4.75}
	douJiaoMenMian := Dish{"Dou Jiao Men Mian", 3.50}

	menu.AddDish(mapoTofu)
	menu.AddDish(douJiaoMenMian)

}
