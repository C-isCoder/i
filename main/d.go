package main

import "fmt"

// 多态

// 接口
type notifer interface {
	notify()
}

type users struct {
	name string
	email string
}

func (u *users) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}

type amdin struct {
	name string
	email string
	}

func (a *amdin)notify()  {
	fmt.Printf("Sending admin email to %s<%s>n",a.name,a.email)
}

func sendNotification(n notifer)  {
	n.notify()
}

func main() {
	bill := users{"Bill","bill@email.com"}
	sendNotification(&bill)

	lisa := amdin{"lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
