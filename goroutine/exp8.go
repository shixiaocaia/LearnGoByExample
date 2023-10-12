package main

import (
	"fmt"
)

// 交替打印字数字和字母

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-ch1:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				ch2 <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			// j 这里是rune类型，对应A的ASCII码，可以用于自增
			select {
			case <-ch2:
				if j > 'Z' {
					ch3 <- true
					return
				}
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				ch1 <- true
			}
		}
	}()
	// 初始情况
	ch1 <- true
	// 阻塞等待ch1，ch2输出完成
	<-ch3
}
