
<原文开始>
// HTTPClient provides an interface allowing us to perform HTTP requests.
<原文结束>

# <翻译开始>
// HTTPClient 提供了一个接口，允许我们执行HTTP请求。 md5:4171bffd13d00b3a
# <翻译结束>


<原文开始>
// truncater is a private interface allowing different response
// Writers to be truncated
<原文结束>

# <翻译开始>
// truncater是一个私有接口，允许使用不同的响应写入器进行截断
// md5:13b36272ac021c7d
# <翻译结束>


<原文开始>
// A Client is a file download client.
//
// Clients are safe for concurrent use by multiple goroutines.
<原文结束>

# <翻译开始>
// Client是一个文件下载客户端。
//
// 客户端是安全的，可以由多个goroutine并发使用。
// md5:adb110330920265b
# <翻译结束>


<原文开始>
	// HTTPClient specifies the http.Client which will be used for communicating
	// with the remote server during the file transfer.
<原文结束>

# <翻译开始>
// HTTPClient 指定了在文件传输过程中与远程服务器通信所使用的 http.Client。
// md5:18f90ae2438c3d95
# <翻译结束>


<原文开始>
	// UserAgent specifies the User-Agent string which will be set in the
	// headers of all requests made by this client.
	//
	// The user agent string may be overridden in the headers of each request.
<原文结束>

# <翻译开始>
// UserAgent 指定将设置为此客户端的所有请求的头中的 User-Agent 字符串。
//
// 可以在每个请求的头中覆盖 User-Agent 字符串。
// md5:416ff171bdbaf067
# <翻译结束>


<原文开始>
	// BufferSize specifies the size in bytes of the buffer that is used for
	// transferring all requested files. Larger buffers may result in faster
	// throughput but will use more memory and result in less frequent updates
	// to the transfer progress statistics. The BufferSize of each request can
	// be overridden on each Request object. Default: 32KB.
<原文结束>

# <翻译开始>
// BufferSize 指定了用于传输所有请求文件的缓冲区的大小（以字节为单位）。较大的缓冲区可能会提高吞吐量，但会消耗更多内存，并导致传输进度统计更新更不频繁。每个请求的BufferSize可以在每个Request对象上重写。默认值：32KB。
// md5:76d8523deca24178
# <翻译结束>


<原文开始>
// NewClient returns a new file download Client, using default configuration.
<原文结束>

# <翻译开始>
// NewClient 使用默认配置返回一个新的文件下载客户端。 md5:37b6062ca04f5dcf
# <翻译结束>


<原文开始>
// DefaultClient is the default client and is used by all Get convenience
// functions.
<原文结束>

# <翻译开始>
// DefaultClient 是默认客户端，所有 Get 方便函数都会使用它。
// md5:3aee22c6e017f0f3
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
// Do 发送文件传输请求并返回一个文件传输响应，遵循客户端的HTTPClient配置（例如重定向、Cookie、身份验证等）。
// 
// 类似于http.Get，Do在启动传输时会阻塞，但一旦传输在后台goroutine中开始传输，或者在早期阶段失败，它就会立即返回。
// 
// 如果错误是由客户端策略（如CheckRedirect）引起的，或者存在HTTP协议或IO错误，将通过Response.Err返回错误。Response.Err会在传输完成，无论成功还是失败，都会阻塞调用者。
// md5:45eff2a029f7d8e0
# <翻译结束>


<原文开始>
// cancel will be called on all code-paths via closeResponse
<原文结束>

# <翻译开始>
// cancel 函数会通过 closeResponse 在所有代码路径上被调用 md5:e670a9d606a72b8c
# <翻译结束>


<原文开始>
// default to Client.BufferSize
<原文结束>

# <翻译开始>
// 默认为Client.BufferSize md5:f7214bed22611823
# <翻译结束>


<原文开始>
	// Run state-machine while caller is blocked to initialize the file transfer.
	// Must never transition to the copyFile state - this happens next in another
	// goroutine.
<原文结束>

# <翻译开始>
// 当调用者被阻塞以初始化文件传输时，运行状态机。永远不要进入copyFile状态 - 这将在另一个goroutine中接下来发生。
// md5:c59c8fd584f98402
# <翻译结束>


<原文开始>
	// Run copyFile in a new goroutine. copyFile will no-op if the transfer is
	// already complete or failed.
<原文结束>

# <翻译开始>
// 在新的goroutine中运行copyFile。如果传输已经完成或失败，copyFile将不执行任何操作。
// md5:0fe588a8a89f07a6
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
// DoChannel 通过给定的 Request 通道顺序执行所有发送的请求，直到由其他goroutine关闭。调用者会在Request通道关闭且所有传输完成后被阻塞。一旦从远程服务器接收到响应，所有响应会立即通过给定的Response通道发送，并可用于跟踪每个下载的进度。
// 
// 如果Response接收器处理速度慢，会导致worker阻塞，从而延迟已经启动连接的传输开始时间——可能导致服务器超时。调用者需要确保为Response通道使用足够的缓冲大小以防止这种情况。
// 
// 在任何文件传输过程中如果发生错误，可以通过关联的Response.Err函数访问到该错误。
// md5:1e8b3d5e8dee8a95
# <翻译结束>


<原文开始>
// TODO: enable cancelling of batch jobs
<原文结束>

# <翻译开始>
// TODO：启用批量任务的取消功能 md5:551d7d31d87cb53d
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
// DoBatch 使用指定数量的并发工作者执行所有给定的请求。一旦工作者启动，控制权就会返回给调用者。
//
// 如果请求的工作者数量小于1，那么将为每个请求创建一个工作者。即，所有请求将并行执行。
//
// 如果在任何文件传输过程中发生错误，可以通过调用关联的Response.Err来获取该错误。
//
// 返回的Response通道只会在所有给定的Requests完成（无论成功或失败）后关闭。
// md5:a143d0d000672213
# <翻译结束>


<原文开始>
// An stateFunc is an action that mutates the state of a Response and returns
// the next stateFunc to be called.
<原文结束>

# <翻译开始>
// StateFunc是一个动作，它会改变Response的状态，并返回下一个要调用的StateFunc。
// md5:2f895dfc8babb8bb
# <翻译结束>


<原文开始>
// run calls the given stateFunc function and all subsequent returned stateFuncs
// until a stateFunc returns nil or the Response.ctx is canceled. Each stateFunc
// should mutate the state of the given Response until it has completed
// downloading or failed.
<原文结束>

# <翻译开始>
// run 调用给定的 stateFunc 函数，以及所有后续返回的 stateFuncs，直到某个 stateFunc 返回 nil 或者 Response 的 ctx 被取消。每个 stateFunc 应该修改给定 Response 的状态，直到下载完成或失败。
// md5:dcde558208dbf3d5
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
// statFileInfo 获取与Response.Filename匹配的任何本地文件的FileInfo。
//
// 如果文件不存在，是一个目录，或者其名称未知，则下一个状态函数为headRequest。
//
// 如果文件存在，则设置Response.fi，下一个状态函数为validateLocal。
//
// 如果发生错误，下一个状态函数为closeResponse。
// md5:6285457c42618336
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
// validateLocal 将下载文件的本地副本与远程文件进行比较。
// 
// 如果本地文件大于远程文件，或者Request.SkipExisting 为 true，则返回错误。
// 
// 如果现有文件的大小与远程文件匹配，下一个状态函数是checksumFile。
// 
// 如果本地文件小于远程文件，并且已知远程服务器支持范围请求，下一个状态函数是getRequest。
// md5:fb339cef020987be
# <翻译结束>


<原文开始>
// determine target file size
<原文结束>

# <翻译开始>
// 确定目标文件大小 md5:b64a2f6b674bfe35
# <翻译结束>


<原文开始>
		// size is either actually 0 or unknown
		// if unknown, we ask the remote server
		// if known to be 0, we proceed with a GET
<原文结束>

# <翻译开始>
// size 是实际大小，可能是0或未知
// 如果未知，我们向远程服务器查询
// 如果已知为0，我们继续使用GET请求
// md5:b8aed7e2a19d006f
# <翻译结束>


<原文开始>
// local file matches remote file size - wrap it up
<原文结束>

# <翻译开始>
// 当本地文件的大小与远程文件相匹配时 - 封装它 md5:9020a4e52bdc768c
# <翻译结束>


<原文开始>
// local file should be overwritten
<原文结束>

# <翻译开始>
// 当地文件应被覆盖 md5:c6963052a01c5c41
# <翻译结束>


<原文开始>
// remote size is known, is smaller than local size and we want to resume
<原文结束>

# <翻译开始>
// 远程大小已知，比本地大小小，并且我们想要恢复下载 md5:4eff80977d3bd5ad
# <翻译结束>


<原文开始>
// set resume range on GET request
<原文结束>

# <翻译开始>
// 在GET请求中设置恢复范围 md5:baef5df94e869e30
# <翻译结束>


<原文开始>
// err should be os.PathError and include file path
<原文结束>

# <翻译开始>
// err应该是os.PathError类型，并且应包含文件路径 md5:255ad01285b97e27
# <翻译结束>


<原文开始>
// doHTTPRequest sends a HTTP Request and returns the response
<原文结束>

# <翻译开始>
// doHTTPRequest 发送一个HTTP请求并返回响应 md5:19242108c18d9170
# <翻译结束>


<原文开始>
// destination path is already known and does not exist
<原文结束>

# <翻译开始>
// 目标路径已知且不存在 md5:d19d7aa9a062a804
# <翻译结束>


<原文开始>
	// In case of redirects during HEAD, record the final URL and use it
	// instead of the original URL when sending future requests.
	// This way we avoid sending potentially unsupported requests to
	// the original URL, e.g. "Range", since it was the final URL
	// that advertised its support.
<原文结束>

# <翻译开始>
// 如果在HEAD请求期间发生重定向，记录最终URL，并在发送后续请求时使用它。
// 这样，我们就避免向原始URL发送可能不被支持的请求，例如"Range"，因为是最终URL表明了其支持。
// md5:5c169e71aa18b9cb
# <翻译结束>


<原文开始>
// TODO: check Content-Range
<原文结束>

# <翻译开始>
// 待办事项：检查 Content-Range md5:1dd4a612f45654be
# <翻译结束>


<原文开始>
// Request.Filename will be empty or a directory
<原文结束>

# <翻译开始>
// Request.Filename将会是空的或者是一个目录 md5:9577464c343c9907
# <翻译结束>


<原文开始>
// openWriter opens the destination file for writing and seeks to the location
// from whence the file transfer will resume.
//
// Requires that Response.Filename and resp.DidResume are already be set.
<原文结束>

# <翻译开始>
// openWriter 为写入打开目标文件，并定位到将继续传输的文件位置。
//
// 需要确保Response.Filename和resp.DidResume已经设置。
// md5:5670da9b8c85a4b1
# <翻译结束>


<原文开始>
				// truncate later in copyFile, if not cancelled
				// by BeforeCopy hook
<原文结束>

# <翻译开始>
// 如果BeforeCopy钩子没有取消，则稍后在copyFile中截断
// md5:8d0fcf3dbf9d4eb0
# <翻译结束>


<原文开始>
// next step is copyFile, but this will be called later in another goroutine
<原文结束>

# <翻译开始>
// 下一步是copyFile，但这个操作将在另一个goroutine中稍后被调用。 md5:bdcea988806d0db1
# <翻译结束>


<原文开始>
// copy transfers content for a HTTP connection established via Client.do()
<原文结束>

# <翻译开始>
// copy 将通过Client.do()建立的HTTP连接的内容进行传输 md5:feda3b303e54e72e
# <翻译结束>


<原文开始>
	// We waited to truncate the file in openWriter() to make sure
	// the BeforeCopy didn't cancel the copy. If this was an existing
	// file that is not going to be resumed, truncate the contents.
<原文结束>

# <翻译开始>
// 我们在openWriter()中推迟了截断文件的操作，以确保BeforeCopy不会取消复制。
// 如果这是一个现有的、不打算恢复的文件，那么截断其内容。
// md5:34b779d0e7dfa278
# <翻译结束>


<原文开始>
// update transfer size if previously unknown
<原文结束>

# <翻译开始>
// 如果之前未知，更新传输大小 md5:d285b51e526307dd
# <翻译结束>


<原文开始>
// close finalizes the Response
<原文结束>

# <翻译开始>
// close 对Response进行最终化处理 md5:5f125695e8212203
# <翻译结束>

