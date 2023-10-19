package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 收到取消信号，退出
				fmt.Println("Goroutine 收到取消信号，退出")
				return
			default:
				// 执行正常操作
				fmt.Println("Goroutine 执行正常操作")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// 等待10秒，触发取消信号
	time.Sleep(10 * time.Second)

	// 手动取消goroutine
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Main 退出")

}
