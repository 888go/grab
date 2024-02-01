
<原文开始>
// setLastModified sets the last modified timestamp of a local file according to
// the Last-Modified header returned by a remote server.
<原文结束>

# <翻译开始>
// setLastModified 根据远程服务器返回的 Last-Modified 头部信息，设置本地文件的最后修改时间戳。
# <翻译结束>


<原文开始>
	// https://tools.ietf.org/html/rfc7232#section-2.2
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
<原文结束>

# <翻译开始>
// 参考RFC 7232文档第2.2章节
// 及MDN Web文档中关于HTTP Headers的Last-Modified部分
// RFC 7232是HTTP/1.1协议中定义条件请求语义的标准，其中第2.2节详细说明了Last-Modified头部字段的用法。
// MDN Web文档（Mozilla开发者网络）对HTTP Headers中的Last-Modified做了详细介绍，这是一个用于指示资源最后修改时间的HTTP头部字段。
# <翻译结束>


<原文开始>
// mkdirp creates all missing parent directories for the destination file path.
<原文结束>

# <翻译开始>
// mkdirp 为目标文件路径创建所有缺失的父级目录。
# <翻译结束>


<原文开始>
// guessFilename returns a filename for the given http.Response. If none can be
// determined ErrNoFilename is returned.
//
// TODO: NoStore operations should not require a filename
<原文结束>

# <翻译开始>
// guessFilename 为给定的http.Response返回一个文件名。如果无法确定文件名，则返回ErrNoFilename错误。
//
// TODO: 对于NoStore操作，不应要求提供文件名
# <翻译结束>


<原文开始>
// else filename directive is missing.. fallback to URL.Path
<原文结束>

# <翻译开始>
// 如果filename指令缺失，则退回到URL.Path
# <翻译结束>

