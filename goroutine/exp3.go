package main

import (
	"fmt"
	"sync"
	"time"
)

// 有四个goroutine，编号为1，2，3，4，每秒钟会有一个goroutine打印出自己的编号，
// 要求写一个程序，让输出的编号总是按照1，2，3，4顺序打印出来
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(4)

	go print1(&wg, ch1, ch4)
	go print2(&wg, ch2, ch1)
	go print3(&wg, ch3, ch2)
	go print4(&wg, ch4, ch3)

	wg.Wait()
}

func print1(wg *sync.WaitGroup, ch1 chan struct{}, ch4 chan struct{}) {
	defer wg.Done()
	for {
		ch1 <- struct{}{}
		fmt.Println(1)
		time.Sleep(time.Second)
		<-ch4
	}
}

func print2(wg *sync.WaitGroup, ch2 chan struct{}, ch1 chan struct{}) {
	defer wg.Done()
	for {
		<-ch1
		time.Sleep(time.Second)
		fmt.Println(2)
		ch2 <- struct{}{}
	}
}

func print3(wg *sync.WaitGroup, ch3 chan struct{}, ch2 chan struct{}) {
	defer wg.Done()
	for {
		<-ch2
		time.Sleep(time.Second)
		fmt.Println(3)
		ch3 <- struct{}{}
	}
}

func print4(wg *sync.WaitGroup, ch4 chan struct{}, ch3 chan struct{}) {
	defer wg.Done()
	for {
		<-ch3
		time.Sleep(time.Second)
		fmt.Println(4)
		time.Sleep(time.Second)
		ch4 <- struct{}{}
	}
}
