package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 定义一个WaitGroup
var wg sync.WaitGroup

func say(name string) {
	time.Sleep(time.Second)
	fmt.Println("I am done, by " + name)
	wg.Done()
}

func main() {
	wg.Add(2) //2. 等待2个goroutine，等待执行两个任务

	go func(id string) {
		fmt.Println(id)
		wg.Done() //3. 完成任务后，减少计数
	}("Hello")

	go say("xiaocai")

	wg.Wait() //4. 挂起，等待所有goroutine完成，如果wg != 0 阻塞
	fmt.Println("exit")

}
