package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"fmt"
)

var (
	// 所有 goroutine 要操作胡变脸
	counter int64

	// wg
	jWg sync.WaitGroup
)
// 原子函数 加锁 atomic
func main() {
	// 增加计数器，2 个 goroutine
	jWg.Add(2)

	go incCounter(1)
	go incCounter(2)

	jWg.Wait()
	// 显示最终的值
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int64) {
	defer jWg.Done()

	for count := 0; count < 2; count++ {
		// 安全胡对 counter +1
		atomic.AddInt64(&counter, id)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}
