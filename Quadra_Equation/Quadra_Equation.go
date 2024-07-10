package main

import (
	"fmt"
	"math"
)

var (
	ErrZeroA       = fmt.Errorf("coefficient a is zero")
	ErrNoRealRoots = fmt.Errorf("equation has no real roots")
	x1, x2         float64
)

func SolveQuadraticEquation(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 {
		err = ErrZeroA
		return
	}
	D := b*b - 4*a*c
	if D < 0 {
		err = ErrNoRealRoots
		return
	}
	if D == 0 {
		x1 = -b / (2 * a)
		x2 = x1
		return
	}
	dRoot := math.Sqrt(D)
	x1 = (-b + dRoot) / (2 * a)
	x2 = (-b - dRoot) / (2 * a)
	return
}
func main() {
	var a, b, c float64
	fmt.Println("Enter arguments quadra equation")
	fmt.Scanln(&a, &b, &c)
	SolveQuadraticEquation(a, b, c)
	x1, x2, err := SolveQuadraticEquation(a, b, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)
	return
}
