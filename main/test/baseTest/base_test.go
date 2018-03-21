package baseTest

import (
	"fmt"
	"testing"
	"strconv"
)

/*
 	基准性能测试 3 种转换字符串的方法 找出最快的
	命令 -bench 可以制定函数单独测试一个 "." 表示所有函数
	-benchtime 设置循环运行的时长
	-benchmem  提供显示每次操作分配内存的次数
	ns/op 		每次执行耗时 纳米/次
	B/op  		每次执行分配的内存字节数
	allocs/op	每次执行从堆上分配内存的次数
	go test -v -run="none" -bench=. -benchtime="3s" -benchmem
	////////////////////////////////////////////////////////////////////////////////////////////////////
	»     go test -v -run="none" -bench=. -benchtime="3s" -benchmem
	goos: darwin
	goarch: amd64
	pkg: i/main/test/baseTest
	BenchmarkSprintf-8      50000000                87.7 ns/op            16 B/op          2 allocs/op
	BenchmarkFormat-8       2000000000               3.28 ns/op            0 B/op          0 allocs/op
	BenchmarkItoa-8         1000000000               4.96 ns/op            0 B/op          0 allocs/op
	PASS
	ok      i/main/test/baseTest    16.802s
	////////////////////////////////////////////////////////////////////////////////////////////////////
 */

// 对 fmt.Sprintf 函数进行基准测试
// 基准测试函数必须以 Benchmark 开头，
// 接受一个只想 testing.B 类型的指针作为唯一参数。
func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// 对 strconv.FormatInt 函数测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// 对 strconv.Itoa 函数测试
func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
