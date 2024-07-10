package main

import "fmt"

type Node struct {
	value     int     // Значение вершины
	neighbors []*Node // Список соседей вершины
}

type Queue struct {
	items []int // Массив для хранения элементов очереди
}

func AddEdge(u, v *Node) {
	u.neighbors = append(u.neighbors, v) // Добавляем вершину v в список соседей вершины u
	v.neighbors = append(v.neighbors, u) // Добавляем вершину u в список соседей вершины v
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item) // Добавляем элемент в конец очереди
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 { // Если очередь пуста
		return -1 // Возвращаем -1
	}
	removed := q.items[0] // Сохраняем первый элемент очереди
	q.items = q.items[1:] // Удаляем первый элемент из очереди

	return removed // Возвращаем удаленный элемент
}

// Функция для поиска в ширину в графе g с начальной вершиной s
func bfs(g []*Node, s int) {
	// Создаем очередь для хранения вершин, которые нужно посетить
	q := &Queue{}
	// Создаем массив для хранения посещенных вершин
	visited := make([]bool, len(g))
	// Добавляем начальную вершину в очередь и отмечаем ее как посещенную
	q.Enqueue(s)
	visited[s] = true
	// Пока очередь не пуста, извлекаем из нее вершину v
	for len(q.items) > 0 {
		v := q.Dequeue()
		// Если v равно -1, значит очередь пуста, и мы пропускаем его
		if v == -1 {
			continue
		}
		// Выводим значение вершины
		fmt.Println(g[v].value)
		// Для каждой смежной с v вершины u
		for _, u := range g[v].neighbors {
			// Если u еще не посещена
			if !visited[u.value] {
				// Добавляем u в очередь и отмечаем ее как посещенную
				q.Enqueue(u.value)
				visited[u.value] = true
			}
		}
	}
}

func main() {
	// Создаем вершины графа
	a := &Node{value: 0}
	b := &Node{value: 1}
	c := &Node{value: 2}
	d := &Node{value: 3}

	// Добавляем ребра в граф
	AddEdge(a, b)
	AddEdge(a, c)
	AddEdge(b, d)
	AddEdge(c, d)

	// Запускаем поиск в ширину от вершины 0
	bfs([]*Node{a, b, c, d}, 0)
	// Результат:
	// 0
	// 1
	// 2
	// 3
}
