
<原文开始>
// Gauge is the common interface for all BPS gauges in this package. Given a
// set of samples over time, each gauge type can be used to measure the Bytes
// Per Second transfer rate of a data stream.
//
// All samples must monotonically increase in timestamp and value. Each sample
// should represent the total number of bytes sent in a stream, rather than
// accounting for the number sent since the last sample.
//
// To ensure a gauge can report progress as quickly as possible, take an initial
// sample when your stream first starts.
//
// All gauge implementations are safe for concurrent use.
<原文结束>

# <翻译开始>
// Gauge 是本包中所有 BPS 计量器的通用接口。给定一段时间内的样本集，每种计量器类型都可以用来测量数据流的每秒字节数（Bytes Per Second，BPS）传输速率。
//
// 所有样本的时间戳和值必须单调递增。每个样本应代表数据流中已发送的总字节数，而不是自上次采样以来发送的字节数。
//
// 为了确保计量器能够尽快报告进度，在数据流开始时立即获取初始样本。
//
// 所有计量器实现都支持安全的并发使用。
# <翻译结束>


<原文开始>
	// Sample adds a new sample of the progress of the monitored stream.
<原文结束>

# <翻译开始>
// Sample 添加对被监控流进度的新采样。
# <翻译结束>


<原文开始>
	// BPS returns the calculated Bytes Per Second rate of the monitored stream.
<原文结束>

# <翻译开始>
// BPS 返回监控流计算出的每秒字节数（Bytes Per Second）速率。
# <翻译结束>


<原文开始>
// SampleFunc is used by Watch to take periodic samples of a monitored stream.
<原文结束>

# <翻译开始>
// SampleFunc 由 Watch 调用，用于周期性地对被监控的流进行采样。
# <翻译结束>


<原文开始>
// Watch will periodically call the given SampleFunc to sample the progress of
// a monitored stream and update the given gauge. SampleFunc should return the
// total number of bytes transferred by the stream since it started.
//
// Watch is a blocking call and should typically be called in a new goroutine.
// To prevent the goroutine from leaking, make sure to cancel the given context
// once the stream is completed or canceled.
<原文结束>

# <翻译开始>
// Watch 会周期性调用给定的 SampleFunc 来采样被监控流的进度，并更新给定的仪表（gauge）。SampleFunc 应该返回流自从开始以来传输的总字节数。
//
// Watch 是一个阻塞调用，通常应在新的 goroutine 中调用。为了避免 goroutine 泄漏，在流完成或被取消时，请确保取消给定的上下文。
# <翻译结束>

