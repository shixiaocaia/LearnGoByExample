package main

import "fmt"

func main() {
	case1 := make(chan int)
	case2 := make(chan int)

	close(case1)
	close(case2)

	// 随机选择case执行，平等访问
	select {
	// 读 closed channel 为 对应nil值
	case <-case1:
		fmt.Println("case1")
	// 写 closed channel panic
	case case2 <- 1:
		fmt.Println("case2")
	default:
		fmt.Println("default")
	}
}
