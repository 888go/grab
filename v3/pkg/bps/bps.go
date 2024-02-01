/*
Package bps provides gauges for calculating the Bytes Per Second transfer rate
of data streams.
*/
package bps

import (
	"context"
	"time"
)

// Gauge 是本包中所有 BPS 计量器的通用接口。给定一段时间内的样本集，每种计量器类型都可以用来测量数据流的每秒字节数（Bytes Per Second，BPS）传输速率。
//
// 所有样本的时间戳和值必须单调递增。每个样本应代表数据流中已发送的总字节数，而不是自上次采样以来发送的字节数。
//
// 为了确保计量器能够尽快报告进度，在数据流开始时立即获取初始样本。
//
// 所有计量器实现都支持安全的并发使用。
type Gauge interface {
// Sample 添加对被监控流进度的新采样。
	Sample(t time.Time, n int64)

// BPS 返回监控流计算出的每秒字节数（Bytes Per Second）速率。
	BPS() float64
}

// SampleFunc 由 Watch 调用，用于周期性地对被监控的流进行采样。
type SampleFunc func() (n int64)

// Watch 会周期性调用给定的 SampleFunc 来采样被监控流的进度，并更新给定的仪表（gauge）。SampleFunc 应该返回流自从开始以来传输的总字节数。
//
// Watch 是一个阻塞调用，通常应在新的 goroutine 中调用。为了避免 goroutine 泄漏，在流完成或被取消时，请确保取消给定的上下文。
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
