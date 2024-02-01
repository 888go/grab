
<原文开始>
// Get sends a HTTP request and downloads the content of the requested URL to
// the given destination file path. The caller is blocked until the download is
// completed, successfully or otherwise.
//
// An error is returned if caused by client policy (such as CheckRedirect), or
// if there was an HTTP protocol or IO error.
//
// For non-blocking calls or control over HTTP client headers, redirect policy,
// and other settings, create a Client instead.
<原文结束>

# <翻译开始>
// Get通过发送HTTP请求并下载请求URL的内容到指定的目标文件路径。调用方会一直阻塞，直到下载完成（无论成功与否）。
//
// 若因客户端策略（例如CheckRedirect）导致错误，或者发生HTTP协议错误或IO错误，将会返回错误。
//
// 对于非阻塞调用，或需要控制HTTP客户端头信息、重定向策略和其他设置的情况，请创建一个Client实例代替。
# <翻译结束>


<原文开始>
// GetBatch sends multiple HTTP requests and downloads the content of the
// requested URLs to the given destination directory using the given number of
// concurrent worker goroutines.
//
// The Response for each requested URL is sent through the returned Response
// channel, as soon as a worker receives a response from the remote server. The
// Response can then be used to track the progress of the download while it is
// in progress.
//
// The returned Response channel will be closed by Grab, only once all downloads
// have completed or failed.
//
// If an error occurs during any download, it will be available via call to the
// associated Response.Err.
//
// For control over HTTP client headers, redirect policy, and other settings,
// create a Client instead.
<原文结束>

# <翻译开始>
// GetBatch 发送多个 HTTP 请求，并使用指定数量的并发工作 goroutine 将请求的 URL 内容下载到给定的目标目录。
//
// 当一个工作线程从远程服务器接收到响应时，每个请求的 URL 的 Response 会立即通过返回的 Response 通道发送出去。这样，在下载过程中就可以利用这个 Response 来跟踪下载进度。
//
// 只有当所有下载任务完成或失败后，由 Grab 返回的 Response 通道才会被关闭。
//
// 如果在任何下载过程中发生错误，可以通过调用相关的 Response.Err 来获取该错误信息。
//
// 如果需要控制 HTTP 客户端头、重定向策略以及其他设置，请创建一个 Client 对象来代替。
# <翻译结束>

