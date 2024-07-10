package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	n := time.Now()
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j <= 10; j++ {
				fmt.Println(j)
			}
			wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println(time.Since(n))

}
