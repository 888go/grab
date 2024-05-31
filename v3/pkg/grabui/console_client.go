package grabui

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/888go/grab/v3"
)

type ConsoleClient struct {
	mu                            sync.Mutex
	client                        *下载类.Client
	succeeded, failed, inProgress int
	responses                     []*下载类.Response
}

func NewConsoleClient(client *下载类.Client) *ConsoleClient {
	return &ConsoleClient{
		client: client,
	}
}

func (c *ConsoleClient) Do(
	ctx context.Context,
	workers int,
	reqs ...*下载类.Request,
) <-chan *下载类.Response {
	// 缓冲区大小可以防止接收者处理速度慢导致的后压（back pressure）问题。 md5:5cad013415a13329
	pump := make(chan *下载类.Response, len(reqs))

	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.failed = 0
		c.inProgress = 0
		c.succeeded = 0
		c.responses = make([]*下载类.Response, 0, len(reqs))
		if c.client == nil {
			c.client = 下载类.DefaultClient
		}

		fmt.Printf("Downloading %d files...\n", len(reqs))
		respch := c.client.X多线程下载(workers, reqs...)
		t := time.NewTicker(200 * time.Millisecond)
		defer t.Stop()

	Loop:
		for {
			select {
			case <-ctx.Done():
				break Loop

			case resp := <-respch:
				if resp != nil {
					// 已经接收到一个新的响应并开始下载 md5:15f5f005f545cc95
					c.responses = append(c.responses, resp)
					pump <- resp // send to caller
				} else {
					// channel已关闭 - 所有下载已完成 md5:7f2447cabb28251e
					break Loop
				}

			case <-t.C:
				// 在时钟Tick时更新UI md5:a52f30bc4b10a415
				c.refresh()
			}
		}

		c.refresh()
		close(pump)

		fmt.Printf(
			"Finished %d successful, %d failed, %d incomplete.\n",
			c.succeeded,
			c.failed,
			c.inProgress)
	}()
	return pump
}

// refresh 会将所有下载的进度输出到终端 md5:b45b3af8576f0c8e
func (c *ConsoleClient) refresh() {
	// 清除不完整下载的行 md5:6f2c2c6ad26e77ea
	if c.inProgress > 0 {
		fmt.Printf("\033[%dA\033[K", c.inProgress)
	}

	// 打印新完成的下载 md5:7ffedfed87948493
	for i, resp := range c.responses {
		if resp != nil && resp.X是否已完成() {
			if resp.X等待错误() != nil {
				c.failed++
				fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n",
					resp.X下载参数.X取下载链接(),
					resp.X等待错误())
			} else {
				c.succeeded++
				fmt.Printf("Finished %s %s / %s (%d%%)\n",
					resp.X文件名,
					byteString(resp.X已完成字节()),
					byteString(resp.X取总字节()),
					int(100*resp.X取进度()))
			}
			c.responses[i] = nil
		}
	}

	// /* 打印不完整下载的进度 */ md5:4b96256fbbba36a5
	c.inProgress = 0
	for _, resp := range c.responses {
		if resp != nil {
			fmt.Printf("Downloading %s %s / %s (%d%%) - %s ETA: %s \033[K\n",
				resp.X文件名,
				byteString(resp.X已完成字节()),
				byteString(resp.X取总字节()),
				int(100*resp.X取进度()),
				bpsString(resp.X取每秒字节()),
				etaString(resp.X取估计完成时间()))
			c.inProgress++
		}
	}
}

func bpsString(n float64) string {
	if n < 1e3 {
		return fmt.Sprintf("%.02fBps", n)
	}
	if n < 1e6 {
		return fmt.Sprintf("%.02fKB/s", n/1e3)
	}
	if n < 1e9 {
		return fmt.Sprintf("%.02fMB/s", n/1e6)
	}
	return fmt.Sprintf("%.02fGB/s", n/1e9)
}

func byteString(n int64) string {
	if n < 1<<10 {
		return fmt.Sprintf("%dB", n)
	}
	if n < 1<<20 {
		return fmt.Sprintf("%dKB", n>>10)
	}
	if n < 1<<30 {
		return fmt.Sprintf("%dMB", n>>20)
	}
	if n < 1<<40 {
		return fmt.Sprintf("%dGB", n>>30)
	}
	return fmt.Sprintf("%dTB", n>>40)
}

func etaString(eta time.Time) string {
	d := eta.Sub(time.Now())
	if d < time.Second {
		return "<1s"
	}
	// 截断到1秒分辨率 md5:d627107a4a28696f
	d /= time.Second
	d *= time.Second
	return d.String()
}
