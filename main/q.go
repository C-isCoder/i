package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// json 对应的结构体
type (
	gResult struct {
		// `` 标签，对应 json 的字段，如果不写则默认 忽略大小写直接用字段名进行匹配（匹配不到 返回 零 值）。
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

// 解析 json 数据
func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println(gr)
}
