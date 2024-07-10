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

	selectionSort(ar)

	fmt.Println(ar)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")

	selectionSortByMax(ar)

	fmt.Println(ar)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")

	BidirectionalSelectionSort(ar)
	fmt.Println(ar)
}

func selectionSort(ar []int) {
	if len(ar) == 0 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}
func selectionSortByMax(ar []int) {
	if len(ar) == 0 {
		return
	}
	for i := len(ar) - 1; i > 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}
func BidirectionalSelectionSort(ar []int) {
	if len(ar) == 0 {
		return
	}
	minPoint := 0
	maxPoint := len(ar) - 1
	for minPoint < maxPoint {
		minIndex := minPoint
		for i := minPoint + 1; i <= maxPoint; i++ {
			if ar[i] < ar[minIndex] {
				minIndex = i
			}
		}
		maxIndex := maxPoint
		for i := maxPoint - 1; i >= minIndex; i-- {
			if ar[i] > ar[maxIndex] {
				maxIndex = i
			}
		}
		ar[minPoint], ar[minIndex] = ar[minIndex], ar[minPoint]
		ar[maxPoint], ar[maxIndex] = ar[maxIndex], ar[maxPoint]
		minPoint++
		maxPoint--
	}

}
