package main

import (
	"fmt"
)

var node = &List{
	val: 1,
}

type List struct {
	node *List
	val  int
}

// 浅拷贝，原对象的值实际上没有发生变化
func (o *List) Updata() {
	o = node
	fmt.Println(&o, &node)
}

func main() {
	l := new(List)
	fmt.Println(&l)
	l.Updata()
	fmt.Println(l)
	fmt.Println(&l)
}
