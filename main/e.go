package main

import "fmt"

// 嵌入类型

type userss struct {
	name string
	email string
}

func (u *userss)notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}

type admin struct {
	userss	// 嵌入类型
	level string
}
func main() {
	ad := admin{
		userss:userss{name:"john smith",email:"john@yahoo.com"},level:"super",
	}

	// 直接访问内部类的方法
	ad.userss.notify()

	// 内部类的方法也被提升到外部类
	ad.notify()
}
