package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) Insert(value int) {
	if bt.root == nil {
		bt.root = &Node{
			value: value,
		}
		return
	}
	cur := bt.root
	for {
		if value > cur.value {
			if cur.right == nil {
				cur.right = &Node{
					value: value,
				}
				return
			}
			cur = cur.right
		} else {
			if cur.left == nil {
				cur.left = &Node{
					value: value,
				}
				return
			}
			cur = cur.left
		}
	}
}

func (bt *BinaryTree) Serach(value int) (*Node, bool) {
	if bt.root == nil {
		return nil, false
	}
	cur := bt.root

	for cur != nil {
		if cur.value == value {
			return cur, true
		}
		if value > cur.value {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return nil, false
}

// 递归的方式先序遍历二叉树
func (bt *BinaryTree) PreorderTraversal(cur *Node) {
	if cur.left != nil {
		bt.PreorderTraversal(cur.left)
	}
	fmt.Print(cur.value, " ")
	if cur.right != nil {
		bt.PreorderTraversal(cur.right)
	}
}

func main() {
	bt := new(BinaryTree)
	for i := 0; i < 10; i++ {
		bt.Insert(rand.Int())
	}
	bt.PreorderTraversal(bt.root)
}
