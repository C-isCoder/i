package main

import (
	"fmt"
	"bytes"
	"io"
	"os"
)
// io.Writer 接口
func main() {
	// 创建一个 buffer
	var b bytes.Buffer
	// buffer 实现了 write 接口 写入 “hello”
	// func (b *Buffer) Write(p []byte) (n int, err error)
	b.Write([]byte("hello "))
	// 用 Fprint 拼接到 buffer 里 将 buffer 的地址作为 io.Writer 类型值传入
	// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	fmt.Fprint(&b,"World")

	io.Copy(os.Stdout,&b)
}
