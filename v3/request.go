package 下载类

import (
	"context"
	"hash"
	"net/http"
	"net/url"
)

// Hook 是用户提供的回调函数，可以在请求生命周期的各个阶段由grab调用。
// 若Hook返回错误，关联的请求将被取消，并且相同的错误会在Response对象上返回。
//
// Hook函数是同步调用的，绝不应无故阻塞。响应方法（如Response.Err、Response.Cancel或Response.Wait）在等待下载完成时会阻塞，因此可能会导致死锁。若要在回调中取消下载，只需返回非空错误即可。
type Hook func(*X响应) error

// Request代表了一个由Client发送的HTTP文件传输请求。
type X下载参数 struct {
// Label 是一个任意字符串，可用于为 Request 指定一个用户友好的名称。
	X名称 string

// Tag 是一个任意接口，可用于将 Request 与其他数据关联起来。
	Tag interface{}

// HTTPRequest 指定了要发送到远程服务器以启动文件传输的 http.Request。它包括请求配置，如 URL、协议版本、HTTP 方法、请求头和身份验证。
	Http协议头 *http.Request

// Filename 指定文件传输将在本地存储中保存的路径。如果 Filename 为空或是一个目录，
// 则会通过 Content-Disposition 头信息或请求 URL 来解析真实的 Filename。
//
// 空字符串表示该传输将在当前工作目录下存储。
	X文件名 string

// SkipExisting 指定当目标路径已存在时应返回 ErrFileExists 错误。如果目标文件已经存在，
// 将不会检查其完整性。
// （注：ErrFileExists 是一个自定义错误类型，表示文件已存在；此选项适用于在复制或移动文件等操作中，遇到同名文件时直接跳过或返回错误，而不是覆盖原有文件。）
	X跳过已存在文件 bool

// NoResume 指定如果一个部分完成的下载将会在不尝试续传任何现有文件的情况下重新开始。但如果下载已经完全完成，则不会重新启动。
	X不续传 bool

// NoStore 指定 grab 应当不写入本地文件系统。
// 取而代之的是，下载内容将存储在内存中，并且只能通过 Response.Open 或 Response.Bytes 来访问。
	X不写入本地文件系统 bool

// NoCreateDirectories 指定在给定的 Filename 路径中，如果不存在任何缺失的目录，则不应自动创建。
	X不自动创建目录 bool

// IgnoreBadStatusCodes 指定 grab 应在接受来自远程服务器响应中的任何状态码。否则，默认情况下，grab 期望响应状态码在重定向后位于 2XX 范围内。
	X忽略错误状态码 bool

// IgnoreRemoteTime 指定 grab 应该不尝试将本地文件的时间戳设置为与远程文件匹配。
	X忽略远程时间 bool

// Size 指定预期的文件传输大小（如果已知）。如果服务器响应的大小不匹配，则取消传输并返回 ErrBadLength 错误。
	X预期文件大小 int64

// BufferSize 指定用于传输请求文件的缓冲区大小（以字节为单位）。更大的缓冲区可能会带来更快的数据传输速度，但会占用更多的内存，并导致传输进度统计信息更新频率降低。如果配置了速率限制器（RateLimiter），则应将 BufferSize 设置得远低于速率限制值。默认值：32KB。
	X缓冲区大小 int

// RateLimiter 允许限制下载的传输速率。给定的
// Request.BufferSize 决定了 RateLimiter 被调用查询的频率。
// （注：这段代码注释描述了一个名为 RateLimiter 的结构或函数，它的功能是限制数据下载的速度。RateLimiter 的工作频率由传入的 Request 结构体中的 BufferSize 参数决定。）
	X速率限制器 RateLimiter

// BeforeCopy 是用户提供的回调函数，在请求开始下载内容前立即调用。如果 BeforeCopy 返回错误，则取消该请求，并在 Response 对象上返回相同的错误。
	X传输开始之前回调 Hook

// AfterCopy 是用户提供的回调函数，在请求完成下载后立即调用，但在校验checksum和关闭连接之前。
// 该钩子仅在传输成功时被调用。如果AfterCopy返回错误，则请求将被取消，并且相同的错误会在Response对象上返回。
	X传输完成之后回调 Hook

// hash, checksum 和 deleteOnError — 通过 SetChecksum 方法设置。
	hash          hash.Hash
	checksum      []byte
	deleteOnError bool

// Context用于取消和超时控制，通过WithContext方法设置
	ctx context.Context
}

// NewRequest 返回一个新的文件传输请求，适合用于
// Client.Do 方法。
func X生成下载参数(保存目录, 下载链接 string) (*X下载参数, error) {
	if 保存目录 == "" {
		保存目录 = "."
	}
	req, err := http.NewRequest("GET", 下载链接, nil)
	if err != nil {
		return nil, err
	}
	return &X下载参数{
		Http协议头: req,
		X文件名:    保存目录,
	}, nil
}

// Context返回请求的上下文。若要更改上下文，请使用
// WithContext。
//
// 返回的上下文始终是非空的，默认为后台上下文。
//
// 上下文控制取消操作。
func (r *X下载参数) I取上下文() context.Context {
	if r.ctx != nil {
		return r.ctx
	}

	return context.Background()
}

// WithContext 返回 r 的浅复制副本，并将其上下文更改为 ctx。提供的 ctx 必须是非空的。
func (r *X下载参数) WithContext(ctx context.Context) *X下载参数 {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(X下载参数)
	*r2 = *r
	r2.ctx = ctx
	r2.Http协议头 = r2.Http协议头.WithContext(ctx)
	return r2
}

// URL 返回将要下载的URL。
func (r *X下载参数) X取下载链接() *url.URL {
	return r.Http协议头.URL
}

// SetChecksum 设置期望的哈希算法和校验值以验证下载文件。
// 下载完成后，将使用给定的哈希算法计算下载文件的实际校验值。
// 如果校验值不匹配，则通过关联的 Response.Err 方法返回错误。
// 如果 deleteOnError 为 true，则在文件校验失败时，下载的文件将被自动删除。
// 为了防止计算出的校验值遭到破坏，给定的哈希必须未被任何其他请求或协程使用。
// 要禁用校验和验证，请使用 nil 哈希调用 SetChecksum。
func (r *X下载参数) X设置完成后效验(h hash.Hash, sum []byte, 出错后删除 bool) {
	r.hash = h
	r.checksum = sum
	r.deleteOnError = 出错后删除
}
