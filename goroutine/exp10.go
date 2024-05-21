package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(10)

	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			if ok := mutex.TryLock(); ok {
				ch <- i
				fmt.Println(i, "from redis")
				time.Sleep(3 * time.Second)
			} else {
				num := <-ch
				fmt.Println(num, "from channel")
				ch <- i
				return
			}
			defer mutex.Unlock()
		}(i)
	}

	wg.Wait()
}
