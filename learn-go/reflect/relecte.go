package main

import (
	"fmt"
	"reflect"
)

var _ Order = new(order)

type Order interface {
	do()
}

type order struct {
	ordID      int
	customerID int
}

func (o *order) do() {}

func createQuery(q Order) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Element ", t.Elem())
	fmt.Println("Kind ", t.Kind(), t.Elem().Kind())
}

func main() {
	o := &order{
		ordID:      456,
		customerID: 56,
	}
	createQuery(o)
}
