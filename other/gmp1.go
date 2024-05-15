//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 区别这里
	runtime.GOMAXPROCS(1)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println(1)
		wg.Done()
	}()

	go func() {
		fmt.Println(2)
		wg.Done()
	}()

	go func() {
		fmt.Println(3)
		wg.Done()
	}()

	wg.Wait()
}
