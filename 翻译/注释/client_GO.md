
<原文开始>
// HTTPClient provides an interface allowing us to perform HTTP requests.
<原文结束>

# <翻译开始>
// HTTPClient 提供了一个接口，使我们能够执行 HTTP 请求。
# <翻译结束>


<原文开始>
// truncater is a private interface allowing different response
// Writers to be truncated
<原文结束>

# <翻译开始>
// truncater 是一个私有接口，用于支持不同的响应写入器进行截断操作
# <翻译结束>


<原文开始>
// A Client is a file download client.
//
// Clients are safe for concurrent use by multiple goroutines.
<原文结束>

# <翻译开始>
// Client 是一个文件下载客户端。
//
// Client 对象支持在多个goroutine中并发安全地使用。
# <翻译结束>


<原文开始>
	// HTTPClient specifies the http.Client which will be used for communicating
	// with the remote server during the file transfer.
<原文结束>

# <翻译开始>
// HTTPClient 指定在文件传输过程中与远程服务器通信时所使用的 http.Client。
# <翻译结束>


<原文开始>
	// UserAgent specifies the User-Agent string which will be set in the
	// headers of all requests made by this client.
	//
	// The user agent string may be overridden in the headers of each request.
<原文结束>

# <翻译开始>
// UserAgent 指定此客户端发起的所有请求中将在头部设置的 User-Agent 字符串。
//
// 用户代理字符串可以在每个请求的头部进行覆盖。
# <翻译结束>


<原文开始>
	// BufferSize specifies the size in bytes of the buffer that is used for
	// transferring all requested files. Larger buffers may result in faster
	// throughput but will use more memory and result in less frequent updates
	// to the transfer progress statistics. The BufferSize of each request can
	// be overridden on each Request object. Default: 32KB.
<原文结束>

# <翻译开始>
// BufferSize 指定用于传输所有请求文件的缓冲区大小（以字节为单位）。更大的缓冲区可能会带来更快的数据传输速度，但会消耗更多的内存，并导致传输进度统计信息更新频率降低。每个请求都可以在各自的 Request 对象上覆盖 BufferSize 属性。默认值：32KB。
# <翻译结束>


<原文开始>
// NewClient returns a new file download Client, using default configuration.
<原文结束>

# <翻译开始>
// NewClient 返回一个使用默认配置的新文件下载客户端。
# <翻译结束>


<原文开始>
// DefaultClient is the default client and is used by all Get convenience
// functions.
<原文结束>

# <翻译开始>
// DefaultClient 是默认的客户端，被所有 Get 方便函数使用。
# <翻译结束>


<原文开始>
// Do sends a file transfer request and returns a file transfer response,
// following policy (e.g. redirects, cookies, auth) as configured on the
// client's HTTPClient.
//
// Like http.Get, Do blocks while the transfer is initiated, but returns as soon
// as the transfer has started transferring in a background goroutine, or if it
// failed early.
//
// An error is returned via Response.Err if caused by client policy (such as
// CheckRedirect), or if there was an HTTP protocol or IO error. Response.Err
// will block the caller until the transfer is completed, successfully or
// otherwise.
<原文结束>

# <翻译开始>
// Do 发送文件传输请求并返回文件传输响应，遵循客户端 HTTPClient 上配置的策略（例如重定向、Cookie 和身份验证）。
//
// 类似于 http.Get，Do 在传输开始时阻塞，但一旦传输在后台 goroutine 中启动或早期失败，则立即返回。
//
// 若因客户端策略（如 CheckRedirect）导致错误，或者发生 HTTP 协议错误或 IO 错误，将通过 Response.Err 返回错误。Response.Err 将阻塞调用者直到传输完成（无论成功与否）。
# <翻译结束>


<原文开始>
	// cancel will be called on all code-paths via closeResponse
<原文结束>

# <翻译开始>
// 当通过 closeResponse 在所有代码路径上调用时，cancel 将被调用
# <翻译结束>


<原文开始>
		// default to Client.BufferSize
<原文结束>

# <翻译开始>
// 默认为Client.BufferSize
# <翻译结束>


<原文开始>
	// Run state-machine while caller is blocked to initialize the file transfer.
	// Must never transition to the copyFile state - this happens next in another
	// goroutine.
<原文结束>

# <翻译开始>
// 当调用者阻塞以初始化文件传输时，运行状态机。
// 绝对不能转换到 copyFile 状态——这将在另一个 goroutine 中紧接着发生。
# <翻译结束>


<原文开始>
	// Run copyFile in a new goroutine. copyFile will no-op if the transfer is
	// already complete or failed.
<原文结束>

# <翻译开始>
// 在一个新的goroutine中运行copyFile。如果传输已经完成或失败，copyFile将执行空操作（不做任何事）。
# <翻译结束>


<原文开始>
// DoChannel executes all requests sent through the given Request channel, one
// at a time, until it is closed by another goroutine. The caller is blocked
// until the Request channel is closed and all transfers have completed. All
// responses are sent through the given Response channel as soon as they are
// received from the remote servers and can be used to track the progress of
// each download.
//
// Slow Response receivers will cause a worker to block and therefore delay the
// start of the transfer for an already initiated connection - potentially
// causing a server timeout. It is the caller's responsibility to ensure a
// sufficient buffer size is used for the Response channel to prevent this.
//
// If an error occurs during any of the file transfers it will be accessible via
// the associated Response.Err function.
<原文结束>

# <翻译开始>
// DoChannel 针对给定的 Request 通道执行所有发送过来的请求，每次执行一个，直到被另一个 goroutine 关闭。调用方将被阻塞，直到 Request 通道关闭且所有传输完成。所有从远程服务器接收到的响应都会立即通过给定的 Response 通道发送，并可用于跟踪每个下载进度。
//
// 如果 Response 接收者处理速度较慢，将会导致工作线程阻塞，从而延迟已建立连接的传输开始时间，可能引发服务器超时。调用方有责任确保为 Response 通道使用足够大的缓冲区大小以防止此类情况发生。
//
// 如果在文件传输过程中出现任何错误，可通过相应的 Response.Err 函数访问该错误。
# <翻译结束>


<原文开始>
	// TODO: enable cancelling of batch jobs
<原文结束>

# <翻译开始>
// TODO: 启用批量作业的取消功能
# <翻译结束>


<原文开始>
// DoBatch executes all the given requests using the given number of concurrent
// workers. Control is passed back to the caller as soon as the workers are
// initiated.
//
// If the requested number of workers is less than one, a worker will be created
// for every request. I.e. all requests will be executed concurrently.
//
// If an error occurs during any of the file transfers it will be accessible via
// call to the associated Response.Err.
//
// The returned Response channel is closed only after all of the given Requests
// have completed, successfully or otherwise.
<原文结束>

# <翻译开始>
// DoBatch 使用给定数量的并发工作者执行所有给定的请求。一旦工作者启动，控制权将返回给调用者。
//
// 如果请求的工作者数量少于1个，则每个请求都会创建一个工作者。即，所有请求都将并发执行。
//
// 如果在任何文件传输过程中发生错误，可以通过调用相关联的Response.Err获取该错误。
//
// 只有在所有给定的Requests都完成（无论成功与否）后，才会关闭返回的Response通道。
# <翻译结束>


<原文开始>
// An stateFunc is an action that mutates the state of a Response and returns
// the next stateFunc to be called.
<原文结束>

# <翻译开始>
// stateFunc 是一种操作，它会修改 Response 的状态，并返回接下来要调用的下一个 stateFunc。
# <翻译结束>


<原文开始>
// run calls the given stateFunc function and all subsequent returned stateFuncs
// until a stateFunc returns nil or the Response.ctx is canceled. Each stateFunc
// should mutate the state of the given Response until it has completed
// downloading or failed.
<原文结束>

# <翻译开始>
// run调用给定的stateFunc函数及其后续返回的所有stateFuncs，
// 直到某个stateFunc返回nil或者Response.ctx被取消为止。每个stateFunc
// 应该根据需要改变给定Response的状态，直到下载完成或失败。
# <翻译结束>


<原文开始>
// statFileInfo retrieves FileInfo for any local file matching
// Response.Filename.
//
// If the file does not exist, is a directory, or its name is unknown the next
// stateFunc is headRequest.
//
// If the file exists, Response.fi is set and the next stateFunc is
// validateLocal.
//
// If an error occurs, the next stateFunc is closeResponse.
<原文结束>

# <翻译开始>
// statFileInfo 函数用于获取与 Response.Filename 匹配的本地文件的 FileInfo 信息。
//
// 如果该文件不存在，或者是一个目录，亦或是名称未知，则下一个状态函数为 headRequest。
//
// 如果文件存在，将设置 Response.fi，并且下一个状态函数为 validateLocal。
//
// 如果出现错误，下一个状态函数则为 closeResponse。
# <翻译结束>


<原文开始>
// validateLocal compares a local copy of the downloaded file to the remote
// file.
//
// An error is returned if the local file is larger than the remote file, or
// Request.SkipExisting is true.
//
// If the existing file matches the length of the remote file, the next
// stateFunc is checksumFile.
//
// If the local file is smaller than the remote file and the remote server is
// known to support ranged requests, the next stateFunc is getRequest.
<原文结束>

# <翻译开始>
// validateLocal 用于比较本地下载文件的副本与远程文件。
//
// 当本地文件大于远程文件，或者 Request.SkipExisting 为真时，会返回错误。
//
// 如果已存在的文件大小与远程文件相同，则下一个状态函数是 checksumFile。
//
// 如果本地文件小于远程文件，并且已知远程服务器支持范围请求，则下一个状态函数是 getRequest。
# <翻译结束>


<原文开始>
	// determine target file size
<原文结束>

# <翻译开始>
// 确定目标文件大小
# <翻译结束>


<原文开始>
		// size is either actually 0 or unknown
		// if unknown, we ask the remote server
		// if known to be 0, we proceed with a GET
<原文结束>

# <翻译开始>
// size 实际上是0或者未知
// 如果是未知，我们将询问远程服务器
// 如果已知为0，我们将继续进行GET请求
# <翻译结束>


<原文开始>
		// local file matches remote file size - wrap it up
<原文结束>

# <翻译开始>
// 当本地文件与远程文件大小相匹配时 - 完成处理
# <翻译结束>


<原文开始>
		// local file should be overwritten
<原文结束>

# <翻译开始>
// 本地文件应被覆盖
# <翻译结束>


<原文开始>
		// remote size is known, is smaller than local size and we want to resume
<原文结束>

# <翻译开始>
// 远程大小已知，且小于本地大小，我们希望从中断处继续恢复
# <翻译结束>


<原文开始>
		// set resume range on GET request
<原文结束>

# <翻译开始>
// 在GET请求中设置恢复范围
# <翻译结束>


<原文开始>
				// err should be os.PathError and include file path
<原文结束>

# <翻译开始>
// err 应该是 os.PathError 类型，并且包含文件路径信息
# <翻译结束>


<原文开始>
// doHTTPRequest sends a HTTP Request and returns the response
<原文结束>

# <翻译开始>
// doHTTPRequest 发送一个HTTP请求并返回响应
# <翻译结束>


<原文开始>
		// destination path is already known and does not exist
<原文结束>

# <翻译开始>
// 目标路径已知且不存在
# <翻译结束>


<原文开始>
	// In case of redirects during HEAD, record the final URL and use it
	// instead of the original URL when sending future requests.
	// This way we avoid sending potentially unsupported requests to
	// the original URL, e.g. "Range", since it was the final URL
	// that advertised its support.
<原文结束>

# <翻译开始>
// 当HEAD请求过程中出现重定向时，记录最终的URL，并在发送后续请求时使用它替代原始URL。
// 这样做可以避免向原始URL发送可能不受支持的请求，例如"Range"请求，因为是最终的URL声明了对该请求的支持。
# <翻译结束>


<原文开始>
	// TODO: check Content-Range
<原文结束>

# <翻译开始>
// TODO: 检查 Content-Range
# <翻译结束>


<原文开始>
	// check status code
<原文结束>

# <翻译开始>
// 检查状态码
# <翻译结束>


<原文开始>
	// check expected size
<原文结束>

# <翻译开始>
// 检查期望的大小
# <翻译结束>


<原文开始>
		// Request.Filename will be empty or a directory
<原文结束>

# <翻译开始>
// Request.Filename 将会是空值或者是一个目录名
# <翻译结束>


<原文开始>
// openWriter opens the destination file for writing and seeks to the location
// from whence the file transfer will resume.
//
// Requires that Response.Filename and resp.DidResume are already be set.
<原文结束>

# <翻译开始>
// openWriter 打开目标文件以进行写入操作，并定位到文件传输将从中恢复的位置。
//
// 要求已预先设置 Response.Filename 和 resp.DidResume。
# <翻译结束>


<原文开始>
		// compute write flags
<原文结束>

# <翻译开始>
// 计算写入标志
# <翻译结束>


<原文开始>
				// truncate later in copyFile, if not cancelled
				// by BeforeCopy hook
<原文结束>

# <翻译开始>
// 如果在copyFile过程中未被BeforeCopy钩子取消，则稍后截断
# <翻译结束>


<原文开始>
		// seek to start or end
<原文结束>

# <翻译开始>
// 寻找并定位到起始或结束位置
# <翻译结束>


<原文开始>
	// next step is copyFile, but this will be called later in another goroutine
<原文结束>

# <翻译开始>
// 下一步是调用copyFile函数，但该函数将在稍后的另一个goroutine中被调用
# <翻译结束>


<原文开始>
// copy transfers content for a HTTP connection established via Client.do()
<原文结束>

# <翻译开始>
// copy 将通过Client.do()建立的HTTP连接的内容进行传输
# <翻译结束>


<原文开始>
	// run BeforeCopy hook
<原文结束>

# <翻译开始>
// 执行 BeforeCopy 钩子函数
# <翻译结束>


<原文开始>
	// We waited to truncate the file in openWriter() to make sure
	// the BeforeCopy didn't cancel the copy. If this was an existing
	// file that is not going to be resumed, truncate the contents.
<原文结束>

# <翻译开始>
// 我们在openWriter()函数中等待截断文件，是为了确保BeforeCopy不会取消复制操作。如果这是一个现有且不打算续传的文件，则截断其内容。
# <翻译结束>


<原文开始>
	// set file timestamp
<原文结束>

# <翻译开始>
// 设置文件时间戳
# <翻译结束>


<原文开始>
	// update transfer size if previously unknown
<原文结束>

# <翻译开始>
// 如果之前未知，则更新传输大小
# <翻译结束>


<原文开始>
	// run AfterCopy hook
<原文结束>

# <翻译开始>
// 运行 AfterCopy 钩子
# <翻译结束>


<原文开始>
// close finalizes the Response
<原文结束>

# <翻译开始>
// close 方法用于最终化（关闭）Response
# <翻译结束>

