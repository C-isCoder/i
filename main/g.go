package main

import (
	"runtime"
	"sync"
	"fmt"
)

// 如何创建 goroutine 和 调度器行为
func main() {
	// 分配 1 个逻辑处理器给调度器。
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	var wg sync.WaitGroup

	// 添加 2 个计数器，代表 2 个 goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 匿名函数创建一个 goroutine
	go func() {
		// defer 函数退出时候 总会执行。 通知 wg 函数运行结束
		defer wg.Done()

		// 打印字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// 打印 3 次大写字母表
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Wating To finish")
	// 等待 goroutine 执行结束
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
