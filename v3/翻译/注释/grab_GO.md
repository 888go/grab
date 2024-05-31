
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
// Get 发送一个HTTP请求，并将请求的URL内容下载到给定的目标文件路径。调用者会在下载完成，无论成功还是失败时被阻塞。
// 
// 如果由于客户端策略（如CheckRedirect）导致错误，或者出现HTTP协议或IO错误，会返回一个错误。
// 
// 若要进行非阻塞调用，或者控制HTTP客户端头、重定向策略和其他设置，请创建一个Client。
// md5:1bbde4e6e89ceb15
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
// GetBatch 发送多个HTTP请求，并使用给定的并发工作goroutine数量将请求的URL内容下载到指定的目标目录。
//
// 每个请求的URL的Response会通过返回的Response通道发送，一旦工作goroutine从远程服务器接收到响应。然后可以使用Response来跟踪下载进度。
//
// Grab会在所有下载完成或失败后关闭返回的Response通道。
//
// 如果在任何下载过程中发生错误，可以通过调用关联的Response.Err获取该错误。
//
// 若要控制HTTP客户端头部、重定向策略和其他设置，应创建一个Client。
// md5:cfa454826e483447
# <翻译结束>

