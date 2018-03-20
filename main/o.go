package main

import "log"

// log 包
func init() {
	log.SetPrefix("CID: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println() 之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 在调用 Println() 之后会接着嗲用 panic()
	log.Panicln("panic message")
}

