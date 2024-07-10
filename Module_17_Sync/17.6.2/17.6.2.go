package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			chan1 <- rand.Int()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			chan2 <- rand.Int()
		}
	}()

	go func() {
		for {
			fmt.Println(time.Now().Format("15:04:05"))
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-chan1:
			fmt.Printf("Получено сообщение из chan1: %d\n", msg1)
		case msg2 := <-chan2:
			fmt.Printf("Получено сообщение из chan2: %d\n", msg2)
		}
	}
}
