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

func FindMinAndParent(n, p *Node) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}

	for n.Left != nil {
		p = n
		n = n.Left
	}

	return n, p
}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		}

		if n.Left == nil {
			return n.Right
		}

		if n.Right == nil {
			return n.Left
		}

		min, parent := FindMinAndParent(n.Right, n)
		min.Left = n.Left
		if parent != n {
			parent.Left = min.Right
			min.Right = n.Right
		}
		n = min
	}

	return n
}

func (n *Node) Find(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		return n.Left.Find(value)
	} else if value > n.Value {
		return n.Right.Find(value)
	} else {
		return n
	}
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
	root.Delete(10)
	root.Delete(10)
	root.Delete(3)
	node8 := root.Find(8)
	if node8 != nil {
		fmt.Println("Узел со значением 8 найден")
	}
	InOrderTraversal(root)
}
