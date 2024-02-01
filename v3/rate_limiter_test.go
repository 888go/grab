package 下载类

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/cavaliergopher/grab/v3/pkg/grabtest"
)

// testRateLimiter 是一个简单的速率限制器，其功能是将吞吐量限制为每秒 r 个令牌。已发放令牌的总数通过变量 n 进行跟踪。
// 注释翻译如下：
// ```go
// testRateLimiter 是一个粗略的速率限制器，用于将处理速率限制为每秒r个令牌。已发放的令牌总数量通过变量n进行统计。
type testRateLimiter struct {
	r, n int
}

func NewLimiter(r int) RateLimiter {
	return &testRateLimiter{r: r}
}

func (c *testRateLimiter) WaitN(ctx context.Context, n int) (err error) {
	c.n += n
	time.Sleep(
		time.Duration(1.00 / float64(c.r) * float64(n) * float64(time.Second)))
	return
}

func TestRateLimiter(t *testing.T) {
// 下载一个128字节的文件，每次下载8字节，并采用一个简单的512bps限速器
// 预计耗时超过250毫秒
	filesize := 128
	filename := ".testRateLimiter"
	defer os.Remove(filename)

	grabtest.WithTestServer(t, func(url string) {
		// limit to 512bps
		lim := &testRateLimiter{r: 512}
		req := mustNewRequest(filename, url)

// 确保通过每次下载8字节的方式多次访问速率限制器
		req.X缓冲区大小 = 8
		req.X速率限制器 = lim

		resp := mustDo(req)
		testComplete(t, resp)
		if lim.n != filesize {
			t.Errorf("expected %d bytes to pass through limiter, got %d", filesize, lim.n)
		}
		if resp.X取下载已持续时间().Seconds() < 0.25 {
// BUG: 如果由于无关原因传输速度较慢，这个测试可能会通过
			t.Errorf("expected transfer to take >250ms, took %v", resp.X取下载已持续时间())
		}
	}, grabtest.ContentLength(filesize))
}

func ExampleRateLimiter() {
	req, _ := X生成下载参数("", "http://www.golang-book.com/public/pdf/gobook.pdf")

// 绑定一个1Mbps的限速器，类似于来自golang.org/x/time/rate包中的令牌桶实现。
	req.X速率限制器 = NewLimiter(1048576)

	resp := X默认全局客户端.X下载(req)
	if err := resp.X等待错误(); err != nil {
		log.Fatal(err)
	}
}
