
# <翻译开始>
type Request
X下载参数
# <翻译结束>

# <翻译开始>
Label string
X名称
<跳到行首>
# <翻译结束>

# <翻译开始>
HTTPRequest *http.Request
Http协议头
<跳到行首>
# <翻译结束>

# <翻译开始>
Filename string
X文件名
<跳到行首>
# <翻译结束>

# <翻译开始>
SkipExisting bool
X跳过已存在文件
<跳到行首>
# <翻译结束>

# <翻译开始>
NoResume bool
X不续传
<跳到行首>
# <翻译结束>

# <翻译开始>
NoStore bool
X不写入本地文件系统
<跳到行首>
# <翻译结束>

# <翻译开始>
NoCreateDirectories bool
X不自动创建目录
<跳到行首>
# <翻译结束>

# <翻译开始>
IgnoreBadStatusCodes bool
X忽略错误状态码
<跳到行首>
# <翻译结束>

# <翻译开始>
IgnoreRemoteTime bool
X忽略远程时间
<跳到行首>
# <翻译结束>

# <翻译开始>
Size int64
X预期文件大小
<跳到行首>
# <翻译结束>

# <翻译开始>
BufferSize int
X缓冲区大小
<跳到行首>
# <翻译结束>

# <翻译开始>
RateLimiter RateLimiter
X速率限制器
<跳到行首>
# <翻译结束>

# <翻译开始>
BeforeCopy Hook
X传输开始之前回调
<跳到行首>
# <翻译结束>

# <翻译开始>
AfterCopy Hook
X传输完成之后回调
<跳到行首>
# <翻译结束>

# <翻译开始>
func NewRequest(dst, urlStr
下载链接
# <翻译结束>

# <翻译开始>
func NewRequest(dst
保存目录
# <翻译结束>

# <翻译开始>
func NewRequest
X生成下载参数
# <翻译结束>

# <翻译开始>
) Context
I取上下文
# <翻译结束>

# <翻译开始>
) URL
X取下载链接
# <翻译结束>

# <翻译开始>
) SetChecksum(h hash.Hash, sum []byte, deleteOnError
出错后删除
# <翻译结束>

# <翻译开始>
) SetChecksum
X设置完成后效验
# <翻译结束>
