package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 的数量
	taskLoad         = 10 // 要处理的工作数量
)

var nWg sync.WaitGroup

// init 初始化包，Go 语言运行时会在其他代码执行之前
func init() {
	rand.Seed(time.Now().Unix())
}

// 有缓存的通道 buffered channel
func main() {
	// 创建一个有缓存的通道
	tasks := make(chan string, taskLoad)
	// 启动 goroutine 来处理工作
	nWg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完时关闭通道，以便所有 goroutine 退出
	close(tasks)
	nWg.Wait()
}

func worker(tasks chan string, worker int) {
	defer nWg.Done()
	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
