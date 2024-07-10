package main

import "fmt"

func Intersection(mas1, mas2 []int) (c []int) {
	mapa := make(map[int]bool)
	for _, i := range mas1 {
		mapa[i] = true
	}
	for _, i := range mas2 {
		if _, ok := mapa[i]; ok {
			c = append(c, i)
		}
	}
	return
}

func main() {
	var amount1, amount2, num int
	fmt.Println("Enter first array size:")
	if _, err := fmt.Scanln(&amount1); err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	fmt.Println("Enter second array size:")
	if _, err := fmt.Scanln(&amount2); err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	if amount1 < 0 && amount2 < 0 {
		fmt.Println("Array size cannot be negative")
		return
	}
	mas1 := make([]int, amount1)
	mas2 := make([]int, amount2)
	fmt.Println("Enter first array:")
	for i := 0; i < amount1; i++ {
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		mas1[i] = num
	}
	fmt.Println("Enter second array:")
	for i := 0; i < amount2; i++ {
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		mas2[i] = num
	}
	if len(mas1) == 0 && len(mas2) == 0 {
		fmt.Println("Array cannot be empty")
		return
	}
	fmt.Println("Outside data: ", Intersection(mas1, mas2))
}
