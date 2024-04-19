package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

// 测试1：不执行resp.Body.Close()，也不执行ioutil.ReadAll(resp.Body)
func testFun1() {
	for i := 0; i < 5; i++ {
		_, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println("testFun1 http.Get err: ", err)
			return
		}
	}

	time.Sleep(time.Second * 1)
	fmt.Println("testFun1() 当前goroutine数量=", runtime.NumGoroutine())
}

// 测试2：执行resp.Body.Close()，不执行ioutil.ReadAll(resp.Body)
func testFun2() {
	for i := 0; i < 5; i++ {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println("testFun2 http.Get err: ", err)
			return
		}

		resp.Body.Close() // 执行resp.Body.Close()
	}

	// Close()过程需要一定时间，如果直接输出goroutine数量，可能出现连接还未完全回收的情况，结果有时为1有时为3
	// 因此，为了结果的准确性，我们这里休眠等待1秒，使得连接完全被回收。
	time.Sleep(time.Second * 1)
	fmt.Println("testFun2() 当前goroutine数量=", runtime.NumGoroutine())
}

// 测试3：不执行resp.Body.Close()，执行ioutil.ReadAll(resp.Body)
func testFun3() {
	for i := 0; i < 5; i++ {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println("testFun3 http.Get err: ", err)
			return
		}

		_, _ = ioutil.ReadAll(resp.Body) // 执行ioutil.ReadAll(resp.Body)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("testFun3() 当前goroutine数量=", runtime.NumGoroutine())
}

// 测试4：执行resp.Body.Close()，也执行ioutil.ReadAll(resp.Body)
func testFun4() {
	for i := 0; i < 5; i++ {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println("testFun4 http.Get err: ", err)
			return
		}

		_, _ = ioutil.ReadAll(resp.Body) // 执行ioutil.ReadAll(resp.Body)
		resp.Body.Close()                // 执行resp.Body.Close()
	}

	time.Sleep(time.Second * 1)
	fmt.Println("testFun4() 当前goroutine数量=", runtime.NumGoroutine())
}

func main() {
	// 2 * 5 + 1 = 每个http.Get()会创建两个goroutine，一个用于发送请求，一个用于接收响应，再加上主goroutine，共11个goroutine
	// testFun1()

	// 执行 resp.Body.Close()后，回收了底层都会创建两个协程，只剩下主协程本身，所以一共就 1 个协程
	// testFun2()

	// 将 resp.body 读取完之后，会把当前的连接放入空闲列表中，供下次复用，所以算上主协程本身一共就 3 个协程
	// testFun3()

	// 在于 执行了 ioutil.ReadAll(resp.Body)，将 resp.body 读取完之后，会优先把当前的连接放入空闲列表中，
	// 供下次复用，即使你执行了resp.Body.Close()，所以算上主协程本身一共就 3 个协程。
	testFun4()
}
