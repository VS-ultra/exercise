package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

func (n *Node) Insert(value int) *Node {
	if n == nil {
		return NewNode(value)
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}

	return n
}

func InOrderTraversal(n *Node) {
	if n != nil {
		InOrderTraversal(n.Left)
		fmt.Print(n.Value, " ")
		InOrderTraversal(n.Right)
	}
}

func main() {
	root := NewNode(10)

	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(3)
	root.Insert(12)
	root.Insert(5)
	root.Insert(10)
	root.Insert(10)
	root.Insert(10)
	root.Insert(10)

	InOrderTraversal(root)

}
