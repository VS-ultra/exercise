package main

import "fmt"

// node - это тип, который представляет узел в связном списке.
type node struct {
	key   string
	value string
	next  *node
}

// hashmap - это тип, который реализует хеш-таблицу, используя строки в качестве ключей и значений.
type hashmap struct {
	// buckets - это срез указателей на узлы, которые хранят пары ключ-значение.
	buckets []*node
}

// hashstr - это функция, которая возвращает хеш-значение для данной строки.
func hashstr(val string) int { // Изменил тип возвращаемого значения на int
	var strUint uint64 = 1
	hashFactor := uint64(11)
	r := []rune(val)
	for _, runeValue := range r {
		strUint *= hashFactor
		strUint += uint64(runeValue)
	}
	hashStr := strUint % 1000
	return int(hashStr) // Привел uint64 к int
}

// Set - это метод, который вставляет или обновляет пару ключ-значение в хеш-таблице.
func (h *hashmap) Set(key, val string) {
	hash_value := hashstr(key)
	bucket := h.buckets[hash_value] // Использовал одно имя для переменной
	for n := bucket; n != nil; n = n.next {
		if n.key == key {
			n.value = val
			return
		}
	}
	newNode := &node{key, val, bucket} // Использовал тип node, а не hashmap
	h.buckets[hash_value] = newNode    // Обновил указатель на корзину
}

// Get - это метод, который возвращает значение, связанное с ключом в хеш-таблице.
// Он также возвращает булево значение, указывающее, существует ли ключ или нет.
func (h *hashmap) Get(key string) (value string, ok bool) {
	hash_value := hashstr(key)
	bucket := h.buckets[hash_value] // Использовал одно имя для переменной
	for n := bucket; n != nil; n = n.next {
		if n.key == key {
			return n.value, true
		}
	}
	return "", false
}

// Delete - это метод, который удаляет пару ключ-значение из хеш-таблицы.
func (h *hashmap) Delete(key string) {
	hash_value := hashstr(key)
	bucket := &h.buckets[hash_value] // Использовал одно имя для переменной и сделал его указателем на указатель
	if *bucket == nil {
		return
	}
	if (*bucket).key == key {
		*bucket = (*bucket).next
		return
	}
	prev := *bucket
	for n := prev.next; n != nil; n = n.next {
		if n.key == key {
			prev.next = n.next
			return
		}
		prev = n
	}
}

func main() {
	// Создаем новую хеш-таблицу с 1000 корзинами
	h := &hashmap{buckets: make([]*node, 1000)}

	// Добавляем несколько пар ключ-значение
	h.Set("apple", "red")
	h.Set("banana", "yellow")
	h.Set("orange", "orange")

	// Получаем значения по ключам
	fmt.Println(h.Get("apple"))  // red true
	fmt.Println(h.Get("banana")) // yellow true
	fmt.Println(h.Get("grape"))  //  false

	// Удаляем пару по ключу
	h.Delete("banana")

	// Проверяем, что пара удалена
	fmt.Println(h.Get("banana")) //  false
}
