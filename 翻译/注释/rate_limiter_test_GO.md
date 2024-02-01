
<原文开始>
// testRateLimiter is a naive rate limiter that limits throughput to r tokens
// per second. The total number of tokens issued is tracked as n.
<原文结束>

# <翻译开始>
// testRateLimiter 是一个简单的速率限制器，其功能是将吞吐量限制为每秒 r 个令牌。已发放令牌的总数通过变量 n 进行跟踪。
// 注释翻译如下：
// ```go
// testRateLimiter 是一个粗略的速率限制器，用于将处理速率限制为每秒r个令牌。已发放的令牌总数量通过变量n进行统计。
# <翻译结束>


<原文开始>
	// download a 128 byte file, 8 bytes at a time, with a naive 512bps limiter
	// should take > 250ms
<原文结束>

# <翻译开始>
// 下载一个128字节的文件，每次下载8字节，并采用一个简单的512bps限速器
// 预计耗时超过250毫秒
# <翻译结束>


<原文开始>
		// ensure multiple trips to the rate limiter by downloading 8 bytes at a time
<原文结束>

# <翻译开始>
// 确保通过每次下载8字节的方式多次访问速率限制器
# <翻译结束>


<原文开始>
			// BUG: this test can pass if the transfer was slow for unrelated reasons
<原文结束>

# <翻译开始>
// BUG: 如果由于无关原因传输速度较慢，这个测试可能会通过
# <翻译结束>


<原文开始>
	// Attach a 1Mbps rate limiter, like the token bucket implementation from
	// golang.org/x/time/rate.
<原文结束>

# <翻译开始>
// 绑定一个1Mbps的限速器，类似于来自golang.org/x/time/rate包中的令牌桶实现。
# <翻译结束>

