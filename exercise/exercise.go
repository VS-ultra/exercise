package main

import "fmt"

const (
	January = iota + 1
	February
	March
	April
	May
)

func main() {
	fmt.Println("January:", January)
	fmt.Println("February:", February)
	fmt.Println("March:", March)
	fmt.Println("April:", April)
	fmt.Println("May:", May)
}
