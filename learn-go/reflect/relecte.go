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
	// kind 获取 type 的类型
	fmt.Println("Kind ", t.Kind(), t.Elem().Kind())
	// name 获取 type 的名称
	fmt.Println("Name ", t.Name(), "ss", t.Elem().Name())
}

/*
Type  *main.order
Value  &{456 56}
Element  main.order
Kind  ptr struct
Name   ss order
*/

func main() {
	o := &order{
		ordID:      456,
		customerID: 56,
	}
	createQuery(o)
}
