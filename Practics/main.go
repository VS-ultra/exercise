package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//FizzBuzzFor()
	//CountFor50()
	// var x1, x2, y1, y2 float32
	// fmt.Scan(&x1, &y1, &x2, &y2)
	// FindCordinate(x1, y1, x2, y2)
	//SimleNumbe()
	/* 	var n int
	   	fmt.Print("The enter is number: ")
	   	fmt.Scan(&n)
	   	fmt.Printf("The inversion case: %v! = %v\n", n, FactorialInversion(n))
	   	fmt.Printf("The recursion case: %v! = %v", n, FactorialRecursion(n)) */
	/* 	var n int
	   	fmt.Print("The enter is number: ")
	   	fmt.Scan(&n)
	   	polindrom := PolindromNum(n)
	   	if polindrom == true {
	   		fmt.Println("This number is polindrom ")
	   	}
	   	if polindrom == false {
	   		fmt.Println("This number isn't polindrom")
	   	} */
	/*var s string
	fmt.Print("Enter the line: ")
	fmt.Scanln(&s)
	if PolindromNumString(s) == s {
		fmt.Println("This number is polindrom ")
		return
	}
	fmt.Println("This number isn't polindrom")
	*/
	/* 	var s string
	   	fmt.Print("Enter the line: ")
	   	fmt.Scanln(&s)
	   	if PolindromWords(s) == s {
	   		fmt.Println("This number is polindrom ")
	   		return
	   	}
	   	fmt.Println("This number isn't polindrom") */
	var s string
	fmt.Print("Enter the line: ")
	fmt.Scanln(&s)
	if len(s) == 0 {
		fmt.Print("S string is nil")
		return
	}
	fmt.Println("The result of the program execution: ", FindRepitWords(s))
}
func FizzBuzzFor() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
	}
}
func UpgradeFizzBuzzFor() {
	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
func CountFor50() {
	var sum int
	for i := 0; i < 50; i++ {
		sum += i
	}
	log.Println(sum)
}
func FindCordinate(a1, b1, a2, b2 float32) {
	if a1 == a2 {
		fmt.Println("No crossing")
	}
	x := (b2 - b1) / (a1 - a2)
	y := a1*x + b1
	fmt.Printf("x: %f, y: %f\n", x, y)
}
func SimleNumbe() {
	counter := 0
	for i := 3; counter <= 20; i++ {
		num := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				num++
			}
		}
		if num == 0 {
			fmt.Println(i)
			counter++
		}
	}
}
func FactorialInversion(n int) int {
	var f int
	f = 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}
func FactorialRecursion(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursion(n-1)
}
func PolindromNum(n int) bool {
	a, b := 0, n
	for b > 0 {
		a *= 10
		a += b % 10
		b /= 10
	}
	return a == n
}
func PolindromNumString(n string) string {
	str := ""
	for i := len(n) - 1; i >= 0; i-- {
		str += string(n[i])
	}
	return str
}

func FindRepitWords(s string) string {
	str := ""
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			str += string(s[i-1]) + strconv.Itoa(count+1)
			count = 0
		}
	}
	return str
}

/* func Archivate(s string) string {
	var curr rune
	currLen := 0
	res :=""
	for _,v := range s{
		if v != curr{
			if currLen > 0 {
				res+= fmt.Sprintf("%c%d",curr, currLen)
			}
			curr = v
			currLen = 1
		}else{
			currLen++
		}
	}
	res += fmt.Sprintf("%c%d", curr, currLen)
	return res
} */
/* func PolindromWords(s string) string{
	for i := len(s)-1; i >=0 i--{
		str += string(s[i])
	}
	return str
} */
