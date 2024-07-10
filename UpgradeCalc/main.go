package main

import (
	"UpgradeCalc/calc"
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	c := calc.NewCalculator()
	fmt.Println("Enter two numbers and an operator in the console")
	fmt.Println("Введите в консоль два номера и оператора")
	fmt.Scanln(&num1, &num2, &operator)
	ent_result := c.Calculator(num1, num2, operator)
	fmt.Println("The output result:", ent_result)
}
