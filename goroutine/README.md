exp1.go 
- goroutine基本使用

exp2.go
- 使用三个协程，每秒钟依次打印cat dog，fish，要求顺序不能变化

exp3.go 
- 有四个goroutine，编号为1，2，3，4，每秒钟会有一个goroutine打印出自己的编号
- 要求写一个程序，让输出的编号总是按照1，2，3，4顺序打印出来

exp4.go
- 实现两个协程轮流输出A-1, B-2, C-3

exp5.go
- N个goroutine循环打印数字min ~ max范围的数字