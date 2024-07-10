package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	var count_receive_ch1 int64
	var count_receive_ch2 int64
	var wg sync.WaitGroup
	chance_chan := func() {
		for {
			chance := rand.Intn(100) - 1
			if chance <= 50 {
				chan1 <- chance
			}
			if chance > 50 {
				chance = rand.Intn(100) - 1
				if chance > 50 {
					chan2 <- chance
				} else {
					continue
				}
			}
		}
	}
	receive_chan := func() {
		for {
			select {
			case reciev := <-chan1:
				atomic.AddInt64(&count_receive_ch1, 1)
				var total = atomic.LoadInt64(&count_receive_ch1) + atomic.LoadInt64(&count_receive_ch2)
				if total > 0 {
					var ration = (float64(count_receive_ch1) / float64(total)) * 100
					str := fmt.Sprintf("Канал chan1 принимает %.2f%% сообщений, сообщение из канала: %d", ration, reciev)
					fmt.Println(str)
				}
			case reciev := <-chan2:
				atomic.AddInt64(&count_receive_ch2, 1)
				var total = atomic.LoadInt64(&count_receive_ch1) + atomic.LoadInt64(&count_receive_ch2)
				if total > 0 {
					var ration = (float64(count_receive_ch2) / float64(total)) * 100
					str := fmt.Sprintf("Канал chan2 принимает %.2f%% сообщений, сообщение из канала: %d", ration, reciev)
					fmt.Println(str)
				}
			default:
				fmt.Println("Сообщение отсутствует или не отправлено")
			}

		}
	}
	go chance_chan()
	go receive_chan()
	wg.Add(100)
	wg.Wait()
}
