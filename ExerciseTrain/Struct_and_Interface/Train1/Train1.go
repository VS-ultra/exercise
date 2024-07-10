package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func NewCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}
func main() {
	var r float64
	fmt.Println("Enter radius Circle: ")
	fmt.Scan(&r)
	circle := NewCircle(r)
	fmt.Println("Output Area Circle is ", circle.Area())
	fmt.Println("Output Perimeter is ,", circle.Perimeter())
}
