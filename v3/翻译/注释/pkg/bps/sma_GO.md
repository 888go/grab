
<原文开始>
// NewSMA returns a gauge that uses a Simple Moving Average with the given
// number of samples to measure the bytes per second of a byte stream.
//
// BPS is computed using the timestamp of the most recent and oldest sample in
// the sample buffer. When a new sample is added, the oldest sample is dropped
// if the sample count exceeds maxSamples.
//
// The gauge does not account for any latency in arrival time of new samples or
// the desired window size. Any variance in the arrival of samples will result
// in a BPS measurement that is correct for the submitted samples, but over a
// varying time window.
//
// maxSamples should be equal to 1 + (window size / sampling interval) where
// window size is the number of seconds over which the moving average is
// smoothed and sampling interval is the number of seconds between each sample.
//
// For example, if you want a five second window, sampling once per second,
// maxSamples should be 1 + 5/1 = 6.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
	// prevent integer overflow in sampleCount. Values greater or equal to
	// maxSamples have the same semantic meaning.
<原文结束>

# <翻译开始>
	// 防止sampleCount发生整数溢出。值大于或等于maxSamples具有相同的语义含义。
	// md5:c2f0ebfb0b578a39
# <翻译结束>


<原文开始>
// we need two samples to start
<原文结束>

# <翻译开始>
// 我们需要两个样本才能开始 md5:ef64437f5a7ce54b
# <翻译结束>


<原文开始>
// First sample is always the oldest until ring buffer first overflows
<原文结束>

# <翻译开始>
// 第一个样本始终是最旧的，直到环形缓冲区首次溢出 md5:84d3ea67895764fe
# <翻译结束>

