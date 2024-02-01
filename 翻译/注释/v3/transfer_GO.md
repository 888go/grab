
<原文开始>
// must be 64bit aligned on 386
<原文结束>

# <翻译开始>
// 必须在386架构上进行64位对齐
# <翻译结束>


<原文开始>
// five second moving average sampling every second
<原文结束>

# <翻译开始>
// 每隔一秒进行一次采样，计算五秒内的移动平均值
# <翻译结束>


<原文开始>
// copy behaves similarly to io.CopyBuffer except that it checks for cancelation
// of the given context.Context, reports progress in a thread-safe manner and
// tracks the transfer rate.
<原文结束>

# <翻译开始>
// copy 函数的行为类似于 io.CopyBuffer，但除此之外，它还会检查给定的 context.Context 是否已取消，
// 以线程安全的方式报告进度，并跟踪传输速率。
# <翻译结束>


<原文开始>
	// maintain a bps gauge in another goroutine
<原文结束>

# <翻译开始>
// 在另一个goroutine中维护一个bps计数器
# <翻译结束>


<原文开始>
	// start the transfer
<原文结束>

# <翻译开始>
// 开始传输
# <翻译结束>


<原文开始>
			// wait for rate limiter
<原文结束>

# <翻译开始>
// 等待速率限制器
# <翻译结束>


<原文开始>
// N returns the number of bytes transferred.
<原文结束>

# <翻译开始>
// N 返回已传输的字节数。
# <翻译结束>


<原文开始>
// BPS returns the current bytes per second transfer rate using a simple moving
// average.
<原文结束>

# <翻译开始>
// BPS 返回当前的字节每秒传输速率，采用简单移动平均算法计算。
# <翻译结束>

