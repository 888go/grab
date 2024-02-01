package 下载类

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/cavaliergopher/grab/v3/pkg/bps"
)

type transfer struct {
	n     int64 // 必须在386架构上进行64位对齐
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
		gauge: bps.NewSMA(6), // 每隔一秒进行一次采样，计算五秒内的移动平均值
		lim:   lim,
		w:     dst,
		r:     src,
		b:     buf,
	}
}

// copy 函数的行为类似于 io.CopyBuffer，但除此之外，它还会检查给定的 context.Context 是否已取消，
// 以线程安全的方式报告进度，并跟踪传输速率。
func (c *transfer) copy() (written int64, err error) {
// 在另一个goroutine中维护一个bps计数器
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	go bps.Watch(ctx, c.gauge, c.N, time.Second)

// 开始传输
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
// 等待速率限制器
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

// N 返回已传输的字节数。
func (c *transfer) N() (n int64) {
	if c == nil {
		return 0
	}
	n = atomic.LoadInt64(&c.n)
	return
}

// BPS 返回当前的字节每秒传输速率，采用简单移动平均算法计算。
func (c *transfer) BPS() (bps float64) {
	if c == nil || c.gauge == nil {
		return 0
	}
	return c.gauge.BPS()
}
