
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
// 即使在远程服务器接收到 HTTP 响应，但内容尚未开始传输时，也可能返回一个响应。
//
// 所有 Response 方法调用都是线程安全的。
# <翻译结束>


<原文开始>
	// The Request that was submitted to obtain this Response.
<原文结束>

# <翻译开始>
// 该Response所对应的请求。
# <翻译结束>


<原文开始>
	// HTTPResponse represents the HTTP response received from an HTTP request.
	//
	// The response Body should not be used as it will be consumed and closed by
	// grab.
<原文结束>

# <翻译开始>
// HTTPResponse 代表从 HTTP 请求中接收到的 HTTP 响应。
//
// 不应使用响应体（response Body），因为它将被 grab 消耗并关闭。
# <翻译结束>


<原文开始>
	// Filename specifies the path where the file transfer is stored in local
	// storage.
<原文结束>

# <翻译开始>
// Filename 指定文件传输在本地存储中的保存路径。
# <翻译结束>


<原文开始>
	// Size specifies the total expected size of the file transfer.
<原文结束>

# <翻译开始>
// Size 指定文件传输的预期总大小。
# <翻译结束>


<原文开始>
	// Start specifies the time at which the file transfer started.
<原文结束>

# <翻译开始>
// Start 指定文件传输开始的时间。
# <翻译结束>


<原文开始>
	// End specifies the time at which the file transfer completed.
	//
	// This will return zero until the transfer has completed.
<原文结束>

# <翻译开始>
// End 指定文件传输完成的时间。
//
// 在传输尚未完成时，此属性将返回零值。
# <翻译结束>


<原文开始>
	// CanResume specifies that the remote server advertised that it can resume
	// previous downloads, as the 'Accept-Ranges: bytes' header is set.
<原文结束>

# <翻译开始>
// CanResume 指定远程服务器声明它可以恢复先前的下载，因为已设置了 'Accept-Ranges: bytes' 头部。
# <翻译结束>


<原文开始>
	// DidResume specifies that the file transfer resumed a previously incomplete
	// transfer.
<原文结束>

# <翻译开始>
// DidResume 指定文件传输恢复了先前未完成的传输。
# <翻译结束>


<原文开始>
	// Done is closed once the transfer is finalized, either successfully or with
	// errors. Errors are available via Response.Err
<原文结束>

# <翻译开始>
// Done 在传输完成（无论成功或出现错误）后关闭。通过 Response.Err 可获取错误信息
# <翻译结束>


<原文开始>
	// ctx is a Context that controls cancelation of an inprogress transfer
<原文结束>

# <翻译开始>
// ctx 是一个 Context，用于控制正在进行的传输的取消
# <翻译结束>


<原文开始>
	// cancel is a cancel func that can be used to cancel the context of this
	// Response.
<原文结束>

# <翻译开始>
// cancel 是一个取消函数，可用于取消此 Response 的上下文。
# <翻译结束>


<原文开始>
	// fi is the FileInfo for the destination file if it already existed before
	// transfer started.
<原文结束>

# <翻译开始>
// fi 是目标文件在传输开始前（如果已经存在）的 FileInfo 信息。
# <翻译结束>


<原文开始>
	// optionsKnown indicates that a HEAD request has been completed and the
	// capabilities of the remote server are known.
<原文结束>

# <翻译开始>
// optionsKnown 表示已完成 HEAD 请求，并且已知远程服务器的功能。
# <翻译结束>


<原文开始>
	// writer is the file handle used to write the downloaded file to local
	// storage
<原文结束>

# <翻译开始>
// writer 是用于将下载的文件写入本地存储的文件句柄
# <翻译结束>


<原文开始>
	// storeBuffer receives the contents of the transfer if Request.NoStore is
	// enabled.
<原文结束>

# <翻译开始>
// storeBuffer 如果 Request.NoStore 被启用，则接收传输的内容。
# <翻译结束>


<原文开始>
	// bytesCompleted specifies the number of bytes which were already
	// transferred before this transfer began.
<原文结束>

# <翻译开始>
// bytesCompleted 指定在这次传输开始之前已经完成传输的字节数。
# <翻译结束>


<原文开始>
	// transfer is responsible for copying data from the remote server to a local
	// file, tracking progress and allowing for cancelation.
<原文结束>

# <翻译开始>
// transfer 负责从远程服务器复制数据到本地文件，
// 并跟踪进度以及允许取消操作。
# <翻译结束>


<原文开始>
	// bufferSize specifies the size in bytes of the transfer buffer.
<原文结束>

# <翻译开始>
// bufferSize 指定了传输缓冲区的大小（以字节为单位）。
# <翻译结束>


<原文开始>
	// Error contains any error that may have occurred during the file transfer.
	// This should not be read until IsComplete returns true.
<原文结束>

# <翻译开始>
// Error 包含了文件传输过程中可能出现的任何错误。
// 请在 IsComplete 返回 true 之前不要读取此内容。
# <翻译结束>


<原文开始>
// IsComplete returns true if the download has completed. If an error occurred
// during the download, it can be returned via Err.
<原文结束>

# <翻译开始>
// IsComplete 返回一个布尔值，如果下载已完成则返回true。如果在下载过程中发生错误，可以通过 Err 返回该错误。
# <翻译结束>


<原文开始>
// Cancel cancels the file transfer by canceling the underlying Context for
// this Response. Cancel blocks until the transfer is closed and returns any
// error - typically context.Canceled.
<原文结束>

# <翻译开始>
// Cancel 取消文件传输，通过取消此 Response 对应的基础 Context 来实现。Cancel 会阻塞直到传输关闭并返回任何错误——通常是 context.Canceled。
# <翻译结束>


<原文开始>
// Wait blocks until the download is completed.
<原文结束>

# <翻译开始>
// Wait会阻塞直到下载完成。
# <翻译结束>


<原文开始>
// Err blocks the calling goroutine until the underlying file transfer is
// completed and returns any error that may have occurred. If the download is
// already completed, Err returns immediately.
<原文结束>

# <翻译开始>
// Err 阻塞调用该方法的 goroutine，直到底层文件传输完成，并返回在此期间可能发生的任何错误。如果下载已经完成，Err 将立即返回。
# <翻译结束>


<原文开始>
// Size returns the size of the file transfer. If the remote server does not
// specify the total size and the transfer is incomplete, the return value is
// -1.
<原文结束>

# <翻译开始>
// Size 返回文件传输的大小。如果远程服务器没有指定总大小，并且传输未完成，则返回值为-1。
# <翻译结束>


<原文开始>
// BytesComplete returns the total number of bytes which have been copied to
// the destination, including any bytes that were resumed from a previous
// download.
<原文结束>

# <翻译开始>
// BytesComplete 返回已复制到目标位置的总字节数，包括从先前下载恢复的所有字节。
# <翻译结束>


<原文开始>
// BytesPerSecond returns the number of bytes per second transferred using a
// simple moving average of the last five seconds. If the download is already
// complete, the average bytes/sec for the life of the download is returned.
<原文结束>

# <翻译开始>
// BytesPerSecond 返回过去五秒钟内通过简单移动平均计算出的每秒传输字节数。如果下载已经完成，则返回整个下载过程中平均的每秒字节数。
# <翻译结束>


<原文开始>
// Progress returns the ratio of total bytes that have been downloaded. Multiply
// the returned value by 100 to return the percentage completed.
<原文结束>

# <翻译开始>
// Progress 返回已下载总字节的比例。将返回的值乘以100可得到完成的百分比。
# <翻译结束>


<原文开始>
// Duration returns the duration of a file transfer. If the transfer is in
// process, the duration will be between now and the start of the transfer. If
// the transfer is complete, the duration will be between the start and end of
// the completed transfer process.
<原文结束>

# <翻译开始>
// Duration 返回文件传输的持续时间。如果传输正在进行中，
// 持续时间将是从现在到传输开始之间的时间差。如果传输已完成，
// 持续时间将是整个完成传输过程从开始到结束的时间差。
# <翻译结束>


<原文开始>
// ETA returns the estimated time at which the the download will complete, given
// the current BytesPerSecond. If the transfer has already completed, the actual
// end time will be returned.
<原文结束>

# <翻译开始>
// ETA 返回根据当前每秒字节数估算的下载完成时间。如果传输已完成，将返回实际结束时间。
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
// Open 将阻塞调用的 goroutine，直到底层文件传输完成，然后打开已传输的文件以供读取。
// 如果启用了 Request.NoStore，则读取器将从内存中读取数据。
//
// 如果在传输过程中发生错误，将会返回该错误。
//
// 调用者有责任关闭返回的文件句柄。
# <翻译结束>


<原文开始>
// Bytes blocks the calling goroutine until the underlying file transfer is
// completed and then reads all bytes from the completed tranafer. If
// Request.NoStore was enabled, the bytes will be read from memory.
//
// If an error occurred during the transfer, it will be returned.
<原文结束>

# <翻译开始>
// Bytes 阻塞调用它的 goroutine，直到底层文件传输完成，然后从已完成的传输中读取所有字节。
// 如果启用了 Request.NoStore，则字节将从内存中读取。
//
// 如果在传输过程中发生错误，将会返回该错误。
# <翻译结束>

