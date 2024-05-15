package main

import (
	"fmt"
	"sync"
)

//实现两个协程轮流输出A-1, B-2, C-3

// 一个输出字母，一个输出数字

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chanA := make(chan int, 1)
	chanNum := make(chan int, 1)
	chanNum <- 0
	go func() {
		defer wg.Done()
		for i := 65; i <= 90; i++ {
			<-chanNum
			//time.Sleep(time.Second)
			fmt.Printf("%c-", byte(i))
			//time.Sleep(time.Second)
			chanA <- 1
		}
		return
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-chanA
			fmt.Printf("%d ", i)
			chanNum <- 1
		}
		return
	}()

	wg.Wait()

}
