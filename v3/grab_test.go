package 下载类

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/888go/grab/v3/pkg/grabtest"
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
		// 更改为临时目录，以便下载到当前工作目录的测试文件相互隔离并自动清理 md5:1f1f2dbd63bb5e74
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		tmpDir, err := ioutil.TempDir("", "grab-")
		if err != nil {
			panic(err)
		}
		if err := os.Chdir(tmpDir); err != nil {
			panic(err)
		}
		defer func() {
			os.Chdir(cwd)
			if err := os.RemoveAll(tmpDir); err != nil {
				panic(err)
			}
		}()
		return m.Run()
	}())
}

// TestGet tests grab.Get
func TestGet(t *testing.T) {
	filename := ".testGet"
	defer os.Remove(filename)
	grabtest.WithTestServer(t, func(url string) {
		resp, err := X下载(filename, url)
		if err != nil {
			t.Fatalf("error in Get(): %v", err)
		}
		testComplete(t, resp)
	})
}

func ExampleGet() {
	// 将一个文件下载到/tmp目录下 md5:b27e46dcadfbc10c
	resp, err := X下载("/tmp", "http://example.com/example.zip")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.X文件名)
}

func mustNewRequest(dst, urlStr string) *Request {
	req, err := X生成下载参数(dst, urlStr)
	if err != nil {
		panic(err)
	}
	return req
}

func mustDo(req *Request) *Response {
	resp := DefaultClient.X下载(req)
	if err := resp.X等待错误(); err != nil {
		panic(err)
	}
	return resp
}
