package main

import "fmt"

func main() {
	doAppend := func(s []int) {
		s = append(s, 1)
		fmt.Printf("addr %p\n", &s)
		printLengthAndCapacity(s)
	}
	// len = 8,所以默认分配了8个0
	s := make([]int, 8, 8)
	// 截取了[0,4),公用一片地址
	// 在末尾地址追加的位置还在s内，修改影响到了原s
	doAppend(s[:4])
	printLengthAndCapacity(s)
	// 接着对len = 8, cap = 8的切片append，超过cap，发生扩容
	// 没有大于扩容前cap的两倍，变成扩容前两倍容量
	doAppend(s)
	// 发生值拷贝
	printLengthAndCapacity(s)

	slice := make([]int, 0, 2)
	fmt.Printf("slice len = %d, cap = %d, addr = %p\n", len(slice), cap(slice), &slice)

	// append会产生一个新slice
	slice1 := append(slice, 1)
	fmt.Printf("slice1, len = %d, cap = %d, addr = %p\n", len(slice1), cap(slice1), &slice1)

	// 完全复制一个slice，分配到不同的内存空间
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice)
	fmt.Printf("slice1 addr = %p, sliceCopy addr = %p", &slice1, &sliceCopy)
}

func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}
