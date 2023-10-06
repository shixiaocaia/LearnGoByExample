package main

import "fmt"

func main() {
	i := [3]int{10, 10, 10}
	j := []int(i[2:3])
	i[2] = 900
	fmt.Println(j[0])
	fmt.Printf("cap(j) = %d\n", cap(j))
	fmt.Printf("j[0]: %p, i[2]: %p\n", &j[0], &i[2])
	// 扩容
	j = append(j, 2)
	fmt.Printf("cap(j) = %d\n", cap(j))
	i[2] = 100
	fmt.Println(j[0])
	fmt.Printf("j[0]: %p, i[2]: %p", &j[0], &i[2])
}
