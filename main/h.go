package main

import (
	"fmt"
	"runtime"
	"sync"
)

var hWg sync.WaitGroup
// 展示 goroutine 调度器在单个线程调度
func main() {
	// 分配 1 个逻辑处理器给调度器
	runtime.GOMAXPROCS(1)
	// 添加 2 个 Goroutine
	hWg.Add(2)

	fmt.Println("Create Goroutines")
	// 运行 2 个 goroutine
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Wating To Finish")
	// 等待 go goroutine 执行
	hWg.Wait()
	// end
	fmt.Println("Terminating Program")
}

// 打印 1-5000 内的素数
func printPrime(prefix string) {
	// 函数结束 通知 sync.WaitGroup
	defer hWg.Done()
next: // 标识？
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next // 跳出执行下一个
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
