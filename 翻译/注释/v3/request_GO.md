
<原文开始>
// A Hook is a user provided callback function that can be called by grab at
// various stages of a requests lifecycle. If a hook returns an error, the
// associated request is canceled and the same error is returned on the Response
// object.
//
// Hook functions are called synchronously and should never block unnecessarily.
// Response methods that block until a download is complete, such as
// Response.Err, Response.Cancel or Response.Wait will deadlock. To cancel a
// download from a callback, simply return a non-nil error.
<原文结束>

# <翻译开始>
// Hook 是用户提供的回调函数，可以在请求生命周期的各个阶段由grab调用。
// 若Hook返回错误，关联的请求将被取消，并且相同的错误会在Response对象上返回。
//
// Hook函数是同步调用的，绝不应无故阻塞。响应方法（如Response.Err、Response.Cancel或Response.Wait）在等待下载完成时会阻塞，因此可能会导致死锁。若要在回调中取消下载，只需返回非空错误即可。
# <翻译结束>


<原文开始>
// A Request represents an HTTP file transfer request to be sent by a Client.
<原文结束>

# <翻译开始>
// Request代表了一个由Client发送的HTTP文件传输请求。
# <翻译结束>


<原文开始>
	// Label is an arbitrary string which may used to label a Request with a
	// user friendly name.
<原文结束>

# <翻译开始>
// Label 是一个任意字符串，可用于为 Request 指定一个用户友好的名称。
# <翻译结束>


<原文开始>
	// Tag is an arbitrary interface which may be used to relate a Request to
	// other data.
<原文结束>

# <翻译开始>
// Tag 是一个任意接口，可用于将 Request 与其他数据关联起来。
# <翻译结束>


<原文开始>
	// HTTPRequest specifies the http.Request to be sent to the remote server to
	// initiate a file transfer. It includes request configuration such as URL,
	// protocol version, HTTP method, request headers and authentication.
<原文结束>

# <翻译开始>
// HTTPRequest 指定了要发送到远程服务器以启动文件传输的 http.Request。它包括请求配置，如 URL、协议版本、HTTP 方法、请求头和身份验证。
# <翻译结束>


<原文开始>
	// Filename specifies the path where the file transfer will be stored in
	// local storage. If Filename is empty or a directory, the true Filename will
	// be resolved using Content-Disposition headers or the request URL.
	//
	// An empty string means the transfer will be stored in the current working
	// directory.
<原文结束>

# <翻译开始>
// Filename 指定文件传输将在本地存储中保存的路径。如果 Filename 为空或是一个目录，
// 则会通过 Content-Disposition 头信息或请求 URL 来解析真实的 Filename。
//
// 空字符串表示该传输将在当前工作目录下存储。
# <翻译结束>


<原文开始>
	// SkipExisting specifies that ErrFileExists should be returned if the
	// destination path already exists. The existing file will not be checked for
	// completeness.
<原文结束>

# <翻译开始>
// SkipExisting 指定当目标路径已存在时应返回 ErrFileExists 错误。如果目标文件已经存在，
// 将不会检查其完整性。
// （注：ErrFileExists 是一个自定义错误类型，表示文件已存在；此选项适用于在复制或移动文件等操作中，遇到同名文件时直接跳过或返回错误，而不是覆盖原有文件。）
# <翻译结束>


<原文开始>
	// NoResume specifies that a partially completed download will be restarted
	// without attempting to resume any existing file. If the download is already
	// completed in full, it will not be restarted.
<原文结束>

# <翻译开始>
// NoResume 指定如果一个部分完成的下载将会在不尝试续传任何现有文件的情况下重新开始。但如果下载已经完全完成，则不会重新启动。
# <翻译结束>


<原文开始>
	// NoStore specifies that grab should not write to the local file system.
	// Instead, the download will be stored in memory and accessible only via
	// Response.Open or Response.Bytes.
<原文结束>

# <翻译开始>
// NoStore 指定 grab 应当不写入本地文件系统。
// 取而代之的是，下载内容将存储在内存中，并且只能通过 Response.Open 或 Response.Bytes 来访问。
# <翻译结束>


<原文开始>
	// NoCreateDirectories specifies that any missing directories in the given
	// Filename path should not be created automatically, if they do not already
	// exist.
<原文结束>

# <翻译开始>
// NoCreateDirectories 指定在给定的 Filename 路径中，如果不存在任何缺失的目录，则不应自动创建。
# <翻译结束>


<原文开始>
	// IgnoreBadStatusCodes specifies that grab should accept any status code in
	// the response from the remote server. Otherwise, grab expects the response
	// status code to be within the 2XX range (after following redirects).
<原文结束>

# <翻译开始>
// IgnoreBadStatusCodes 指定 grab 应在接受来自远程服务器响应中的任何状态码。否则，默认情况下，grab 期望响应状态码在重定向后位于 2XX 范围内。
# <翻译结束>


<原文开始>
	// IgnoreRemoteTime specifies that grab should not attempt to set the
	// timestamp of the local file to match the remote file.
<原文结束>

# <翻译开始>
// IgnoreRemoteTime 指定 grab 应该不尝试将本地文件的时间戳设置为与远程文件匹配。
# <翻译结束>


<原文开始>
	// Size specifies the expected size of the file transfer if known. If the
	// server response size does not match, the transfer is cancelled and
	// ErrBadLength returned.
<原文结束>

# <翻译开始>
// Size 指定预期的文件传输大小（如果已知）。如果服务器响应的大小不匹配，则取消传输并返回 ErrBadLength 错误。
# <翻译结束>


<原文开始>
	// BufferSize specifies the size in bytes of the buffer that is used for
	// transferring the requested file. Larger buffers may result in faster
	// throughput but will use more memory and result in less frequent updates
	// to the transfer progress statistics. If a RateLimiter is configured,
	// BufferSize should be much lower than the rate limit. Default: 32KB.
<原文结束>

# <翻译开始>
// BufferSize 指定用于传输请求文件的缓冲区大小（以字节为单位）。更大的缓冲区可能会带来更快的数据传输速度，但会占用更多的内存，并导致传输进度统计信息更新频率降低。如果配置了速率限制器（RateLimiter），则应将 BufferSize 设置得远低于速率限制值。默认值：32KB。
# <翻译结束>


<原文开始>
	// RateLimiter allows the transfer rate of a download to be limited. The given
	// Request.BufferSize determines how frequently the RateLimiter will be
	// polled.
<原文结束>

# <翻译开始>
// RateLimiter 允许限制下载的传输速率。给定的
// Request.BufferSize 决定了 RateLimiter 被调用查询的频率。
// （注：这段代码注释描述了一个名为 RateLimiter 的结构或函数，它的功能是限制数据下载的速度。RateLimiter 的工作频率由传入的 Request 结构体中的 BufferSize 参数决定。）
# <翻译结束>


<原文开始>
	// BeforeCopy is a user provided callback that is called immediately before
	// a request starts downloading. If BeforeCopy returns an error, the request
	// is cancelled and the same error is returned on the Response object.
<原文结束>

# <翻译开始>
// BeforeCopy 是用户提供的回调函数，在请求开始下载内容前立即调用。如果 BeforeCopy 返回错误，则取消该请求，并在 Response 对象上返回相同的错误。
# <翻译结束>


<原文开始>
	// AfterCopy is a user provided callback that is called immediately after a
	// request has finished downloading, before checksum validation and closure.
	// This hook is only called if the transfer was successful. If AfterCopy
	// returns an error, the request is canceled and the same error is returned on
	// the Response object.
<原文结束>

# <翻译开始>
// AfterCopy 是用户提供的回调函数，在请求完成下载后立即调用，但在校验checksum和关闭连接之前。
// 该钩子仅在传输成功时被调用。如果AfterCopy返回错误，则请求将被取消，并且相同的错误会在Response对象上返回。
# <翻译结束>


<原文开始>
	// hash, checksum and deleteOnError - set via SetChecksum.
<原文结束>

# <翻译开始>
// hash, checksum 和 deleteOnError — 通过 SetChecksum 方法设置。
# <翻译结束>


<原文开始>
	// Context for cancellation and timeout - set via WithContext
<原文结束>

# <翻译开始>
// Context用于取消和超时控制，通过WithContext方法设置
# <翻译结束>


<原文开始>
// NewRequest returns a new file transfer Request suitable for use with
// Client.Do.
<原文结束>

# <翻译开始>
// NewRequest 返回一个新的文件传输请求，适合用于
// Client.Do 方法。
# <翻译结束>


<原文开始>
// Context returns the request's context. To change the context, use
// WithContext.
//
// The returned context is always non-nil; it defaults to the background
// context.
//
// The context controls cancelation.
<原文结束>

# <翻译开始>
// Context返回请求的上下文。若要更改上下文，请使用
// WithContext。
//
// 返回的上下文始终是非空的，默认为后台上下文。
//
// 上下文控制取消操作。
# <翻译结束>


<原文开始>
// WithContext returns a shallow copy of r with its context changed
// to ctx. The provided ctx must be non-nil.
<原文结束>

# <翻译开始>
// WithContext 返回 r 的浅复制副本，并将其上下文更改为 ctx。提供的 ctx 必须是非空的。
# <翻译结束>


<原文开始>
// URL returns the URL to be downloaded.
<原文结束>

# <翻译开始>
// URL 返回将要下载的URL。
# <翻译结束>


<原文开始>
// SetChecksum sets the desired hashing algorithm and checksum value to validate
// a downloaded file. Once the download is complete, the given hashing algorithm
// will be used to compute the actual checksum of the downloaded file. If the
// checksums do not match, an error will be returned by the associated
// Response.Err method.
//
// If deleteOnError is true, the downloaded file will be deleted automatically
// if it fails checksum validation.
//
// To prevent corruption of the computed checksum, the given hash must not be
// used by any other request or goroutines.
//
// To disable checksum validation, call SetChecksum with a nil hash.
<原文结束>

# <翻译开始>
// SetChecksum 设置期望的哈希算法和校验值以验证下载文件。
// 下载完成后，将使用给定的哈希算法计算下载文件的实际校验值。
// 如果校验值不匹配，则通过关联的 Response.Err 方法返回错误。
// 如果 deleteOnError 为 true，则在文件校验失败时，下载的文件将被自动删除。
// 为了防止计算出的校验值遭到破坏，给定的哈希必须未被任何其他请求或协程使用。
// 要禁用校验和验证，请使用 nil 哈希调用 SetChecksum。
# <翻译结束>

