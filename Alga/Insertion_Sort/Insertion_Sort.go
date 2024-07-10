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

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}
	for i := 1; i < len(ar)-1; i++ {
		if ar[i] < ar[i-1] {
			for j := i; j > 0 && ar[j] < ar[j-1]; j-- {
				if ar[j] < ar[j-1] {
					ar[j], ar[j-1] = ar[j-1], ar[j]
				}
			}
		}
	}
}
