package test

import (
	"net/http"
	"testing"
	"fmt"
	"net/http/httptest"
)

// xml 数据
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
</channel>
</rss>`

// mockServer() 模拟服务器请求
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

const wCheckMark = "🙆️"
const wBallotX = "🙅"

// mock 模仿调用 基础单元测试

// 测试文件命名 xx_test.go
// 测试函数 大写 Test 开头，函数签名必须接收 testing.T 类型的指针，且函数不能有返回值
// 遵循上面约定 go 测试工具才会执行。
func TestDownload_w(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", wBallotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", wCheckMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v %v", statusCode, wCheckMark, resp.StatusCode)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, wBallotX, resp.StatusCode)
			}
		}
	}
}
