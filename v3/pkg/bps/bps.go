/*
Package bps provides gauges for calculating the Bytes Per Second transfer rate
of data streams.
*/
package bps

import (
	"context"
	"time"
)

// Gauge是该包中所有BPS度量的通用接口。给定一段时间内的多个采样，每种度量类型都可以用来测量数据流的字节/秒传输速率。
// 
// 所有采样点的时间戳和值都必须单调递增。每个采样点应表示流中的总字节数，而不是自上一个采样点以来发送的字节数。
// 
// 为了确保度量能够尽快报告进度，请在数据流开始时立即获取初始采样。
// 
// 所有度量实现都支持并发使用。
// md5:96b611c42ceb040c
type Gauge interface {
	// Sample 添加监控流的新进度样本。 md5:d2b29305c97808d1
	Sample(t time.Time, n int64)

	// BPS 返回监控流的计算出的每秒字节数速率。 md5:6ca06acbcf94e710
	BPS() float64
}

// SampleFunc被Watch用来定期从监控流中采样。 md5:4f93786f07892dce
type SampleFunc func() (n int64)

// Watch 会定期调用给定的 SampleFunc 来采样监控流的进度，并更新给定的指标。SampleFunc 应返回自流开始以来已传输的总字节数。
// 
// Watch 是一个阻塞调用，通常应在新的goroutine中进行。为了防止goroutine泄漏，请确保在流完成或取消时取消给定的上下文。
// md5:866a6d9ed56020b3
func Watch(ctx context.Context, g Gauge, f SampleFunc, interval time.Duration) {
	g.Sample(time.Now(), f())
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-t.C:
			g.Sample(now, f())
		}
	}
}
