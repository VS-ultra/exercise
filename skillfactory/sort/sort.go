package sort

import (
	"math/rand"
)

func BubbleSortRecursive(ar []int, lastIndex int) {
	if lastIndex == 0 {
		return
	}
	for i := 0; i < lastIndex-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
		}
	}
	BubbleSortRecursive(ar, lastIndex-1)
}
func InsertionSort(ar []int) {
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
func QuickSort(ar []int) []int {
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

	QuickSort(ar[:left])
	QuickSort(ar[left+1:])

	return ar
}
func MergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	middle := len(ar) / 2

	sortedAr := make([]int, 0, len(ar))
	left, right := MergeSort(ar[:middle]), MergeSort(ar[middle:])

	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortedAr = append(sortedAr, right[j])
			j++
		} else {
			sortedAr = append(sortedAr, left[i])
			i++
		}
	}

	sortedAr = append(sortedAr, left[i:]...)
	sortedAr = append(sortedAr, right[j:]...)

	return sortedAr
}
