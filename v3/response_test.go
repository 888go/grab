package 下载类

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/cavaliergopher/grab/v3/pkg/grabtest"
)

// testComplete 验证已完成的 Response 是否包含所有期望的字段。
func testComplete(t *testing.T, resp *X响应) {
	<-resp.Done
	if !resp.X是否已完成() {
		t.Errorf("Response.IsComplete returned false")
	}

	if resp.X传输开始时间.IsZero() {
		t.Errorf("Response.Start is zero")
	}

	if resp.X传输完成时间.IsZero() {
		t.Error("Response.End is zero")
	}

	if eta := resp.X取估计完成时间(); eta != resp.X传输完成时间 {
		t.Errorf("Response.ETA is not equal to Response.End: %v", eta)
	}

// 以下字段仅在未发生错误时才应设置
	if resp.X等待错误() == nil {
		if resp.X文件名 == "" {
			t.Errorf("Response.Filename is empty")
		}

		if resp.X取总字节() == 0 {
			t.Error("Response.Size is zero")
		}

		if p := resp.X取进度(); p != 1.00 {
			t.Errorf("Response.Progress returned %v (%v/%v bytes), expected 1", p, resp.X已完成字节(), resp.X取总字节())
		}
	}
}

// TestResponseProgress 测试正在进行的文件传输进度指示功能。
func TestResponseProgress(t *testing.T) {
	filename := ".testResponseProgress"
	defer os.Remove(filename)

	sleep := 300 * time.Millisecond
	size := 1024 * 8 // bytes

	grabtest.WithTestServer(t, func(url string) {
// 请求一个慢速传输
		req := mustNewRequest(filename, url)
		resp := X默认全局客户端.X下载(req)

// 确保转账尚未开始
		if resp.X是否已完成() {
			t.Errorf("Transfer should not have started")
		}

		if p := resp.X取进度(); p != 0 {
			t.Errorf("Transfer should not have started yet but progress is %v", p)
		}

// 等待传输完成
		<-resp.Done

// 确保传输已完整完成
		if p := resp.X取进度(); p != 1 {
			t.Errorf("Transfer is complete but progress is %v", p)
		}

		if s := resp.X已完成字节(); s != int64(size) {
			t.Errorf("Expected to transfer %v bytes, got %v", size, s)
		}
	},
		grabtest.TimeToFirstByte(sleep),
		grabtest.ContentLength(size),
	)
}

func TestResponseOpen(t *testing.T) {
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest("", url+"/someFilename"))
		f, err := resp.X等待完成后打开文件()
		if err != nil {
			t.Error(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				t.Error(err)
			}
		}()
		grabtest.AssertSHA256Sum(t, grabtest.DefaultHandlerSHA256ChecksumBytes, f)
	})
}

func TestResponseBytes(t *testing.T) {
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest("", url+"/someFilename"))
		b, err := resp.X等待完成后取字节集()
		if err != nil {
			t.Error(err)
			return
		}
		grabtest.AssertSHA256Sum(
			t,
			grabtest.DefaultHandlerSHA256ChecksumBytes,
			bytes.NewReader(b),
		)
	})
}
