package main

import (
	"fmt"
	"sync"
)

// N个goroutine循环打印数字min ~ max范围的数字

var max = 100
var min = 1
var n = 5

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(n)

	var chanArray = make([]chan int, n)
	for i := 0; i < n; i++ {
		chanArray[i] = make(chan int)
	}

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			// 注释掉这一句后 fatal error: all goroutines are asleep - deadlock!
			// 在num > max时return, 触发defer close关闭下一个通道
			// 下一个goroutine的chanArray[i]被关闭停止阻塞，接着关闭下一个通道，反复
			defer close(chanArray[(i+1)%n])
			// channel0为空时阻塞, 被关闭时停止阻塞
			for num := range chanArray[i] {
				if num > max {
					return
				}
				fmt.Printf("g %d 进行打印 %d \n", i, num)
				// 自增，传递给下一个goroutine(轮询传递）
				chanArray[(i+1)%n] <- num + 1
			}
		}(i)
	}

	// 传递第一个数，从main goroutine-> goroutine 0
	chanArray[0] <- min
	wg.Wait()
}
