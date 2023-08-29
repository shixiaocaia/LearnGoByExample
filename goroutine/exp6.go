package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 创建goroutine
// 1. 初始化new一个g
// 2. 存放p：runnext:下一个执行的g runq: 一个本地P的G队列 256数组，实际上可以存放256 + 1 globalg: 全局的g队列

func main() {
	// 1. 设置一个 P工作
	runtime.GOMAXPROCS(1)

	n := 258

	wg := &sync.WaitGroup{}
	wg.Add(n)

	// runnext: g2 -> g1
	// runq: g1存放到runq
	// 最终 runnext: g10, runq = g1~g9
	// 由于runnext是下一个执行的g，所以先输出g10

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("I am goroutine ", i)
		}(i)
	}

	// 当g超过257，超过本地队列能存放
	// 会把当前队列（满）一半， + runnext（旧的）存放到全局队列中，然后执行顺序就变了
	// 1-128 + 257 放到全局G队列中
	// 调度器每61次会从全局G中抽取1个执行一次，避免一直执行本地
	// 实际会在61次前，因为中间穿插一些系统级别的

	wg.Wait()
}
