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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	/* n := len(ar) */

	/* quickSort(ar, 0, n-1) */
	quickSort(ar)

	fmt.Println(ar)
}

/*
	 func partition(ar []int, low, high int) int {
		pivot := ar[high]
		i := low - 1
		for j := low; j < high; j++ {
			if ar[j] < pivot {
				i++
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
		ar[i+1], ar[high] = ar[high], ar[i+1]
		return i + 1
	}

	func quickSort(ar []int, low, high int) {
		if low < high {
			p := partition(ar, low, high)
			quickSort(ar, low, p-1)
			quickSort(ar, p+1, high)
		}
	}
*/
func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	left, right := 0, len(ar)-1

	pivot := rand.Int() % len(ar)

	ar[pivot], ar[right] = ar[right], ar[pivot]

	for i, _ := range ar {
		if ar[i] < ar[right] {
			ar[left], ar[i] = ar[i], ar[left]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])

	return ar
}
