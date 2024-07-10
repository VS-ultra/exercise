package main

import "fmt"

func main() {
	var array = []int{1, 2, 4, 5, 6}
	fmt.Println(findMaxNegative(array))
}
func findMaxNegative(array []int) (min int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}
	min = array[0]
	for _, val := range array[1:] {
		if val < min {
			min = val
		}
	}

	return min, nil
}
