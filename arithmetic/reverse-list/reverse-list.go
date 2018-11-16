package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (o *ListNode) Insert(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	if o == nil {
		o = node
		return o
	}
	node.Next = o
	o = node
	return o
}

func (o *ListNode) Display() {
	for o != nil {
		fmt.Print(o.Val, " ")
		o = o.Next
	}
	fmt.Println()
}

// 递归的方法反转链表
func Reverse(o *ListNode) *ListNode {
	if o == nil || o.Next == nil {
		return o
	}
	list := Reverse(o.Next)
	o.Next.Next = o
	o.Next = nil
	return list
}

// 非递归的方式反转链表
func Reverse2(o *ListNode) *ListNode {
	if o == nil || o.Next == nil {
		return o
	}
	prev := o
	// 从链表的第二个开始反转
	cur := o.Next
	for cur != nil {
		// 暂存下一个节点的位置
		temp := cur.Next
		// 当前节点指向上一个节点的位置
		cur.Next = prev
		// 指针向后移动一位
		prev = cur
		// 处理下一个节点
		cur = temp
	}
	o.Next = nil
	return prev
}

func main() {
	var list *ListNode
	for i := 0; i < 10; i++ {
		list = list.Insert(i)
	}

	list.Display()
	list = Reverse(list)
	list.Display()
	list = Reverse2(list)
	list.Display()
}
