package i

import "fmt"
/**
 *	接口
 */
type notifier interface {
	notify()
}

// 定义结构体
type user struct {
	name  string
	email string
}

// user 指针接收者实现接口
func (u *user) notify() {
	fmt.Println("Sending user email to %s<%s>\n", u.name, u.email)
}

// 方法接受实现了 notifier 接口的值
func sendNotify(u notifier) {
	u.notify()
}
func main() {
	u := user{"tom", "tom@gmail.com"}
	// user 是指针接收实现的接口 u 是一个类型值 编译不通过
	//sendNotify(u)

	// ok 传指针（地址）
	sendNotify(&u)
	// notify method has pointer receiver
}
