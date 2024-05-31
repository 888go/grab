package 下载类

import (
	"context"
	"hash"
	"net/http"
	"net/url"
)

// Hook 是用户提供的回调函数，可以在请求生命周期的各个阶段由 grab 调用。
// 如果钩子函数返回错误，关联的请求将被取消，并且相同的错误将返回在 Response 对象上。
//
// 钩子函数是同步调用的，应避免不必要的阻塞。
// 诸如 Response.Err、Response.Cancel 或 Response.Wait 这些会阻塞直到下载完成的 Response 方法会导致死锁。
// 若要从回调函数中取消下载，只需返回一个非空错误。
// md5:b4013e6bf57cee77
type Hook func(*Response) error

// Request表示客户端要发送的HTTP文件传输请求。 md5:3bbb6f4b848425e8
type Request struct {
// Label是一个任意字符串，可以用来将Request标记为一个用户友好的名称。
// md5:48c76f58a675938f
	X名称 string

// Tag 是一个任意的接口，可以用来将请求与其他数据相关联。
// md5:cca083a2da336d7a
	Tag interface{}

// HTTPRequest 指定了要发送到远程服务器以启动文件传输的HTTP请求。它包括URL、协议版本、HTTP方法、请求头和身份验证等请求配置信息。
// md5:8b29a36ebccc4fbf
	Http协议头 *http.Request

// Filename 指定了文件传输将在本地存储的位置。如果 Filename 为空或是一个目录，将使用 Content-Disposition 头部或请求 URL 来解析实际的文件名。
// 
// 空字符串表示文件传输将存储在当前工作目录下。
// md5:b2487dc28865af4e
	X文件名 string

// SkipExisting 指定如果目标路径已经存在，应该返回 ErrFileExists 错误。不会检查已存在的文件是否完整。
// md5:fcfcc6296df94857
	X跳过已存在文件 bool

// NoResume 指定在不尝试恢复任何现有文件的情况下，部分下载将被重新启动。如果下载已经完整完成，它将不会被重新启动。
// md5:5d7097b2011a8aed
	X不续传 bool

// NoStore 指定 grab 不应将内容写入本地文件系统。相反，下载的内容将存储在内存中，仅可通过 Response.Open 或 Response.Bytes 访问。
// md5:b250d7f0915c2e35
	X不写入本地文件系统 bool

// NoCreateDirectories 指定在给定的Filename路径中，任何缺失的目录不应该自动创建，如果它们尚不存在的话。
// md5:f9491a6bdafa1a03
	X不自动创建目录 bool

// IgnoreBadStatusCodes 指定grab应该接受远程服务器返回的任何状态码。否则，grab期望响应状态码在200（跟随重定向后）范围内。
// md5:0cfc61395f3c8fda
	X忽略错误状态码 bool

// IgnoreRemoteTime 指定 grab 不应尝试将本地文件的时间戳设置为与远程文件匹配。
// md5:4b121c28c096e635
	X忽略远程时间 bool

// Size 指定如果已知的话，文件传输的预期大小。如果服务器响应的大小与预期不符， 
// 传输将被取消并返回 ErrBadLength 错误。
// md5:4d6feeb012ab2e19
	X预期文件大小 int64

// BufferSize 指定了用于传输请求文件的缓冲区的字节数。更大的缓冲区可能会提高吞吐量，但会消耗更多内存，并减少传输进度统计信息的更新频率。如果配置了速率限制器，BufferSize 应远低于速率限制。默认值：32KB。
// md5:c943b23cdf634593
	X缓冲区大小 int

// RateLimiter 用于限制下载的传输速率。给定的 Request.BufferSize 决定了RateLimiter将多久被查询一次。
// md5:86949911252236b7
	X速率限制器 RateLimiter

// BeforeCopy 是用户提供的回调函数，它在请求开始下载之前立即被调用。如果 BeforeCopy 返回错误，请求将被取消，并且相同的错误将返回在响应对象中。
// md5:865a949d97572fb2
	X传输开始之前回调 Hook

// AfterCopy 是用户提供的回调，它在请求下载完成后立即调用，但在校验和验证和关闭之前。此钩子仅在传输成功时调用。如果 AfterCopy 返回错误，请求将被取消，并在 Response 对象上返回相同的错误。
// md5:9e8a840d6dbbf503
	X传输完成之后回调 Hook

	// hash、checksum和deleteOnError - 通过SetChecksum设置。 md5:cc1ef74b5265e5bb
	hash          hash.Hash
	checksum      []byte
	deleteOnError bool

	// 用于取消和超时的上下文 - 通过WithContext设置 md5:b3765a4b0c8785c9
	ctx context.Context
}

// NewRequest 返回一个新的文件传输请求，适合与 Client.Do 使用。
// md5:1bf325ab23123d38
func X生成下载参数(保存目录, 下载链接 string) (*Request, error) {
	if 保存目录 == "" {
		保存目录 = "."
	}
	req, err := http.NewRequest("GET", 下载链接, nil)
	if err != nil {
		return nil, err
	}
	return &Request{
		Http协议头: req,
		X文件名:    保存目录,
	}, nil
}

// Context 返回请求的上下文。要更改上下文，请使用
// WithContext。
//
// 返回的上下文始终非空；它默认为背景上下文。
//
// 上下文控制取消操作。
// md5:48c4bb357919bed1
func (r *Request) I取上下文() context.Context {
	if r.ctx != nil {
		return r.ctx
	}

	return context.Background()
}

// WithContext 返回一个与 r 拥有相同结构但上下文改为 ctx 的浅拷贝。提供的 ctx 必须非 nil。
// md5:717222e0b2288cb8
func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(Request)
	*r2 = *r
	r2.ctx = ctx
	r2.Http协议头 = r2.Http协议头.WithContext(ctx)
	return r2
}

// URL 返回要下载的 URL。 md5:aebf141338bb627e
func (r *Request) X取下载链接() *url.URL {
	return r.Http协议头.URL
}

// SetChecksum 设置期望的哈希算法和校验值，用于验证下载的文件。下载完成后，将使用给定的哈希算法计算实际下载文件的校验和。如果校验和不匹配，关联的Response.Err方法将返回错误。
// 
// 如果deleteOnError为true，如果校验和验证失败，下载的文件将自动删除。
//
// 为了防止计算出的校验和被破坏，提供的哈希值在其他请求或goroutine中不应被使用。
//
// 要禁用校验和验证，请使用nil哈希值调用SetChecksum。
// md5:56e3d06ddceed502
func (r *Request) X设置完成后效验(h hash.Hash, sum []byte, 出错后删除 bool) {
	r.hash = h
	r.checksum = sum
	r.deleteOnError = 出错后删除
}
