package main

import "fmt"

func main() {
	var a, b float64
	var f string
	fmt.Scanln(&a)
	fmt.Scanln(&f)
	fmt.Scanln(&b)
	switch f {
	case "+":
		fmt.Println("Результат суммирования:", a+b)
	case "-":
		fmt.Println("Результат вычитания:", a-b)
	case "*":
		fmt.Println("Результат умножения:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на 0")
		} else {
			fmt.Println("Результат деления:", a/b)
		}
	}
}
