package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)
	cSend := func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}
	cRecv := func() {
		for v := range c {
			fmt.Println(v)
		}
		wg.Done()
	}
	go cSend()
	go cRecv()
	wg.Add(2)
	wg.Wait()
}
