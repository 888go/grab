
<原文开始>
// This example uses DoChannel to create a Producer/Consumer model for
// downloading multiple files concurrently. This is similar to how DoBatch uses
// DoChannel under the hood except that it allows the caller to continually send
// new requests until they wish to close the request channel.
<原文结束>

# <翻译开始>
// 此示例使用DoChannel创建一个生产者/消费者模型，用于并发下载多个文件。
// 这与DoBatch在内部使用DoChannel的工作方式类似，不同之处在于它允许调用者持续发送新请求，
// 直到他们希望关闭请求通道为止。
// md5:8cfd63343a82362c
# <翻译结束>


<原文开始>
// create a request and a buffered response channel
<原文结束>

# <翻译开始>
// 创建一个请求和一个缓冲的响应通道 md5:c4b7d02b204abf79
# <翻译结束>


<原文开始>
// wait for workers to finish
<原文结束>

# <翻译开始>
// 等待工人完成 md5:b758a6709f2c0bc0
# <翻译结束>


<原文开始>
// create multiple download requests
<原文结束>

# <翻译开始>
// 创建多个下载请求 md5:763389fc7a54db5a
# <翻译结束>


<原文开始>
// start downloads with 4 workers
<原文结束>

# <翻译开始>
// 使用4个工人开始下载 md5:f23bbcd916bb7cfc
# <翻译结束>

