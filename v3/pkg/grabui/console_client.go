package grabui

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

type ConsoleClient struct {
	mu                            sync.Mutex
	client                        *下载类.Client
	succeeded, failed, inProgress int
	responses                     []*下载类.X响应
}

func NewConsoleClient(client *下载类.Client) *ConsoleClient {
	return &ConsoleClient{
		client: client,
	}
}

func (c *ConsoleClient) Do(
	ctx context.Context,
	workers int,
	reqs ...*下载类.X下载参数,
) <-chan *下载类.X响应 {
// 缓冲区大小防止慢速接收者造成回压
	pump := make(chan *下载类.X响应, len(reqs))

	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.failed = 0
		c.inProgress = 0
		c.succeeded = 0
		c.responses = make([]*下载类.X响应, 0, len(reqs))
		if c.client == nil {
			c.client = 下载类.X默认全局客户端
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
// 新的响应已收到并已开始下载
					c.responses = append(c.responses, resp)
					pump <- resp // send to caller
				} else {
// 通道已关闭 - 所有下载已完成
					break Loop
				}

			case <-t.C:
// 在时钟滴答时更新UI
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

// refresh 将所有下载进度打印到终端
func (c *ConsoleClient) refresh() {
// 清除不完整下载的行
	if c.inProgress > 0 {
		fmt.Printf("\033[%dA\033[K", c.inProgress)
	}

// 打印新完成的下载内容
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

// 打印未完成下载的进度
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
// 截断至1秒分辨率
	d /= time.Second
	d *= time.Second
	return d.String()
}
