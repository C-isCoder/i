package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// map 转换成 JSON
func main() {
	// 创建一个实现空接口的 map json 包方法接受实现空接口的任意类型
	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	c := make(map[string]interface{})
	// key	----  value
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}
	// 将 map 序列化成 JSON 字符串MarshalIndent 带缩进的 @prefix 前缀 @indent 缩进
	/**
		{
    		"contact": {
        		"cell": "415.555.5555",
        		"home": "415.333.3333"
    		},
    		"name": "Gopher",
    		"title": "programmer"
		}
	 */
	data, err := json.MarshalIndent(c, "", "    ")
	// 不带缩进的适合 web api
	// {"contact":{"cell":"415.555.5555","home":"415.333.3333"},"name":"Gopher","title":"programmer"}
	//data, err := json.Marshal(c)
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(data)
}
