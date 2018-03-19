package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mWg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 无缓冲 unbuffered channel 模拟网球比赛
func main() {
	// 创建无缓冲通道
	court := make(chan int)

	mWg.Add(2)

	// 启动两个选手
	go player("Tom", court)
	go player("iCong", court)

	// 发球 channel 赋值
	court <- 1
	// 阻塞 等待结束
	mWg.Wait()
}

func player(name string, court chan int) {
	defer mWg.Done()

	for {
		// 等待球被击打过来  <- court 从 court 通道取值
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道 输了
			close(court)
			return
		}
		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball ++
		// 把值放到管道里
		court <- ball
	}
}
