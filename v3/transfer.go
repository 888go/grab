package 下载类

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/888go/grab/v3/pkg/bps"
)

type transfer struct {
	n     int64 // 在386架构上必须是64位对齐 md5:abe49438c68a8c64
	ctx   context.Context
	gauge bps.Gauge
	lim   RateLimiter
	w     io.Writer
	r     io.Reader
	b     []byte
}

func newTransfer(ctx context.Context, lim RateLimiter, dst io.Writer, src io.Reader, buf []byte) *transfer {
	return &transfer{
		ctx:   ctx,
		gauge: bps.NewSMA(6), // 每秒采样五秒滑动平均值 md5:5bcc165c9abe93fa
		lim:   lim,
		w:     dst,
		r:     src,
		b:     buf,
	}
}

// copy 的行为类似于 io.CopyBuffer，但它会检查给定的 context.Context 是否被取消，以线程安全的方式报告进度，并跟踪传输速率。
// md5:41ad622882e2be6a
func (c *transfer) copy() (written int64, err error) {
	// 在另一个goroutine中维护一个bps（每秒字节）仪表盘 md5:e4bdbe611923ed71
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	go bps.Watch(ctx, c.gauge, c.N, time.Second)

	// start the transfer
	if c.b == nil {
		c.b = make([]byte, 32*1024)
	}
	for {
		select {
		case <-c.ctx.Done():
			err = c.ctx.Err()
			return
		default:
			// keep working
		}
		nr, er := c.r.Read(c.b)
		if nr > 0 {
			nw, ew := c.w.Write(c.b[0:nr])
			if nw > 0 {
				written += int64(nw)
				atomic.StoreInt64(&c.n, written)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
			// wait for rate limiter
			if c.lim != nil {
				err = c.lim.WaitN(c.ctx, nr)
				if err != nil {
					return
				}
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

// N返回传输的字节数。 md5:9d2049135c9567cb
func (c *transfer) N() (n int64) {
	if c == nil {
		return 0
	}
	n = atomic.LoadInt64(&c.n)
	return
}

// BPS 返回当前的每秒字节数传输速率，使用简单的移动平均方法。
// md5:f09c67b9534a83a1
func (c *transfer) BPS() (bps float64) {
	if c == nil || c.gauge == nil {
		return 0
	}
	return c.gauge.BPS()
}
