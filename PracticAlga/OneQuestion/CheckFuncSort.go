package main

import "fmt"

func checkSliceIsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pointStart := len(a) / 2
	a[pointStart], a[right] = a[right], a[pointStart]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a
}
func main() {
	a := []int{}
	fmt.Println("Enter arrays")
	fmt.Scanln(a)
	fmt.Println("Our array: ", a)
	fmt.Println("Our array sort quick: ", QuickSort(a))
	fmt.Println("Check in our array sort :", checkSliceIsSorted(a))

}
