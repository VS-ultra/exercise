package main

import "fmt"

type hashElement struct {
	key   string
	value int
	next  *hashElement
}

type HashTable struct {
	size     int
	elements []*hashElement
}

func NewHashTable(size int) *HashTable {
	return &HashTable{size: size, elements: make([]*hashElement, size)}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}

func (ht *HashTable) Insert(key string, value int) {
	index := ht.hash(key)
	newElement := &hashElement{key: key, value: value}

	if ht.elements[index] == nil {
		ht.elements[index] = newElement
	} else {
		current := ht.elements[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newElement
	}
}
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hash(key)
	current := ht.elements[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func main() {
	table := NewHashTable(10)

	table.Insert("Alice", 30)
	table.Insert("Sam", 35)
	// table.Insert("Bob", 25)

	age, found := table.Get("Bob")
	if found {
		fmt.Println("Возраст Bob:", age)
	} else {
		fmt.Println("Информация о Bob отсутствует")
	}
}
