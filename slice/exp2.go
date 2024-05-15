package main

import "fmt"

func main() {
	a := make([]int, 0, 3)

	b := append(a, 1)
	// 	a ___
	//  b 1__

	x := append(b, 3)
	// 	a ___
	//  b 1__
	//  x 13_
	fmt.Println(x)
	b = append(b, 10)
	// 如果这个时候b[1] = 10，会影响x[1] = 10
	fmt.Println(b, x)
	// b = [1 10], x = [1, 10]

	_ = append(b, 5)
	fmt.Println(a, b, x)
	// a = [], b = [1 10], x = [1 10]
	// 此时在b的基础上append，没有影响到b，前面空占位符，区别上面x append
}
