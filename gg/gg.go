package main

import (
	"fmt"
)

func main() {
	doSomething(func(number int) { fmt.Printf("number: %d", number) })
}

func doSomething(anonymousFunc func(number int)) {
	anonymousFunc(100)
}
