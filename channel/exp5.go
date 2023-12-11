package main

import "fmt"

func main() {
	c := make(chan int, 1)

	done := false
	for !done {
		select {
		// 读不到
		// 第二次读到，输出 1，channel = nil
		case <-c:
			fmt.Print(1)
			c = nil
		// 写入 1，输出 2
		case c <- 1:
			fmt.Print(2)
		// channel 关闭后，非阻塞进入 default
		default:
			fmt.Print(3)
			done = true
		}
	}
}
