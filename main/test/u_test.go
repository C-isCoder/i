package test

import (
	"net/http"
	"testing"
)

//const checkMark = "\u2713" // 对号 ✅
const checkMark = "🙆️"
//const ballotX = "\u2717" // 错 ❌
const ballotX = "🙅"

// 基础单元测试
// 测试文件命名 xx_test.go
// 测试函数 大写 Test 开头，函数签名必须接收 testing.T 类型的指针，且函数不能有返回值
// 遵循上面约定 go 测试工具才会执行。
func TestDownload(t *testing.T) {
	//url := "http://www.goinggo.net/feeds/posts/defautl?alt=rss" //404
	url := "http://www.baidu.com"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v %v", statusCode, checkMark, resp.StatusCode)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
