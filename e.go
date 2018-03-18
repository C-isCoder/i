package main

import "fmt"

// 嵌入类型

type user1 struct {
	name string
	email string
}

func (u *user1)notify()  {
	fmt.Println("Sending user email to %s<%s>\n",u.name,u.email)
}

type admin struct {
	user1	// 嵌入类型
	level string
}
func main() {
	ad := admin{
		user1:user1{name:"john smith",email:"john@yahoo.com"},level:"super",
	}

	// 直接访问内部类的方法
	ad.user1.notify()

	// 内部类的方法也被提升到外部类
	ad.notify()
}
