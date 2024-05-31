
<原文开始>
// Response represents the response to a completed or in-progress download
// request.
//
// A response may be returned as soon a HTTP response is received from a remote
// server, but before the body content has started transferring.
//
// All Response method calls are thread-safe.
<原文结束>

# <翻译开始>
// Response 表示已完成或正在进行的下载请求的响应。
//
//一旦从远程服务器接收到HTTP响应，就可能返回一个响应，但在开始传输主体内容之前。
//
//对Response方法的所有调用都是线程安全的。
// md5:ddbfa271b1d4fd24
# <翻译结束>


<原文开始>
// The Request that was submitted to obtain this Response.
<原文结束>

# <翻译开始>
// 提交以获取此响应的请求。 md5:75fc5e8790cf3c19
# <翻译结束>


<原文开始>
	// HTTPResponse represents the HTTP response received from an HTTP request.
	//
	// The response Body should not be used as it will be consumed and closed by
	// grab.
<原文结束>

# <翻译开始>
// HTTPResponse 表示从 HTTP 请求接收到的 HTTP 响应。
// 
// 响应体不应该被使用，因为它会被 grab 消费并关闭。
// md5:de050e04bc066b54
# <翻译结束>


<原文开始>
	// Filename specifies the path where the file transfer is stored in local
	// storage.
<原文结束>

# <翻译开始>
// Filename 指定了文件传输在本地存储中的路径。
// md5:2739dccaf6d8bae1
# <翻译结束>


<原文开始>
// Size specifies the total expected size of the file transfer.
<原文结束>

# <翻译开始>
// Size 指定文件传输的总预期大小。 md5:6c3a6da7f95a7aa6
# <翻译结束>


<原文开始>
// Start specifies the time at which the file transfer started.
<原文结束>

# <翻译开始>
// Start 指定文件传输开始的时间。 md5:0b7ef91daaa881e6
# <翻译结束>


<原文开始>
	// End specifies the time at which the file transfer completed.
	//
	// This will return zero until the transfer has completed.
<原文结束>

# <翻译开始>
// End 指定文件传输完成的时间。
//
// 在传输完成之前，此函数将返回零值。
// md5:22aad30fb1342e5f
# <翻译结束>


<原文开始>
	// CanResume specifies that the remote server advertised that it can resume
	// previous downloads, as the 'Accept-Ranges: bytes' header is set.
<原文结束>

# <翻译开始>
// CanResume 指定远程服务器已宣布它可以恢复以前的下载，因为已设置 "Accept-Ranges: bytes" 头。
// md5:d5b949adac4b5a26
# <翻译结束>


<原文开始>
	// DidResume specifies that the file transfer resumed a previously incomplete
	// transfer.
<原文结束>

# <翻译开始>
// DidResume 指定文件传输恢复了先前未完成的传输。
// md5:9e59c01300744b13
# <翻译结束>


<原文开始>
	// Done is closed once the transfer is finalized, either successfully or with
	// errors. Errors are available via Response.Err
<原文结束>

# <翻译开始>
// Done 在传输完成时关闭，无论是成功还是出现错误。可以通过 Response.Err 获取错误信息。
// md5:a87b182e4efaa89f
# <翻译结束>


<原文开始>
// ctx is a Context that controls cancelation of an inprogress transfer
<原文结束>

# <翻译开始>
// ctx是一个控制进行中传输取消的Context md5:a1a2dcb143eb803b
# <翻译结束>


<原文开始>
	// cancel is a cancel func that can be used to cancel the context of this
	// Response.
<原文结束>

# <翻译开始>
// cancel 是一个取消函数，可以用来取消这个Response的上下文。
// md5:7803555a55294ebb
# <翻译结束>


<原文开始>
	// fi is the FileInfo for the destination file if it already existed before
	// transfer started.
<原文结束>

# <翻译开始>
// fi 是目标文件在传输开始前已存在的FileInfo。
// md5:80de0f1066a0dd4a
# <翻译结束>


<原文开始>
	// optionsKnown indicates that a HEAD request has been completed and the
	// capabilities of the remote server are known.
<原文结束>

# <翻译开始>
// optionsKnown表示已经完成了HEAD请求，并且知道了远程服务器的功能。
// md5:1808cb9d787356c6
# <翻译结束>


<原文开始>
	// writer is the file handle used to write the downloaded file to local
	// storage
<原文结束>

# <翻译开始>
// writer 是用于将下载的文件写入本地存储的文件句柄
// md5:bd2f23db00719ebb
# <翻译结束>


<原文开始>
	// storeBuffer receives the contents of the transfer if Request.NoStore is
	// enabled.
<原文结束>

# <翻译开始>
// storeBuffer 接收传输的内容，如果 Request.NoStore 开关启用的话。
// md5:c4f455f8b39f4595
# <翻译结束>


<原文开始>
	// bytesCompleted specifies the number of bytes which were already
	// transferred before this transfer began.
<原文结束>

# <翻译开始>
// bytesCompleted 指定的是在本次传输开始之前已经传输的字节数。
// md5:03aa448103fe412e
# <翻译结束>


<原文开始>
	// transfer is responsible for copying data from the remote server to a local
	// file, tracking progress and allowing for cancelation.
<原文结束>

# <翻译开始>
// transfer负责从远程服务器复制数据到本地文件，跟踪进度并允许取消。
// md5:9a7cce1d1337a757
# <翻译结束>


<原文开始>
// bufferSize specifies the size in bytes of the transfer buffer.
<原文结束>

# <翻译开始>
// bufferSize 指定了传输缓冲区的字节大小。 md5:fc0bc9d5d97bffef
# <翻译结束>


<原文开始>
	// Error contains any error that may have occurred during the file transfer.
	// This should not be read until IsComplete returns true.
<原文结束>

# <翻译开始>
// Error 包含文件传输过程中可能发生的任何错误。在 IsComplete 返回 true 之前，不应读取此值。
// md5:5938a32425f8b1ad
# <翻译结束>


<原文开始>
// IsComplete returns true if the download has completed. If an error occurred
// during the download, it can be returned via Err.
<原文结束>

# <翻译开始>
// IsComplete 返回如果下载已完成为true。如果在下载过程中发生错误，可以通过Err返回该错误。
// md5:eb0348c8e5031ff2
# <翻译结束>


<原文开始>
// Cancel cancels the file transfer by canceling the underlying Context for
// this Response. Cancel blocks until the transfer is closed and returns any
// error - typically context.Canceled.
<原文结束>

# <翻译开始>
// Cancel 通过取消此 Response 的底层上下文来取消文件传输。Cancel 函数会阻塞直到传输关闭，并返回任何错误——通常为 context.Canceled。
// md5:4277ad03dbb17f34
# <翻译结束>


<原文开始>
// Wait blocks until the download is completed.
<原文结束>

# <翻译开始>
// Wait 会阻塞直到下载完成。 md5:2c909b1e4febc570
# <翻译结束>


<原文开始>
// Err blocks the calling goroutine until the underlying file transfer is
// completed and returns any error that may have occurred. If the download is
// already completed, Err returns immediately.
<原文结束>

# <翻译开始>
// Err会阻塞调用的goroutine，直到底层文件传输完成，并返回可能发生的任何错误。如果下载已经完成，Err会立即返回。
// md5:1eec3a720d62a01e
# <翻译结束>


<原文开始>
// Size returns the size of the file transfer. If the remote server does not
// specify the total size and the transfer is incomplete, the return value is
// -1.
<原文结束>

# <翻译开始>
// Size 返回文件传输的大小。如果远程服务器没有指定总大小且传输不完整，返回值为 -1。
// md5:53bd23aba7b5b22a
# <翻译结束>


<原文开始>
// BytesComplete returns the total number of bytes which have been copied to
// the destination, including any bytes that were resumed from a previous
// download.
<原文结束>

# <翻译开始>
// BytesComplete返回已复制到目标的总字节数，包括从先前下载中恢复的任何字节。
// md5:8b9d515c7cdd428f
# <翻译结束>


<原文开始>
// BytesPerSecond returns the number of bytes per second transferred using a
// simple moving average of the last five seconds. If the download is already
// complete, the average bytes/sec for the life of the download is returned.
<原文结束>

# <翻译开始>
// BytesPerSecond 返回过去五秒内的平均每秒传输字节数。如果下载已完成，则返回整个下载过程中的平均字节数/秒。
// md5:dd425949dfaaf72f
# <翻译结束>


<原文开始>
// Progress returns the ratio of total bytes that have been downloaded. Multiply
// the returned value by 100 to return the percentage completed.
<原文结束>

# <翻译开始>
// Progress 返回已下载的总字节数的比例。将返回值乘以100可得到完成的百分比。
// md5:163669af2595df84
# <翻译结束>


<原文开始>
// Duration returns the duration of a file transfer. If the transfer is in
// process, the duration will be between now and the start of the transfer. If
// the transfer is complete, the duration will be between the start and end of
// the completed transfer process.
<原文结束>

# <翻译开始>
// Duration 返回文件传输的持续时间。如果传输正在进行中，持续时间将是当前时间和传输开始时间之间的差。如果传输已完成，持续时间将是完成传输过程的开始和结束时间之间的差。
// md5:851e9a4335644cc7
# <翻译结束>


<原文开始>
// ETA returns the estimated time at which the the download will complete, given
// the current BytesPerSecond. If the transfer has already completed, the actual
// end time will be returned.
<原文结束>

# <翻译开始>
// ETA 返回给定当前每秒字节数（BytesPerSecond）下载预计完成的时间。如果传输已经完成，将返回实际结束时间。
// md5:7f8fcb12ee64da7f
# <翻译结束>


<原文开始>
// Open blocks the calling goroutine until the underlying file transfer is
// completed and then opens the transferred file for reading. If Request.NoStore
// was enabled, the reader will read from memory.
//
// If an error occurred during the transfer, it will be returned.
//
// It is the callers responsibility to close the returned file handle.
<原文结束>

# <翻译开始>
// Open函数会让调用的goroutine阻塞，直到底层文件传输完成，然后打开已传输的文件以供阅读。如果Request.NoStore选项被启用，读取器将从内存中读取。
//
// 如果在传输过程中发生错误，该错误将被返回。
//
// 调用者有责任关闭返回的文件句柄。
// md5:92d1addd1c15e703
# <翻译结束>


<原文开始>
// Bytes blocks the calling goroutine until the underlying file transfer is
// completed and then reads all bytes from the completed tranafer. If
// Request.NoStore was enabled, the bytes will be read from memory.
//
// If an error occurred during the transfer, it will be returned.
<原文结束>

# <翻译开始>
// Bytes 使调用的goroutine阻塞，直到底层文件传输完成，然后读取已完成传输的所有字节。如果启用了Request.NoStore，将从内存中读取字节。
// 
// 如果在传输过程中发生错误，将返回该错误。
// md5:af539342438b13c1
# <翻译结束>

