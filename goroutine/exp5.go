package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10
	ch := make([]chan int, n)
	for i := 0; i < n; i++ {
		ch[i] = make(chan int)
		go xPrint(ch[i], i)
	}

	for i := 1; i <= 100; i++ {
		ch[i%n] <- i
		time.Sleep(time.Second)
	}

	// 关闭所有通道
	// 防止潜在错误，向closed channel 发送数据会panic
	// 避免内存泄露
	for _, c := range ch {
		close(c)
	}

	// 主协程可以选择在此等待一段时间或进行其他操作
	// 等待足够的时间以确保所有输出都已完成
	time.Sleep(time.Second * 1)
	fmt.Println("All channels are closed, and all goroutines should have exited.")
}

func xPrint(ch chan int, i int) {
	for num := range ch {
		fmt.Printf("G:%d, print number:%d\n", i, num)
	}
}
