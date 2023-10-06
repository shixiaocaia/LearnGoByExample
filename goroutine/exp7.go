package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用两个协程交替输出奇数偶数
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chanA := make(chan int, 1)
	chanB := make(chan int, 1)
	n := 100

	chanA <- 1

	go func() {
		for {
			cnt, ok := <-chanA
			if !ok {
				return
			}
			fmt.Println("G1: ", cnt)
			if cnt == n {
				close(chanB)
				wg.Done()
				return
			}
			time.Sleep(time.Second)
			chanB <- cnt + 1
		}
	}()

	go func() {
		for {
			cnt, ok := <-chanB
			if !ok {
				wg.Done()
				return
			}
			fmt.Println("G2: ", cnt)
			if cnt == n {
				close(chanA)
				wg.Done()
				return
			}
			time.Sleep(time.Second)
			chanA <- cnt + 1
		}
	}()

	wg.Wait()
}
