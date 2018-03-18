package main

import "fmt"

// 多态

// 接口
type notifer interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify()  {
	fmt.Println("Sending user email to %s<%s>\n",u.name,u.email)
}

type amdin struct {
	name string
	email string
	}

func (a *amdin)notify()  {
	fmt.Println("Sending admin email to %s<%s>n",a.name,a.email)
}

func sendNotification(n notifer)  {
	n.notify()
}

func main() {
	bill := user{"Bill","bill@email.com"}
	sendNotification(&bill)

	lisa := amdin{"lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
