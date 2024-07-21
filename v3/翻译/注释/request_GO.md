
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
// Hook 是用户提供的回调函数，可以在请求生命周期的各个阶段由 grab 调用。
// 如果钩子函数返回错误，关联的请求将被取消，并且相同的错误将返回在 Response 对象上。
//
// 钩子函数是同步调用的，应避免不必要的阻塞。
// 诸如 Response.Err、Response.Cancel 或 Response.Wait 这些会阻塞直到下载完成的 Response 方法会导致死锁。
// 若要从回调函数中取消下载，只需返回一个非空错误。
// md5:b4013e6bf57cee77
# <翻译结束>


<原文开始>
// A Request represents an HTTP file transfer request to be sent by a Client.
<原文结束>

# <翻译开始>
// Request表示客户端要发送的HTTP文件传输请求。 md5:3bbb6f4b848425e8
# <翻译结束>


<原文开始>
	// Label is an arbitrary string which may used to label a Request with a
	// user friendly name.
<原文结束>

# <翻译开始>
	// Label是一个任意字符串，可以用来将Request标记为一个用户友好的名称。
	// md5:48c76f58a675938f
# <翻译结束>


<原文开始>
	// Tag is an arbitrary interface which may be used to relate a Request to
	// other data.
<原文结束>

# <翻译开始>
	// Tag 是一个任意的接口，可以用来将请求与其他数据相关联。
	// md5:cca083a2da336d7a
# <翻译结束>


<原文开始>
	// HTTPRequest specifies the http.Request to be sent to the remote server to
	// initiate a file transfer. It includes request configuration such as URL,
	// protocol version, HTTP method, request headers and authentication.
<原文结束>

# <翻译开始>
	// HTTPRequest 指定了要发送到远程服务器以启动文件传输的HTTP请求。它包括URL、协议版本、HTTP方法、请求头和身份验证等请求配置信息。
	// md5:8b29a36ebccc4fbf
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
	// Filename 指定了文件传输将在本地存储的位置。如果 Filename 为空或是一个目录，将使用 Content-Disposition 头部或请求 URL 来解析实际的文件名。
	// 
	// 空字符串表示文件传输将存储在当前工作目录下。
	// md5:b2487dc28865af4e
# <翻译结束>


<原文开始>
	// SkipExisting specifies that ErrFileExists should be returned if the
	// destination path already exists. The existing file will not be checked for
	// completeness.
<原文结束>

# <翻译开始>
	// SkipExisting 指定如果目标路径已经存在，应该返回 ErrFileExists 错误。不会检查已存在的文件是否完整。
	// md5:fcfcc6296df94857
# <翻译结束>


<原文开始>
	// NoResume specifies that a partially completed download will be restarted
	// without attempting to resume any existing file. If the download is already
	// completed in full, it will not be restarted.
<原文结束>

# <翻译开始>
	// NoResume 指定在不尝试恢复任何现有文件的情况下，部分下载将被重新启动。如果下载已经完整完成，它将不会被重新启动。
	// md5:5d7097b2011a8aed
# <翻译结束>


<原文开始>
	// NoStore specifies that grab should not write to the local file system.
	// Instead, the download will be stored in memory and accessible only via
	// Response.Open or Response.Bytes.
<原文结束>

# <翻译开始>
	// NoStore 指定 grab 不应将内容写入本地文件系统。相反，下载的内容将存储在内存中，仅可通过 Response.Open 或 Response.Bytes 访问。
	// md5:b250d7f0915c2e35
# <翻译结束>


<原文开始>
	// NoCreateDirectories specifies that any missing directories in the given
	// Filename path should not be created automatically, if they do not already
	// exist.
<原文结束>

# <翻译开始>
	// NoCreateDirectories 指定在给定的Filename路径中，任何缺失的目录不应该自动创建，如果它们尚不存在的话。
	// md5:f9491a6bdafa1a03
# <翻译结束>


<原文开始>
	// IgnoreBadStatusCodes specifies that grab should accept any status code in
	// the response from the remote server. Otherwise, grab expects the response
	// status code to be within the 2XX range (after following redirects).
<原文结束>

# <翻译开始>
	// IgnoreBadStatusCodes 指定grab应该接受远程服务器返回的任何状态码。否则，grab期望响应状态码在200（跟随重定向后）范围内。
	// md5:0cfc61395f3c8fda
# <翻译结束>


<原文开始>
	// IgnoreRemoteTime specifies that grab should not attempt to set the
	// timestamp of the local file to match the remote file.
<原文结束>

# <翻译开始>
	// IgnoreRemoteTime 指定 grab 不应尝试将本地文件的时间戳设置为与远程文件匹配。
	// md5:4b121c28c096e635
# <翻译结束>


<原文开始>
	// Size specifies the expected size of the file transfer if known. If the
	// server response size does not match, the transfer is cancelled and
	// ErrBadLength returned.
<原文结束>

# <翻译开始>
	// Size 指定如果已知的话，文件传输的预期大小。如果服务器响应的大小与预期不符， 
	// 传输将被取消并返回 ErrBadLength 错误。
	// md5:4d6feeb012ab2e19
# <翻译结束>


<原文开始>
	// BufferSize specifies the size in bytes of the buffer that is used for
	// transferring the requested file. Larger buffers may result in faster
	// throughput but will use more memory and result in less frequent updates
	// to the transfer progress statistics. If a RateLimiter is configured,
	// BufferSize should be much lower than the rate limit. Default: 32KB.
<原文结束>

# <翻译开始>
	// BufferSize 指定了用于传输请求文件的缓冲区的字节数。更大的缓冲区可能会提高吞吐量，但会消耗更多内存，并减少传输进度统计信息的更新频率。如果配置了速率限制器，BufferSize 应远低于速率限制。默认值：32KB。
	// md5:c943b23cdf634593
# <翻译结束>


<原文开始>
	// RateLimiter allows the transfer rate of a download to be limited. The given
	// Request.BufferSize determines how frequently the RateLimiter will be
	// polled.
<原文结束>

# <翻译开始>
	// RateLimiter 用于限制下载的传输速率。给定的 Request.BufferSize 决定了RateLimiter将多久被查询一次。
	// md5:86949911252236b7
# <翻译结束>


<原文开始>
	// BeforeCopy is a user provided callback that is called immediately before
	// a request starts downloading. If BeforeCopy returns an error, the request
	// is cancelled and the same error is returned on the Response object.
<原文结束>

# <翻译开始>
	// BeforeCopy 是用户提供的回调函数，它在请求开始下载之前立即被调用。如果 BeforeCopy 返回错误，请求将被取消，并且相同的错误将返回在响应对象中。
	// md5:865a949d97572fb2
# <翻译结束>


<原文开始>
	// AfterCopy is a user provided callback that is called immediately after a
	// request has finished downloading, before checksum validation and closure.
	// This hook is only called if the transfer was successful. If AfterCopy
	// returns an error, the request is canceled and the same error is returned on
	// the Response object.
<原文结束>

# <翻译开始>
	// AfterCopy 是用户提供的回调，它在请求下载完成后立即调用，但在校验和验证和关闭之前。此钩子仅在传输成功时调用。如果 AfterCopy 返回错误，请求将被取消，并在 Response 对象上返回相同的错误。
	// md5:9e8a840d6dbbf503
# <翻译结束>


<原文开始>
// hash, checksum and deleteOnError - set via SetChecksum.
<原文结束>

# <翻译开始>
// hash、checksum和deleteOnError - 通过SetChecksum设置。 md5:cc1ef74b5265e5bb
# <翻译结束>


<原文开始>
// Context for cancellation and timeout - set via WithContext
<原文结束>

# <翻译开始>
// 用于取消和超时的上下文 - 通过WithContext设置 md5:b3765a4b0c8785c9
# <翻译结束>


<原文开始>
// NewRequest returns a new file transfer Request suitable for use with
// Client.Do.
<原文结束>

# <翻译开始>
// NewRequest 返回一个新的文件传输请求，适合与 Client.Do 使用。
// md5:1bf325ab23123d38
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
// Context 返回请求的上下文。要更改上下文，请使用
// WithContext。
//
// 返回的上下文始终非空；它默认为背景上下文。
//
// 上下文控制取消操作。
// md5:48c4bb357919bed1
# <翻译结束>


<原文开始>
// WithContext returns a shallow copy of r with its context changed
// to ctx. The provided ctx must be non-nil.
<原文结束>

# <翻译开始>
// WithContext 返回一个与 r 拥有相同结构但上下文改为 ctx 的浅拷贝。提供的 ctx 必须非 nil。
// md5:717222e0b2288cb8
# <翻译结束>


<原文开始>
// URL returns the URL to be downloaded.
<原文结束>

# <翻译开始>
// URL 返回要下载的 URL。 md5:aebf141338bb627e
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
// SetChecksum 设置期望的哈希算法和校验值，用于验证下载的文件。下载完成后，将使用给定的哈希算法计算实际下载文件的校验和。如果校验和不匹配，关联的Response.Err方法将返回错误。
// 
// 如果deleteOnError为true，如果校验和验证失败，下载的文件将自动删除。
//
// 为了防止计算出的校验和被破坏，提供的哈希值在其他请求或goroutine中不应被使用。
//
// 要禁用校验和验证，请使用nil哈希值调用SetChecksum。
// md5:56e3d06ddceed502
# <翻译结束>

