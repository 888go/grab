package 下载类

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/888go/grab/v3/pkg/grabtest"
)

// testRateLimiter是一个简单的速率限制器，每秒限制r个令牌的通过量。已发出的令牌总数由n跟踪。
// md5:d54804ac9a0ceb1a
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
// 使用简单的512bps限制，一次下载128字节的文件，每次8字节。应该花费超过250毫秒。
// md5:e92aa1bfdda7be8a
	filesize := 128
	filename := ".testRateLimiter"
	defer os.Remove(filename)

	grabtest.WithTestServer(t, func(url string) {
		// limit to 512bps
		lim := &testRateLimiter{r: 512}
		req := mustNewRequest(filename, url)

		// 确保每次下载8字节，以实现对速率限制器的多次访问 md5:d541457cb761ceaf
		req.X缓冲区大小 = 8
		req.X速率限制器 = lim

		resp := mustDo(req)
		testComplete(t, resp)
		if lim.n != filesize {
			t.Errorf("expected %d bytes to pass through limiter, got %d", filesize, lim.n)
		}
		if resp.X取下载已持续时间().Seconds() < 0.25 {
			// BUG：如果由于无关原因（transfer速度慢）导致测试通过，这个注释表示这是一个已知问题。 md5:a7882bdf6de8b17f
			t.Errorf("expected transfer to take >250ms, took %v", resp.X取下载已持续时间())
		}
	}, grabtest.ContentLength(filesize))
}

func ExampleRateLimiter() {
	req, _ := X生成下载参数("", "http://www.golang-book.com/public/pdf/gobook.pdf")

// 附加一个1Mbps速率限制器，类似于golang.org/x/time/rate包中的令牌桶实现。
// md5:a33151d1a418a247
	req.X速率限制器 = NewLimiter(1048576)

	resp := DefaultClient.X下载(req)
	if err := resp.X等待错误(); err != nil {
		log.Fatal(err)
	}
}
