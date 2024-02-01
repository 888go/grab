package bps

import (
	"sync"
	"time"
)

// NewSMA 返回一个使用给定样本数量计算简单移动平均值的仪表，
// 用于测量字节流的每秒字节数（BPS）。
// BPS 的计算基于样本缓冲区中最新和最旧样本的时间戳。当添加新样本时，
// 如果样本数量超过 maxSamples，则丢弃最旧的样本。
// 该仪表并未考虑新样本到达时间的延迟或期望窗口大小。样本到达时间的任何变化
// 都会导致 BPS 测量值在提交的样本上是正确的，但其覆盖的时间窗口会有所变动。
// maxSamples 应等于 1 + (窗口大小 / 采样间隔)，其中窗口大小是指移动平均值平滑处理所跨越的秒数，
// 而采样间隔则是每个样本之间的秒数。
// 例如，如果你想设置一个5秒的窗口，并且每秒采样一次，那么 maxSamples 应为 1 + 5/1 = 6。
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

// 防止在sampleCount中发生整数溢出。当值大于或等于
// maxSamples时，它们具有相同的语义含义。
	c.sampleCount++
	if c.sampleCount > c.maxSamples {
		c.sampleCount = c.maxSamples
	}
}

func (c *sma) BPS() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

// 我们需要两个样本开始
	if c.sampleCount < 2 {
		return 0
	}

// 第一个样本始终是最旧的，直到环形缓冲区首次溢出为止
	oldest := c.index
	if c.sampleCount < c.maxSamples {
		oldest = 0
	}

	newest := (c.index + c.maxSamples - 1) % c.maxSamples
	seconds := c.timestamps[newest].Sub(c.timestamps[oldest]).Seconds()
	bytes := float64(c.samples[newest] - c.samples[oldest])
	return bytes / seconds
}
