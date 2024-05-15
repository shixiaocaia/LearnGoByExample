package main

import "fmt"

type DoubleListNode struct {
	key   int
	value int
	Pre   *DoubleListNode
	Next  *DoubleListNode
}

type LRUCache struct {
	maxCapacity int
	nowCapacity int
	checkMap    map[int]*DoubleListNode
	head        *DoubleListNode
	tail        *DoubleListNode
}

func NewLRUCache(cap int) LRUCache {
	headNode := &DoubleListNode{}
	tailNode := &DoubleListNode{}
	headNode.Next = tailNode
	tailNode.Next = headNode
	return LRUCache{
		cap,
		0,
		make(map[int]*DoubleListNode),
		headNode,
		tailNode,
	}
}

func (this *LRUCache) InsertNode(node *DoubleListNode) {
	tmp := this.head.Next
	this.head.Next = node
	node.Pre = this.head
	node.Next = tmp
	tmp.Pre = node
	// head.Next->node->tmp
}

func (this *LRUCache) DeleteNode(node *DoubleListNode) {
	tmp := node.Next
	tmp.Pre = node.Pre
	node.Pre.Next = tmp
	// node1 -> deleteNode -> tmp 改变指向
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.checkMap[key]; !ok {
		return -1
	} else {
		this.DeleteNode(value)
		this.InsertNode(value)
		return value.value
	}
}

func (this *LRUCache) Put(key, value int) {
	if val, ok := this.checkMap[key]; ok {
		// 存在, 删除node, 插入新node
		this.DeleteNode(val)
		newNode := &DoubleListNode{key, value, nil, nil}
		this.InsertNode(newNode)
		this.checkMap[key] = newNode
	} else if !ok && this.maxCapacity < this.nowCapacity+1 {
		// 不存在, 超出capacity, 删除tailNode, 插入新node
		t := this.tail.Pre
		delete(this.checkMap, t.key)
		this.DeleteNode(t)

		newNode := &DoubleListNode{key, value, nil, nil}
		this.InsertNode(newNode)
		this.checkMap[key] = newNode
	} else {
		// 不存在, 不超过capacity
		newNode := &DoubleListNode{key, value, nil, nil}
		this.InsertNode(newNode)
		this.checkMap[key] = newNode
		this.nowCapacity++
	}
}

func main() {
	lruCache := NewLRUCache(3)

	// get notFound
	fmt.Printf("get key-%v: %v\n", 1, lruCache.Get(1))

	// get 1
	lruCache.Put(1, 100)
	lruCache.Get(1)
	fmt.Printf("get key-%v: %v\n", 1, lruCache.Get(1))

	// put 已存在key
	lruCache.Put(1, -100)
	lruCache.Get(1)
	fmt.Printf("get key-%v: %v\n", 1, lruCache.Get(1))

	// put 大于cap
	lruCache.Put(0, -100)
	lruCache.Put(2, -100)
	lruCache.Put(3, -100)
	fmt.Println("lruCache capacity:", lruCache.nowCapacity)
	lruCache.Put(4, -100)
	// 1已经被删除
	lruCache.Get(1)
	fmt.Printf("get key-%v: %v\n", 1, lruCache.Get(1))
	lruCache.Get(2)
	fmt.Printf("get key-%v: %v\n", 2, lruCache.Get(2))

}
