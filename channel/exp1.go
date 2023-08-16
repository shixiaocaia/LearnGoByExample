package main

// channel本质是值拷贝

import (
	"fmt"
	"time"
)

func print(u <-chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("print int", <-u)
}

type people struct {
	name string
}

var u = people{name: "A"}

func printPeople(u <-chan *people) {
	time.Sleep(2 * time.Second)
	a := <-u
	fmt.Printf("&a address-3: %p\n", a)
	fmt.Println("print people", a)
}

func main() {
	c := make(chan *people, 5)
	var a = &u
	c <- a
	fmt.Println(a)
	fmt.Printf("&a address-1: %p\n", a)
	a = &people{name: "xiaocai"}
	fmt.Printf("&a address-2: %p\n", a)

	go printPeople(c)
	time.Sleep(5 * time.Second)
	fmt.Printf("&a address-4: %p\n", a)
	fmt.Println(a)
}
