
<原文开始>
	// ErrBadLength indicates that the server response or an existing file does
	// not match the expected content length.
<原文结束>

# <翻译开始>
// ErrBadLength 表示服务器响应或已存在的文件内容长度与预期的不匹配。
# <翻译结束>


<原文开始>
	// ErrBadChecksum indicates that a downloaded file failed to pass checksum
	// validation.
<原文结束>

# <翻译开始>
// ErrBadChecksum 表示下载的文件未能通过校验和验证。
# <翻译结束>


<原文开始>
	// ErrNoFilename indicates that a reasonable filename could not be
	// automatically determined using the URL or response headers from a server.
<原文结束>

# <翻译开始>
// ErrNoFilename 表示无法通过 URL 或服务器响应头自动生成一个合理的文件名。
# <翻译结束>


<原文开始>
	// ErrNoTimestamp indicates that a timestamp could not be automatically
	// determined using the response headers from the remote server.
<原文结束>

# <翻译开始>
// ErrNoTimestamp 表示无法通过远程服务器响应头自动确定时间戳。
# <翻译结束>


<原文开始>
	// ErrFileExists indicates that the destination path already exists.
<原文结束>

# <翻译开始>
// ErrFileExists 表示目标路径已存在。
# <翻译结束>


<原文开始>
// StatusCodeError indicates that the server response had a status code that
// was not in the 200-299 range (after following any redirects).
<原文结束>

# <翻译开始>
// StatusCodeError 表示服务器响应的状态码不在 200-299 范围内（在跟随所有重定向之后）。
// 注：此注释描述了一个Go语言错误类型，当HTTP请求的最终状态码不在成功范围（200-299）内时，会返回这个错误。
# <翻译结束>


<原文开始>
// IsStatusCodeError returns true if the given error is of type StatusCodeError.
<原文结束>

# <翻译开始>
// IsStatusCodeError 判断给定的错误是否为 StatusCodeError 类型，如果是则返回 true。
# <翻译结束>

