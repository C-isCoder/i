package handlers

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"log"
	"fmt"
)

// go DOC 生成示例（编写基础示例）
// 代码示例名字 Example 开头方法不传值
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println(u)
	// Output:
	// {Bill bill@mail.com}
}
