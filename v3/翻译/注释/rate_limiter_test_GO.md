
<原文开始>
// testRateLimiter is a naive rate limiter that limits throughput to r tokens
// per second. The total number of tokens issued is tracked as n.
<原文结束>

# <翻译开始>
// testRateLimiter是一个简单的速率限制器，每秒限制r个令牌的通过量。已发出的令牌总数由n跟踪。
// md5:d54804ac9a0ceb1a
# <翻译结束>


<原文开始>
	// download a 128 byte file, 8 bytes at a time, with a naive 512bps limiter
	// should take > 250ms
<原文结束>

# <翻译开始>
// 使用简单的512bps限制，一次下载128字节的文件，每次8字节。应该花费超过250毫秒。
// md5:e92aa1bfdda7be8a
# <翻译结束>


<原文开始>
// ensure multiple trips to the rate limiter by downloading 8 bytes at a time
<原文结束>

# <翻译开始>
// 确保每次下载8字节，以实现对速率限制器的多次访问 md5:d541457cb761ceaf
# <翻译结束>


<原文开始>
// BUG: this test can pass if the transfer was slow for unrelated reasons
<原文结束>

# <翻译开始>
// BUG：如果由于无关原因（transfer速度慢）导致测试通过，这个注释表示这是一个已知问题。 md5:a7882bdf6de8b17f
# <翻译结束>


<原文开始>
	// Attach a 1Mbps rate limiter, like the token bucket implementation from
	// golang.org/x/time/rate.
<原文结束>

# <翻译开始>
// 附加一个1Mbps速率限制器，类似于golang.org/x/time/rate包中的令牌桶实现。
// md5:a33151d1a418a247
# <翻译结束>

