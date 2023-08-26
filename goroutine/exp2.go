package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用三个协程，每秒钟依次打印cat dog，fish，要求顺序不能变化
func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	chcat := make(chan int)
	chdog := make(chan int)
	chfish := make(chan int)

	go printCat(&wg, chcat, chfish)
	go printDog(&wg, chdog, chcat)
	go printFish(&wg, chfish, chdog)

	wg.Wait()
}

func printCat(wg *sync.WaitGroup, cat chan int, fish chan int) {
	defer wg.Done()
	for {
		cat <- 1 // 往cat里传，表明cat已经输出，chdog收到后输出
		fmt.Println("cat")
		time.Sleep(time.Second)
		<-fish // 阻塞等待，收到chfish输入，表明执行下一轮
	}
}

func printDog(wg *sync.WaitGroup, dog chan int, cat chan int) {
	defer wg.Done()
	for {
		<-cat
		// 因为是无缓存的，cat和dog会同时打印，这里做一个延迟
		time.Sleep(time.Second)
		fmt.Println("dog")
		dog <- 1
	}
}

func printFish(wg *sync.WaitGroup, fish chan int, dog chan int) {
	defer wg.Done()
	for {
		<-dog
		time.Sleep(time.Second)
		fmt.Println("fish")
		time.Sleep(time.Second)
		fish <- 1
	}
}
