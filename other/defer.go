package main

import "fmt"

func test1() (result int) {
	defer func() {
		result++
	}()
	return 0

	/*
		result = 0
		defer func(){
			result++
		}
		return

		非匿名返回值
	*/
}

func test2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t

	/*
		t := 5
		r = t           // 看到最后return t
		defer func(){
			t = t + 5   // 值传递，不改变 r
		}()
		return
	*/
}

func test3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1

	/*
		r = 1
		defer func(r int){
			 r = r + 5 //值传递
		}(r)
		return
		// NOTE 如果这里不是传参，直接 r + 5,就是对非匿名返回值修改，对比test1
	*/
}

func test4() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
	// defer函数的顺序后入先出，类似堆栈，先输出defer2，再输出defer1
	// 返回值是匿名返回值，不影响
}

func test5(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	return b

	// b = 0
	// return b
}

func main() {
	fmt.Println("test1:", test1())
	fmt.Println("test2:", test2())
	fmt.Println("test3:", test3())
	fmt.Println("test4:", test4())
	fmt.Println("test5:", test5(1))
}
