package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counters int
	lWg     sync.WaitGroup

	// mutex 用来定义一段代码临界值
	mutex sync.Mutex
)
// mutex 互斥锁
func main() {
	lWg.Add(2)

	go incCounterL(1)
	go incCounterL(2)

	lWg.Wait()
	fmt.Printf("Final Counter: %d\n", counters)
}

func incCounterL(id int) {
	defer lWg.Done()
	for count := 0; count < 2; count++ {
		// 锁上临界区，只允许一个 进入
		mutex.Lock()
		// 临界区
		{
			// 捕获 counter 的值
			value := counters
			// 当前 goroutine 从线程退出，并放回到列队
			runtime.Gosched()
			// 增加本地 value 变量的值
			value++
			// 将该值保存回 counter
			counters = value
		}
		// 释放锁，允许其他正在等待的 goroutine 进入
		mutex.Unlock()
	}
}
