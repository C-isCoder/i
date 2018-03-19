package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 通知正在执行的 goroutine 停止工作的标志
	shutdown int64
	kWg      sync.WaitGroup
)
// store 和 load 函数进行安全访问
func main() {
	// 添加计数
	kWg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
}

func doWork(name string) {
	defer kWg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		// 当 shutdown == 1 停止
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
