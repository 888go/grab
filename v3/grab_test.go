package 下载类

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/cavaliergopher/grab/v3/pkg/grabtest"
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
// 切换到临时目录，以便下载到当前工作目录的测试文件能够被隔离并清理掉
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

// TestGet 测试 grab.Get 功能
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
// 下载文件到 /tmp 目录
	resp, err := X下载("/tmp", "http://example.com/example.zip")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.X文件名)
}

func mustNewRequest(dst, urlStr string) *X下载参数 {
	req, err := X生成下载参数(dst, urlStr)
	if err != nil {
		panic(err)
	}
	return req
}

func mustDo(req *X下载参数) *X响应 {
	resp := X默认全局客户端.X下载(req)
	if err := resp.X等待错误(); err != nil {
		panic(err)
	}
	return resp
}
