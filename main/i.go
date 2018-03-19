package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 展示 goroutine 调度器在多个逻辑处理器上的行为
func main() {
	// 分配 2 个逻辑处理器给调度器
	runtime.GOMAXPROCS(2)
	// wg 用来等待程序完成
	var wg sync.WaitGroup
	// 添加 2 个 goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// 运行一个 goroutine 匿名函数
	go func() {
		// 函数退出通知 wg
		defer wg.Done()
		// 打印 3 次字母表
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	// 运行一个 goroutine 匿名函数
	go func() {
		// 函数退出通知 wg
		defer wg.Done()
		// 打印 3 次字母表
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Wating To Finish")
	// 等待结束
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
