package test

import (
	"net/http"
	"testing"
)

const uCheckMark = "🙆️"
const uBallotX = "🙅"

// 表组测试
// 测试文件命名 xx_test.go
// 测试函数 大写 Test 开头，函数签名必须接收 testing.T 类型的指针，且函数不能有返回值
// 遵循上面约定 go 测试工具才会执行。
func TestDownload_U(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK, // 200
		},
		{
			"http://www.google.com",
			http.StatusOK,
		},
	}
	t.Log("Given the need to test downloading content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", uBallotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", uCheckMark)
				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, uCheckMark, resp.StatusCode)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, uBallotX, resp.StatusCode)
				}
			}
		}

	}
}
