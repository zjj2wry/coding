package main

import "fmt"

type Node struct {
	value int
	Next  *Node
}

type Linkedlist struct {
	head *Node
}

// 往链表的头部插入数据
func (o *Linkedlist) InsertFirst(value int) {
	node := &Node{
		value: value,
	}
	node.Next = o.head
	o.head = node
	return
}

// 删除指定值的某个节点
func (o *Linkedlist) Delete(value int) *Node {
	if o.head == nil {
		return nil
	}
	cur := o.head
	pre := o.head
	for cur.value != value {
		if cur.Next == nil {
			return nil
		}
		pre = cur
		cur = cur.Next
	}
	if cur == o.head {
		o.head = o.head.Next
	} else {
		pre.Next = cur.Next
	}
	return cur
}

// 删除倒数第 n 个节点
func (o *Linkedlist) DeleteNthFromEnd(n int) *Node {
	if o.head == nil {
		return nil
	}
	first := o.head
	second := o.head
	for i := 0; i <= n; i++ {
		first = first.Next
		// 链表的长度小于 n
		if first == nil {
			return nil
		}
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	deletedNode := second.Next
	second.Next = second.Next.Next
	return deletedNode
}

func (o *Linkedlist) Display() {
	cur := o.head
	for cur != nil {
		fmt.Print(cur.value, " ")
		cur = cur.Next
	}
	fmt.Println()
}

// 递归的方法反转链表
func Reverse(o *Node) *Node {
	if o == nil || o.Next == nil {
		return o
	}
	list := Reverse(o.Next)
	o.Next.Next = o
	o.Next = nil
	return list
}

// 非递归的方式反转链表
func Reverse2(o *Node) *Node {
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
	// 插入节点
	link := new(Linkedlist)
	for i := 0; i < 10; i++ {
		link.InsertFirst(i)
	}
	link.Display()
	// 删除节点
	node := link.Delete(0)
	fmt.Println("deleted node: ", node.value)
	link.Display()
	// 删除倒数第 n 个节点
	node = link.DeleteNthFromEnd(2)
	fmt.Println("deleted node: ", node.value)
	link.Display()
	// 反转链表
	node = Reverse(link.head)
	link2 := new(Linkedlist)
	link2.head = node
	link2.Display()
}
