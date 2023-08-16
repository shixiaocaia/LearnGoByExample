package main

import (
	"fmt"
	"time"
)

// 使用channel实现一个限流器

func main() {
	chLimit := make(chan struct{}, 3)
	// 通过一个带有缓冲大小为3的无类型通道chLimit来控制并发数量，确保一次只有三个协程在进行

	for i := 0; i < 20; i++ {
		// 发送一个值，占用一个协程，chLimit是一个struct{}类型，并用{}表示空结构体
		chLimit <- struct{}{}
		go func(i int) {
			fmt.Println("下游服务处理逻辑...", i)
			time.Sleep(time.Second * 3)
			<-chLimit
			// 获取一个值，释放一个协程
		}(i)
	}
	// 保证所有协程全部处理完
	time.Sleep(30 * time.Second)
}

// 每次三个协程处理逻辑，释放后才能允许协程进行处理
