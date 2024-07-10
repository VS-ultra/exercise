/*
	package main

import (

	"fmt"
	"sync"
	"time"

)

	func main() {
		mu := sync.Mutex{}
		_ = mu

		times := 5
		value := 0

		go func() {
			for {
				time.Sleep(time.Millisecond)
				go func() {
					for {
						time.Sleep(time.Millisecond)
						value++
					}
				}()
			}
		}()

		time.Sleep(time.Second * time.Duration(times))
		fmt.Println(value)
	}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите количество секунд работы программы: ")
	fmt.Scan(&n)

	value := 0

	ticker := time.NewTicker(200 * time.Millisecond)
	timer := time.NewTimer(time.Duration(n) * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(value)
			value++
		case <-timer.C:
			ticker.Stop()
			return
		}
	}
}
