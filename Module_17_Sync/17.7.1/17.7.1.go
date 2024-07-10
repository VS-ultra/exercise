package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	valueChan chan int
}

func NewCounter() *Counter {
	return &Counter{valueChan: make(chan int, 1)}
}

func (c *Counter) Increment(done chan bool, endValue int) {
	val := <-c.valueChan
	val++
	c.valueChan <- val
	if val >= endValue {
		done <- true
	}
}

func (c *Counter) Value() int {
	return <-c.valueChan
}

func main() {
	var wg sync.WaitGroup
	var goroutines, endValue int
	counter := NewCounter()
	fmt.Print("Enter the number of goroutines: ")
	fmt.Scanln(&goroutines)
	fmt.Print("Enter the final counter value: ")
	fmt.Scanln(&endValue)

	counter.valueChan <- 0
	done := make(chan bool, 1)

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					counter.Increment(done, endValue)
				}
			}
		}()
	}

	wg.Wait()
	close(counter.valueChan)
	finalValue := counter.Value()
	fmt.Printf("Target value %d reached.\n", endValue)
	fmt.Printf("Final counter value:%d\n", finalValue)
}
