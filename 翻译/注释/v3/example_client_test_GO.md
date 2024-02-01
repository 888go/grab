
<原文开始>
// This example uses DoChannel to create a Producer/Consumer model for
// downloading multiple files concurrently. This is similar to how DoBatch uses
// DoChannel under the hood except that it allows the caller to continually send
// new requests until they wish to close the request channel.
<原文结束>

# <翻译开始>
// 此示例使用 DoChannel 创建一个用于并发下载多个文件的生产者/消费者模型。这与 DoBatch 在底层使用 DoChannel 的方式类似，但不同之处在于它允许调用者持续发送新的请求，直到他们希望关闭请求通道为止。
# <翻译结束>


<原文开始>
	// create a request and a buffered response channel
<原文结束>

# <翻译开始>
// 创建一个请求和一个缓冲的响应通道
# <翻译结束>


<原文开始>
		// wait for workers to finish
<原文结束>

# <翻译开始>
// 等待所有工作者完成
# <翻译结束>


<原文开始>
	// check each response
<原文结束>

# <翻译开始>
// 检查每个响应
# <翻译结束>


<原文开始>
		// block until complete
<原文结束>

# <翻译开始>
// 等待直到完成
# <翻译结束>


<原文开始>
	// create multiple download requests
<原文结束>

# <翻译开始>
// 创建多个下载请求
# <翻译结束>


<原文开始>
	// start downloads with 4 workers
<原文结束>

# <翻译开始>
// 使用4个工人开始下载
# <翻译结束>

