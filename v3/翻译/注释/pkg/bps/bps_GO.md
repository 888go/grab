
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
// Gauge是该包中所有BPS度量的通用接口。给定一段时间内的多个采样，每种度量类型都可以用来测量数据流的字节/秒传输速率。
// 
// 所有采样点的时间戳和值都必须单调递增。每个采样点应表示流中的总字节数，而不是自上一个采样点以来发送的字节数。
// 
// 为了确保度量能够尽快报告进度，请在数据流开始时立即获取初始采样。
// 
// 所有度量实现都支持并发使用。
// md5:96b611c42ceb040c
# <翻译结束>


<原文开始>
// Sample adds a new sample of the progress of the monitored stream.
<原文结束>

# <翻译开始>
// Sample 添加监控流的新进度样本。 md5:d2b29305c97808d1
# <翻译结束>


<原文开始>
// BPS returns the calculated Bytes Per Second rate of the monitored stream.
<原文结束>

# <翻译开始>
// BPS 返回监控流的计算出的每秒字节数速率。 md5:6ca06acbcf94e710
# <翻译结束>


<原文开始>
// SampleFunc is used by Watch to take periodic samples of a monitored stream.
<原文结束>

# <翻译开始>
// SampleFunc被Watch用来定期从监控流中采样。 md5:4f93786f07892dce
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
// Watch 会定期调用给定的 SampleFunc 来采样监控流的进度，并更新给定的指标。SampleFunc 应返回自流开始以来已传输的总字节数。
// 
// Watch 是一个阻塞调用，通常应在新的goroutine中进行。为了防止goroutine泄漏，请确保在流完成或取消时取消给定的上下文。
// md5:866a6d9ed56020b3
# <翻译结束>

