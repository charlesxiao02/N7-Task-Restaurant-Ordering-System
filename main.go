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
type OrderItem struct {
	dish       Dish
	itemStatus bool
}

// struct for an Order, has list of OrderDishes and status for overall order
// maybe consider a map[name]OrderDish?
type Order struct {
	dishes      []OrderItem
	orderStatus OrderStatus
}

type OrderStatus int

const (
	StatusOrdering OrderStatus = iota
	StatusPending
	StatusComplete
)

var orderStatusName = map[OrderStatus]string{
	StatusOrdering: "ordering",
	StatusPending:  "pending",
	StatusComplete: "complete",
}

func (os OrderStatus) String() string {
	return orderStatusName[os]
}

var itemStatusName = map[bool]string{
	false: "preparing",
	true:  "ready",
}

func (o Order) GetOrderStatus() {
	fmt.Printf("Order status: %s\n", o.orderStatus)
	if o.orderStatus == StatusPending {
		for i, item := range o.dishes {
			fmt.Printf("%d: %s: Item status: %s\n", i+1, item.dish.name, itemStatusName[item.itemStatus])
		}
	}
}

// struct for the Kitchen, has a list of Orders
type Kitchen struct {
	orders []Order
}

// struct to represent a customer?
type Customer struct {
	poOrder *Order
}

// method to add a dish to a menu
func (m *Menu) AddDish(d Dish) {
	m.dishes = append(m.dishes, d)
}

// method to add a dish to an order
func (o *Order) OrderDish(d Dish) {
	oItem := OrderItem{d, false}
	o.dishes = append(o.dishes, oItem)
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

	tMenu := make(map[string]float32)

	tMenu["Mapo Tofu"] = 4.75
	tMenu["Dou Jiao Men Mian"] = 3.50

}
