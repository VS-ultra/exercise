package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	bubbleSort(ar)
	fmt.Println(ar)
	fmt.Println("////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	bubbleSortReversed(ar)
	fmt.Println(ar)
	bubbleSortRecursive(ar, len(ar))
	fmt.Println(ar)
}

func bubbleSort(ar []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				swapped = true
			}
		}
	}
	return
}
func bubbleSortReversed(ar []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := len(ar) - 1; i > 0; i-- {
			if ar[i] > ar[i-1] {
				ar[i], ar[i-1] = ar[i-1], ar[i]
				swapped = true
			}
		}
	}
}
func bubbleSortRecursive(ar []int, lastIndex int) {
	if lastIndex == 0 {
		return
	}
	for i := 0; i < lastIndex-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
		}
	}
	bubbleSortRecursive(ar, lastIndex-1)
}
