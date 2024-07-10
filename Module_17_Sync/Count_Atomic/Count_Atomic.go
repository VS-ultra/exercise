/* package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 10

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup
	increment := func() {
		defer wg.Done()
		atomic.AddInt64(&counter, step)
	}
	// Не всегда вычисление этой переменной будет приводить к верному
	// результату в счётчике, но для правильных значений
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / step)
	for i := 1; i <= iterationCount; i++ {
		wg.Add(1)
		go increment()
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const step int64 = 1
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func(i int) {
		atomic.AddInt64(&counter, step)
		if i == interationAmount {
			c.Signal()
		}
	}
	for i := 1; i <= interationAmount; i++ {
		go increment(i)
	}
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
*/
package main

import (
	"fmt"
	"sync"
)

const step int = 1
const interationAmount int = 1000

func main() {
	var counter int = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func() {
		c.L.Lock()                       // Захватываем блокировку
		counter += step                  // Увеличиваем значение counter на единицу
		if counter == interationAmount { // Проверяем, достигло ли counter конечного значения
			c.Signal() // Посылаем сигнал горутине, которая ждет в main
		}
		c.L.Unlock() // Освобождаем блокировку
	}
	for i := 1; i <= interationAmount; i++ {
		go increment() // Запускаем горутину с функцией increment
	}
	c.L.Lock()           // Захватываем блокировку
	c.Wait()             // Ждем сигнала от последней горутины
	c.L.Unlock()         // Освобождаем блокировку
	fmt.Println(counter) // Печатаем результат, надеясь, что будет 1000
}
