package main

import "fmt"

func main() {
	var i interface{} = nil
	var j interface{} = (*int)(nil)
	fmt.Println(i == nil, j == nil, i == j)
}

// i == nil：这里，我们将 i 初始化为 nil，因此 i == nil 返回 true。

// j == nil：尽管我们似乎将 j 初始化为 nil，但事实并非如此。我们实际上将 j 初始化为一个 nil 指针（指向 int 的指针）。
// 因此，j 的值是 nil，但其类型是 *int。在 Go 中，只有当接口的值和类型都为 nil 时，接口才被认为是 nil。因此，j == nil 返回 false。
// i == j：由于 i 和 j 的类型不同（一个没有类型，另一个是 *int），因此它们是不相等的，所以返回 false。
