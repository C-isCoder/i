package main

import (
	"fmt"
	"i/main/f"
)

func main() {
	// alertCounter 是未公开名字，不能直接调用
	//alert := a.alertCounter(10)
	// 通过工厂方法 访问未公开的实例
	alert := f.New(10)
	fmt.Printf("Counter: %d\n", alert)


	// 公开结构类型中未公开字段
	u := f.User{
		Name: "Bill",
		//email: "bill@email.com"  //email 未公开 无法访问
	}
	fmt.Printf("User: %v\n", u)


	// 公开结构体中 嵌入未公开字段
	ad := f.Admin{
		Rights: 10,
	}
	// 设置未公开的内部类的公开字段
	ad.Name = "Bill"
	ad.Email = "bill@email.com"
	fmt.Printf("User: %v\n", ad)

}
