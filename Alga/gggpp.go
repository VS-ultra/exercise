package main

import (
	"fmt"
)

type Node struct {
	Value     int
	Neighbors []*Node
	Weight    float64
	Prev      *Node
	Visited   bool
}

func AddEdge(u, v *Node) {
	u.Neighbors = append(u.Neighbors, v)
}

func MinDistance(nodes []*Node) *Node {
	minDist := maxDist
	var minNode *Node
	for _, node := range nodes {
		if !node.Visited && node.Weight < minDist {
			minDist = node.Weight
			minNode = node
		}
	}
	return minNode
}

func Dijkstra(src, dest *Node) ([]*Node, float64) {
	src.Weight = 0
	nodes := []*Node{src}
	for len(nodes) > 0 {
		u := MinDistance(nodes)
		if u == nil {
			break
		}
		if u == dest {
			break
		}
		u.Visited = true
		for i, node := range nodes {
			if node == u {
				nodes = append(nodes[:i], nodes[i+1:]...)
				break
			}
		}
		for _, v := range u.Neighbors {
			if !v.Visited {
				newDist := u.Weight + 1.0
				if newDist < v.Weight {
					v.Weight = newDist
					v.Prev = u
				}
				found := false
				for _, node := range nodes {
					if node == v {
						found = true
						break
					}
				}
				if !found {
					nodes = append(nodes, v)
				}
			}
		}
	}
	if dest.Weight != maxDist {
		path := []*Node{}
		v := dest
		for v != src {
			path = append([]*Node{v}, path...)
			v = v.Prev
		}
		path = append([]*Node{src}, path...)
		return path, dest.Weight
	} else {
		return nil, -1
	}
}

const maxDist = 1e9

func main() {
	A := &Node{Value: 0, Weight: maxDist}
	B := &Node{Value: 1, Weight: maxDist}
	C := &Node{Value: 2, Weight: maxDist}
	D := &Node{Value: 3, Weight: maxDist}
	E := &Node{Value: 4, Weight: maxDist}
	F := &Node{Value: 5, Weight: maxDist}
	G := &Node{Value: 6, Weight: maxDist}
	H := &Node{Value: 7, Weight: maxDist}
	I := &Node{Value: 8, Weight: maxDist}
	AddEdge(A, B)
	AddEdge(A, H)
	AddEdge(B, C)
	AddEdge(B, H)
	AddEdge(C, D)
	AddEdge(C, F)
	AddEdge(C, I)
	AddEdge(D, E)
	AddEdge(D, F)
	AddEdge(E, F)
	AddEdge(F, G)
	AddEdge(G, H)
	AddEdge(G, I)
	AddEdge(H, I)
	path, length := Dijkstra(A, E)
	fmt.Println("Кратчайший путь:", path)
	fmt.Println("Длина пути:", length)
}
