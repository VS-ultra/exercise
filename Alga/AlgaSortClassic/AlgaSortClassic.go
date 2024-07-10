package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func usedMemory() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
func main() {
	arraySize := 20000
	array := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		array[i] = rand.Intn(arraySize)
	}

	algoritms := map[string]func([]int){
		"Selection sort": SelectionSort,
		"Insertion sort": InsertionSort,
		"Bubble sort":    BubbleSort,
		"Quick sort":     func(arr []int) { QuickSort(arr) },
		"Merge sort":     func(arr []int) { MergeSort(arr) },
	}

	for name, sortFunction := range algoritms {
		tempArray := make([]int, arraySize)
		copy(tempArray, array)

		runtime.GC()
		startMemory := usedMemory()

		start := time.Now()
		sortFunction(tempArray)
		elapsed := time.Since(start)

		endMemory := usedMemory()

		fmt.Printf("%s - Время выполнения %v, Потребляемая память %v\n", name, elapsed, endMemory-startMemory)
	}
}

// BubbleSort выполняет сортировку пузырьком.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// InsertionSort выполняет сортировку вставками.
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// SelectionSort выполняет сортировку выбором.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// QuickSort выполняет быструю сортировку.
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

// MergeSort выполняет сортировку слиянием.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge объединяет два отсортированных массива в один.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
