package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact 结构代表我们的 json 字符串
type Contact struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// json 字符穿
var JSON = `{"name" : "Gopher","title" : "programmer","contact" : {"home" : "415.333.3333","cell" : "415.555.5555"}}`

// 解码 json 字符串
func main() {
	// 在 json 串可以对应到词一个结构类型
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(c)
	// 当 json 串无法对应结构类型的时候可以解码到时 map 中
	var cm map[string]interface{} // 要实现空接口才能 Decode ，decode 要传入一个任意实现空接口的类型
	err1 := json.Unmarshal([]byte(JSON), &cm)
	if err1 != nil {
		log.Println("Error: ", err1)
	}
	fmt.Println("Name: ", cm["name"])
	fmt.Println("Title: ", cm["title"])
	fmt.Println("Contact: ")
	fmt.Println("H:", cm["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", cm["contact"].(map[string]interface{})["cell"])
}
