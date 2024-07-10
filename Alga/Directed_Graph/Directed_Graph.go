package main

import (
	"fmt"
)

// Node представляет собой вершину графа
type Node struct {
	Value     int     // Значение вершины
	Neighbors []*Node // Список соседей вершины
	Weight    float64 // Расстояние от источника до этой вершины
	Prev      *Node   // Предыдущая вершина в кратчайшем пути
	Visited   bool    // Посещена ли эта вершина
}

// AddEdge добавляет ребро между двумя вершинами
func AddEdge(u, v *Node) {
	u.Neighbors = append(u.Neighbors, v)
}

// MinDistance находит вершину с минимальным расстоянием среди непосещенных вершин
func MinDistance(nodes []*Node) *Node {
	minDist := maxDist // Инициализируем минимальное расстояние максимальным значением
	var minNode *Node  // Инициализируем минимальную вершину нулевым указателем
	for _, node := range nodes {
		// Если вершина не посещена и ее расстояние меньше минимального
		if !node.Visited && node.Weight < minDist {
			minDist = node.Weight // Обновляем минимальное расстояние
			minNode = node        // Обновляем минимальную вершину
		}
	}
	return minNode // Возвращаем минимальную вершину
}

// Dijkstra реализует алгоритм Дейкстры для нахождения кратчайшего пути между двумя вершинами
func Dijkstra(src, dest *Node) ([]*Node, float64) {
	src.Weight = 0        // Устанавливаем расстояние от источника до себя равным нулю
	nodes := []*Node{src} // Добавляем источник в список вершин
	for len(nodes) > 0 {  // Пока список вершин не пуст
		u := MinDistance(nodes) // Находим вершину с минимальным расстоянием
		if u == nil {           // Если такой вершины нет
			break // Прерываем цикл
		}
		if u == dest { // Если мы достигли цели
			break // Прерываем цикл
		}
		u.Visited = true // Помечаем вершину как посещенную
		for i, node := range nodes {
			if node == u { // Если мы нашли вершину в списке
				nodes = append(nodes[:i], nodes[i+1:]...) // Удаляем ее из списка
				break                                     // Прерываем цикл
			}
		}
		for _, v := range u.Neighbors { // Для каждого соседа вершины u
			// Если вершина v не посещена
			if !v.Visited {
				newDist := u.Weight + 1.0 // Вычисляем новое расстояние до v
				if newDist < v.Weight {   // Если новое расстояние меньше старого
					v.Weight = newDist // Обновляем расстояние до v
					v.Prev = u         // Обновляем предыдущую вершину в пути до v
				}
				found := false // Флаг, указывающий, находится ли v в списке вершин
				for _, node := range nodes {
					if node == v { // Если мы нашли v в списке
						found = true // Устанавливаем флаг в true
						break        // Прерываем цикл
					}
				}
				if !found { // Если v не в списке
					nodes = append(nodes, v) // Добавляем v в список
				}
			}
		}
	}
	if dest.Weight != maxDist { // Если мы нашли путь до цели
		path := []*Node{} // Создаем список для хранения пути
		v := dest         // Начинаем с целевой вершины
		for v != src {    // Пока не достигнем источника
			path = append([]*Node{v}, path...) // Добавляем вершину в начало пути
			v = v.Prev                         // Переходим к предыдущей вершине
		}
		path = append([]*Node{src}, path...) // Добавляем источник в начало пути
		return path, dest.Weight             // Возвращаем путь и длину пути
	} else {
		// Иначе возвращаем nil, -1
		return nil, -1
	}
}

const maxDist = 1e9 // Константа для обозначения максимального расстояния

func main() {
	// Создаем вершины графа
	A := &Node{Value: 0, Weight: maxDist}
	B := &Node{Value: 1, Weight: maxDist}
	C := &Node{Value: 2, Weight: maxDist}
	D := &Node{Value: 3, Weight: maxDist}
	E := &Node{Value: 4, Weight: maxDist}
	F := &Node{Value: 5, Weight: maxDist}
	G := &Node{Value: 6, Weight: maxDist}
	H := &Node{Value: 7, Weight: maxDist}
	I := &Node{Value: 8, Weight: maxDist}
	// Добавляем ребра между вершинами
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
	// Находим кратчайший путь от вершины A до вершины E
	path, length := Dijkstra(A, E)
	// Выводим результат
	fmt.Println("Кратчайший путь:", path)
	fmt.Println("Длина пути:", length)
}
