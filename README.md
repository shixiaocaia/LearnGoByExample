# Learn Go By Example

Hello World.

## channel

exp1.go: channel本质是值的拷贝

exp2.go: 使用channel实现一个限流器

exp3.go: 使用channel实现一个互斥锁

exp4.go: go select使用

exp5.go: 判断channel情况

## goroutine

exp1.go: goroutine基本使用

exp2.go: 使用三个协程，每秒钟依次打印cat dog，fish，要求顺序不能变化

exp3.go:
- 有四个goroutine，编号为1，2，3，4，每秒钟会有一个goroutine打印出自己的编号
- 要求写一个程序，让输出的编号总是按照1，2，3，4顺序打印出来

exp4.go: 实现两个协程轮流输出A-1, B-2, C-3

exp5.go: N个goroutine循环打印数字min ~ max范围的数字

exp6.go:
- goroutine的创建过程，g是如何存放，如何调度的
- 在GMP模型中，P起到了重要作用，负责G和M之间调度
    - 当前G阻塞，将P转到其他闲置或者新建M继续执行

exp7.go: 使用两个协程交替输出奇数偶数

exp8.go: 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
`12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728`

## other
defer.go: 输出顺序，返回值修改

interface.go: 类型的比较，interface类型只有值和类型都为nil，才为nil

sort.go: 经典排序算法

string1.go: 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
  - 判断两个给定的字符串排序后是否一致，同这一题，可以判断两个字符串中同一个字符的个数是否相同