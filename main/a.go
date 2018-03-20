package main

import (
	"net/http"
	"os"
	"log"
	"io"
)

// 简单的 网络请求 （io.Writer io.Reader） 的基本用法
func main() {
	// Response 的 Body 是 io.Reader 的实现
	r, e := http.Get(os.Args[1])
	if e != nil {
		log.Fatalln(e)
	}

	// 创建文件来保存响应内容
	f, e := os.Create(os.Args[2])
	if e != nil {
		log.Fatalln(e)
	}
	defer f.Close()

	// 使用 MultiWriter 这样可以同时向文件和标准输出设备进行写入操作
	dest := io.MultiWriter(os.Stdout, f)

	// 读出响应的内容，并写入两个目的地
	io.Copy(dest, r.Body)
	if e := r.Body.Close(); e != nil {
		log.Println(e)
	}
}
