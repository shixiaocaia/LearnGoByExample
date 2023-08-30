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
		r = t // 看到最后return t
		defer func(){
			t = t + 5 // 值传递，不改变r
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

		// NOTE 如果这里不是传参，直接 r + 5,就是对非匿名返回值修改
	*/
}

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
}
