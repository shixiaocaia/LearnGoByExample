package main

import (
	"log"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	// 从通道内接受一个值，此时为空，阻塞当前协程
	<-m.ch
}

func (m *Mutex) Unlock() {
	select {
	// 发送一个空结构体，释放一个协程，相当于解锁
	case m.ch <- struct{}{}:
	default:
		// 发送失败，解锁了一个已经解锁的互斥锁
		panic("unlock of unlocked mutex")
	}
}

func (m *Mutex) TryLock() bool {
	select {
	// 初始化锁时，发送了一个数据，如果通道为空，表示锁已经被其他协程持有
	case <-m.ch:
		return true
	default:

	}
	return false
}

func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	// 创建一个倒计时
	timer := time.NewTimer(timeout)

	select {
	// 接受到值，说明并未加锁
	case <-m.ch:
		timer.Stop()
		return true
	// 等待定时器触发
	case <-timer.C:

	}
	// 获取锁超时
	return false
}

func (m *Mutex) IsLocked() bool {
	// != 0 时说明有缓冲元素，并未加锁
	return len(m.ch) == 0
}

// 使用channel实现一个互斥锁
func main() {
	m := NewMutex()
	log.Println("just get Lock, IsLocked: ", m.IsLocked())
	m.Lock()
	log.Println("try to lock, IsLocked: ", m.IsLocked())
	ok := m.TryLock()
	log.Printf("try to lock again, IsLocked %v\n", ok)

	m.Unlock()
	log.Println("try to unlock, IsLocked: ", m.IsLocked())
	ok = m.TryLock()
	log.Printf("try to lock again, IsLocked %v\n", ok)

	go func() {
		time.Sleep(5 * time.Second)
		m.Unlock()
	}()

	ok = m.LockTimeout(10 * time.Second)
	log.Printf("lockTimeout-3 %v\n", ok)

}
