package 下载类

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

// HTTPClient 提供了一个接口，使我们能够执行 HTTP 请求。
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// truncater 是一个私有接口，用于支持不同的响应写入器进行截断操作
type truncater interface {
	Truncate(size int64) error
}

// Client 是一个文件下载客户端。
//
// Client 对象支持在多个goroutine中并发安全地使用。
type Client struct {
// HTTPClient 指定在文件传输过程中与远程服务器通信时所使用的 http.Client。
	HTTPClient HTTPClient

// UserAgent 指定此客户端发起的所有请求中将在头部设置的 User-Agent 字符串。
//
// 用户代理字符串可以在每个请求的头部进行覆盖。
	HTTP_UA string

// BufferSize 指定用于传输所有请求文件的缓冲区大小（以字节为单位）。更大的缓冲区可能会带来更快的数据传输速度，但会消耗更多的内存，并导致传输进度统计信息更新频率降低。每个请求都可以在各自的 Request 对象上覆盖 BufferSize 属性。默认值：32KB。
	X缓冲区大小 int
}

// NewClient 返回一个使用默认配置的新文件下载客户端。
func X创建客户端() *Client {
	return &Client{
		HTTP_UA: "grab",
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}
}

// DefaultClient 是默认的客户端，被所有 Get 方便函数使用。
var X默认全局客户端 = X创建客户端()

// Do 发送文件传输请求并返回文件传输响应，遵循客户端 HTTPClient 上配置的策略（例如重定向、Cookie 和身份验证）。
//
// 类似于 http.Get，Do 在传输开始时阻塞，但一旦传输在后台 goroutine 中启动或早期失败，则立即返回。
//
// 若因客户端策略（如 CheckRedirect）导致错误，或者发生 HTTP 协议错误或 IO 错误，将通过 Response.Err 返回错误。Response.Err 将阻塞调用者直到传输完成（无论成功与否）。
func (c *Client) X下载(下载参数 *X下载参数) *X响应 {
// 当通过 closeResponse 在所有代码路径上调用时，cancel 将被调用
	ctx, cancel := context.WithCancel(下载参数.I取上下文())
	下载参数 = 下载参数.WithContext(ctx)
	resp := &X响应{
		X下载参数:    下载参数,
		X传输开始时间:      time.Now(),
		Done:       make(chan struct{}, 0),
		X文件名:   下载参数.X文件名,
		ctx:        ctx,
		cancel:     cancel,
		bufferSize: 下载参数.X缓冲区大小,
	}
	if resp.bufferSize == 0 {
// 默认为Client.BufferSize
		resp.bufferSize = c.X缓冲区大小
	}

// 当调用者阻塞以初始化文件传输时，运行状态机。
// 绝对不能转换到 copyFile 状态——这将在另一个 goroutine 中紧接着发生。
	c.run(resp, c.statFileInfo)

// 在一个新的goroutine中运行copyFile。如果传输已经完成或失败，copyFile将执行空操作（不做任何事）。
	go c.run(resp, c.copyFile)
	return resp
}

// DoChannel 针对给定的 Request 通道执行所有发送过来的请求，每次执行一个，直到被另一个 goroutine 关闭。调用方将被阻塞，直到 Request 通道关闭且所有传输完成。所有从远程服务器接收到的响应都会立即通过给定的 Response 通道发送，并可用于跟踪每个下载进度。
//
// 如果 Response 接收者处理速度较慢，将会导致工作线程阻塞，从而延迟已建立连接的传输开始时间，可能引发服务器超时。调用方有责任确保为 Response 通道使用足够大的缓冲区大小以防止此类情况发生。
//
// 如果在文件传输过程中出现任何错误，可通过相应的 Response.Err 函数访问该错误。
func (c *Client) DoChannel(reqch <-chan *X下载参数, respch chan<- *X响应) {
// TODO: 启用批量作业的取消功能
	for req := range reqch {
		resp := c.X下载(req)
		respch <- resp
		<-resp.Done
	}
}

// DoBatch 使用给定数量的并发工作者执行所有给定的请求。一旦工作者启动，控制权将返回给调用者。
//
// 如果请求的工作者数量少于1个，则每个请求都会创建一个工作者。即，所有请求都将并发执行。
//
// 如果在任何文件传输过程中发生错误，可以通过调用相关联的Response.Err获取该错误。
//
// 只有在所有给定的Requests都完成（无论成功与否）后，才会关闭返回的Response通道。
func (c *Client) X多线程下载(线程数量 int, 下载参数 ...*X下载参数) <-chan *X响应 {
	if 线程数量 < 1 {
		线程数量 = len(下载参数)
	}
	reqch := make(chan *X下载参数, len(下载参数))
	respch := make(chan *X响应, len(下载参数))
	wg := sync.WaitGroup{}
	for i := 0; i < 线程数量; i++ {
		wg.Add(1)
		go func() {
			c.DoChannel(reqch, respch)
			wg.Done()
		}()
	}

	// queue requests
	go func() {
		for _, req := range 下载参数 {
			reqch <- req
		}
		close(reqch)
		wg.Wait()
		close(respch)
	}()
	return respch
}

// stateFunc 是一种操作，它会修改 Response 的状态，并返回接下来要调用的下一个 stateFunc。
type stateFunc func(*X响应) stateFunc

// run调用给定的stateFunc函数及其后续返回的所有stateFuncs，
// 直到某个stateFunc返回nil或者Response.ctx被取消为止。每个stateFunc
// 应该根据需要改变给定Response的状态，直到下载完成或失败。
func (c *Client) run(resp *X响应, f stateFunc) {
	for {
		select {
		case <-resp.ctx.Done():
			if resp.X是否已完成() {
				return
			}
			resp.err = resp.ctx.Err()
			f = c.closeResponse

		default:
			// keep working
		}
		if f = f(resp); f == nil {
			return
		}
	}
}

// statFileInfo 函数用于获取与 Response.Filename 匹配的本地文件的 FileInfo 信息。
//
// 如果该文件不存在，或者是一个目录，亦或是名称未知，则下一个状态函数为 headRequest。
//
// 如果文件存在，将设置 Response.fi，并且下一个状态函数为 validateLocal。
//
// 如果出现错误，下一个状态函数则为 closeResponse。
func (c *Client) statFileInfo(resp *X响应) stateFunc {
	if resp.X下载参数.X不写入本地文件系统 || resp.X文件名 == "" {
		return c.headRequest
	}
	fi, err := os.Stat(resp.X文件名)
	if err != nil {
		if os.IsNotExist(err) {
			return c.headRequest
		}
		resp.err = err
		return c.closeResponse
	}
	if fi.IsDir() {
		resp.X文件名 = ""
		return c.headRequest
	}
	resp.fi = fi
	return c.validateLocal
}

// validateLocal 用于比较本地下载文件的副本与远程文件。
//
// 当本地文件大于远程文件，或者 Request.SkipExisting 为真时，会返回错误。
//
// 如果已存在的文件大小与远程文件相同，则下一个状态函数是 checksumFile。
//
// 如果本地文件小于远程文件，并且已知远程服务器支持范围请求，则下一个状态函数是 getRequest。
func (c *Client) validateLocal(resp *X响应) stateFunc {
	if resp.X下载参数.X跳过已存在文件 {
		resp.err = ERR_文件已存在
		return c.closeResponse
	}

// 确定目标文件大小
	expectedSize := resp.X下载参数.X预期文件大小
	if expectedSize == 0 && resp.HTTP响应 != nil {
		expectedSize = resp.HTTP响应.ContentLength
	}

	if expectedSize == 0 {
// size 实际上是0或者未知
// 如果是未知，我们将询问远程服务器
// 如果已知为0，我们将继续进行GET请求
		return c.headRequest
	}

	if expectedSize == resp.fi.Size() {
// 当本地文件与远程文件大小相匹配时 - 完成处理
		resp.DidResume = true
		resp.bytesResumed = resp.fi.Size()
		return c.checksumFile
	}

	if resp.X下载参数.X不续传 {
// 本地文件应被覆盖
		return c.getRequest
	}

	if expectedSize >= 0 && expectedSize < resp.fi.Size() {
// 远程大小已知，且小于本地大小，我们希望从中断处继续恢复
		resp.err = ERR_文件长度不匹配
		return c.closeResponse
	}

	if resp.CanResume {
// 在GET请求中设置恢复范围
		resp.X下载参数.Http协议头.Header.Set(
			"Range",
			fmt.Sprintf("bytes=%d-", resp.fi.Size()))
		resp.DidResume = true
		resp.bytesResumed = resp.fi.Size()
		return c.getRequest
	}
	return c.headRequest
}

func (c *Client) checksumFile(resp *X响应) stateFunc {
	if resp.X下载参数.hash == nil {
		return c.closeResponse
	}
	if resp.X文件名 == "" {
		panic("下载类: 开发人员错误:文件名未设置")
	}
	if resp.X取总字节() < 0 {
		panic("下载类: 开发人员错误:大小未知")
	}
	req := resp.X下载参数

	// compute checksum
	var sum []byte
	sum, resp.err = resp.checksumUnsafe()
	if resp.err != nil {
		return c.closeResponse
	}

	// compare checksum
	if !bytes.Equal(sum, req.checksum) {
		resp.err = ERR_文件校验失败
		if !resp.X下载参数.X不写入本地文件系统 && req.deleteOnError {
			if err := os.Remove(resp.X文件名); err != nil {
// err 应该是 os.PathError 类型，并且包含文件路径信息
				resp.err = fmt.Errorf(
					"下载类: 无法删除已下载的文件，因为校验和不匹配: %v",
					err)
			}
		}
	}
	return c.closeResponse
}

// doHTTPRequest 发送一个HTTP请求并返回响应
func (c *Client) doHTTPRequest(req *http.Request) (*http.Response, error) {
	if c.HTTP_UA != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.HTTP_UA)
	}
	return c.HTTPClient.Do(req)
}

func (c *Client) headRequest(resp *X响应) stateFunc {
	if resp.optionsKnown {
		return c.getRequest
	}
	resp.optionsKnown = true

	if resp.X下载参数.X不续传 {
		return c.getRequest
	}

	if resp.X文件名 != "" && resp.fi == nil {
// 目标路径已知且不存在
		return c.getRequest
	}

	hreq := new(http.Request)
	*hreq = *resp.X下载参数.Http协议头
	hreq.Method = "HEAD"

	resp.HTTP响应, resp.err = c.doHTTPRequest(hreq)
	if resp.err != nil {
		return c.closeResponse
	}
	resp.HTTP响应.Body.Close()

	if resp.HTTP响应.StatusCode != http.StatusOK {
		return c.getRequest
	}

// 当HEAD请求过程中出现重定向时，记录最终的URL，并在发送后续请求时使用它替代原始URL。
// 这样做可以避免向原始URL发送可能不受支持的请求，例如"Range"请求，因为是最终的URL声明了对该请求的支持。
	resp.X下载参数.Http协议头.URL = resp.HTTP响应.Request.URL
	resp.X下载参数.Http协议头.Host = resp.HTTP响应.Request.Host

	return c.readResponse
}

func (c *Client) getRequest(resp *X响应) stateFunc {
	resp.HTTP响应, resp.err = c.doHTTPRequest(resp.X下载参数.Http协议头)
	if resp.err != nil {
		return c.closeResponse
	}

// TODO: 检查 Content-Range

// 检查状态码
	if !resp.X下载参数.X忽略错误状态码 {
		if resp.HTTP响应.StatusCode < 200 || resp.HTTP响应.StatusCode > 299 {
			resp.err = X状态码错误(resp.HTTP响应.StatusCode)
			return c.closeResponse
		}
	}

	return c.readResponse
}

func (c *Client) readResponse(resp *X响应) stateFunc {
	if resp.HTTP响应 == nil {
		panic("下载类: 开发人员错误: Response.HTTPResponse 返回 nil")
	}

// 检查期望的大小
	resp.sizeUnsafe = resp.HTTP响应.ContentLength
	if resp.sizeUnsafe >= 0 {
		// remote size is known
		resp.sizeUnsafe += resp.bytesResumed
		if resp.X下载参数.X预期文件大小 > 0 && resp.X下载参数.X预期文件大小 != resp.sizeUnsafe {
			resp.err = ERR_文件长度不匹配
			return c.closeResponse
		}
	}

	// check filename
	if resp.X文件名 == "" {
		filename, err := guessFilename(resp.HTTP响应)
		if err != nil {
			resp.err = err
			return c.closeResponse
		}
// Request.Filename 将会是空值或者是一个目录名
		resp.X文件名 = filepath.Join(resp.X下载参数.X文件名, filename)
	}

	if !resp.X下载参数.X不写入本地文件系统 && resp.requestMethod() == "HEAD" {
		if resp.HTTP响应.Header.Get("Accept-Ranges") == "bytes" {
			resp.CanResume = true
		}
		return c.statFileInfo
	}
	return c.openWriter
}

// openWriter 打开目标文件以进行写入操作，并定位到文件传输将从中恢复的位置。
//
// 要求已预先设置 Response.Filename 和 resp.DidResume。
func (c *Client) openWriter(resp *X响应) stateFunc {
	if !resp.X下载参数.X不写入本地文件系统 && !resp.X下载参数.X不自动创建目录 {
		resp.err = mkdirp(resp.X文件名)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	if resp.X下载参数.X不写入本地文件系统 {
		resp.writer = &resp.storeBuffer
	} else {
// 计算写入标志
		flag := os.O_CREATE | os.O_WRONLY
		if resp.fi != nil {
			if resp.DidResume {
				flag = os.O_APPEND | os.O_WRONLY
			} else {
// 如果在copyFile过程中未被BeforeCopy钩子取消，则稍后截断
				flag = os.O_WRONLY
			}
		}

		// open file
		f, err := os.OpenFile(resp.X文件名, flag, 0666)
		if err != nil {
			resp.err = err
			return c.closeResponse
		}
		resp.writer = f

// 寻找并定位到起始或结束位置
		whence := os.SEEK_SET
		if resp.bytesResumed > 0 {
			whence = os.SEEK_END
		}
		_, resp.err = f.Seek(0, whence)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	// init transfer
	if resp.bufferSize < 1 {
		resp.bufferSize = 32 * 1024
	}
	b := make([]byte, resp.bufferSize)
	resp.transfer = newTransfer(
		resp.X下载参数.I取上下文(),
		resp.X下载参数.X速率限制器,
		resp.writer,
		resp.HTTP响应.Body,
		b)

// 下一步是调用copyFile函数，但该函数将在稍后的另一个goroutine中被调用
	return nil
}

// copy 将通过Client.do()建立的HTTP连接的内容进行传输
func (c *Client) copyFile(resp *X响应) stateFunc {
	if resp.X是否已完成() {
		return nil
	}

// 执行 BeforeCopy 钩子函数
	if f := resp.X下载参数.X传输开始之前回调; f != nil {
		resp.err = f(resp)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	var bytesCopied int64
	if resp.transfer == nil {
		panic("下载类: 开发人员错误: Response.transfer 返回 nil")
	}

// 我们在openWriter()函数中等待截断文件，是为了确保BeforeCopy不会取消复制操作。如果这是一个现有且不打算续传的文件，则截断其内容。
	if t, ok := resp.writer.(truncater); ok && resp.fi != nil && !resp.DidResume {
		t.Truncate(0)
	}

	bytesCopied, resp.err = resp.transfer.copy()
	if resp.err != nil {
		return c.closeResponse
	}
	closeWriter(resp)

// 设置文件时间戳
	if !resp.X下载参数.X不写入本地文件系统 && !resp.X下载参数.X忽略远程时间 {
		resp.err = setLastModified(resp.HTTP响应, resp.X文件名)
		if resp.err != nil {
			return c.closeResponse
		}
	}

// 如果之前未知，则更新传输大小
	if resp.X取总字节() < 0 {
		discoveredSize := resp.bytesResumed + bytesCopied
		atomic.StoreInt64(&resp.sizeUnsafe, discoveredSize)
		if resp.X下载参数.X预期文件大小 > 0 && resp.X下载参数.X预期文件大小 != discoveredSize {
			resp.err = ERR_文件长度不匹配
			return c.closeResponse
		}
	}

// 运行 AfterCopy 钩子
	if f := resp.X下载参数.X传输完成之后回调; f != nil {
		resp.err = f(resp)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	return c.checksumFile
}

func closeWriter(resp *X响应) {
	if closer, ok := resp.writer.(io.Closer); ok {
		closer.Close()
	}
	resp.writer = nil
}

// close 方法用于最终化（关闭）Response
func (c *Client) closeResponse(resp *X响应) stateFunc {
	if resp.X是否已完成() {
		panic("下载类: 开发人员错误: 响应已经关闭")
	}

	resp.fi = nil
	closeWriter(resp)
	resp.closeResponseBody()

	resp.X传输完成时间 = time.Now()
	close(resp.Done)
	if resp.cancel != nil {
		resp.cancel()
	}

	return nil
}
