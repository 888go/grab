package bps

import (
	"sync"
	"time"
)

// NewSMA 返回一个使用简单移动平均法（Simple Moving Average）的仪表，该方法使用给定的样本数量来测量字节流的每秒字节数（BPS）。
//
// BPS 的计算基于样本缓冲区中最新和最旧样本的时间戳。当添加新样本时，如果样本数量超过 maxSamples，最旧的样本将被丢弃。
//
// 该仪表不考虑新样本到达时间的延迟或期望的窗口大小。样本到达的任何偏差都将会导致 BPS 测量值正确地适用于提交的样本，但其覆盖的时间窗口会有所变化。
//
// maxSamples 应等于 1 + (窗口大小 / 采样间隔)，其中窗口大小是平滑移动平均的时间范围（以秒为单位），采样间隔是每个样本之间的时间间隔（以秒为单位）。
//
// 例如，如果你想要一个五秒的窗口，并且每秒采样一次，maxSamples 应该是 1 + 5/1 = 6。
// md5:62f6153626d33f59
func NewSMA(maxSamples int) Gauge {
	if maxSamples < 2 {
		panic("sample count must be greater than 1")
	}
	return &sma{
		maxSamples: uint64(maxSamples),
		samples:    make([]int64, maxSamples),
		timestamps: make([]time.Time, maxSamples),
	}
}

type sma struct {
	mu          sync.Mutex
	index       uint64
	maxSamples  uint64
	sampleCount uint64
	samples     []int64
	timestamps  []time.Time
}

func (c *sma) Sample(t time.Time, n int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamps[c.index] = t
	c.samples[c.index] = n
	c.index = (c.index + 1) % c.maxSamples

// 防止sampleCount发生整数溢出。值大于或等于maxSamples具有相同的语义含义。
// md5:c2f0ebfb0b578a39
	c.sampleCount++
	if c.sampleCount > c.maxSamples {
		c.sampleCount = c.maxSamples
	}
}

func (c *sma) BPS() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 我们需要两个样本才能开始 md5:ef64437f5a7ce54b
	if c.sampleCount < 2 {
		return 0
	}

	// 第一个样本始终是最旧的，直到环形缓冲区首次溢出 md5:84d3ea67895764fe
	oldest := c.index
	if c.sampleCount < c.maxSamples {
		oldest = 0
	}

	newest := (c.index + c.maxSamples - 1) % c.maxSamples
	seconds := c.timestamps[newest].Sub(c.timestamps[oldest]).Seconds()
	bytes := float64(c.samples[newest] - c.samples[oldest])
	return bytes / seconds
}
