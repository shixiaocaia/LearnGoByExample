package main

import (
    "fmt"
    "time"
)

func main() {
    // ch := make(chan int, 0)
	ch := make(chan int, 3)

    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Sending:", i)
            ch <- i
        }
		// 如果没有缓存的channel的，发送会阻塞等待读取
        close(ch) 
		fmt.Println("channel closed...")
    }()

    // 从通道中读取数据
    for val := range ch {
        fmt.Println("Received:", val)
        time.Sleep(2 * time.Second) // 等待2秒钟，模拟处理数据的时间
    }
}
