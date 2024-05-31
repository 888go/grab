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

// HTTPClient 提供了一个接口，允许我们执行HTTP请求。 md5:4171bffd13d00b3a
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// truncater是一个私有接口，允许使用不同的响应写入器进行截断
// md5:13b36272ac021c7d
type truncater interface {
	Truncate(size int64) error
}

// Client是一个文件下载客户端。
//
// 客户端是安全的，可以由多个goroutine并发使用。
// md5:adb110330920265b
type Client struct {
// HTTPClient 指定了在文件传输过程中与远程服务器通信所使用的 http.Client。
// md5:18f90ae2438c3d95
	HTTPClient HTTPClient

// UserAgent 指定将设置为此客户端的所有请求的头中的 User-Agent 字符串。
//
// 可以在每个请求的头中覆盖 User-Agent 字符串。
// md5:416ff171bdbaf067
	HTTP_UA string

// BufferSize 指定了用于传输所有请求文件的缓冲区的大小（以字节为单位）。较大的缓冲区可能会提高吞吐量，但会消耗更多内存，并导致传输进度统计更新更不频繁。每个请求的BufferSize可以在每个Request对象上重写。默认值：32KB。
// md5:76d8523deca24178
	X缓冲区大小 int
}

// NewClient 使用默认配置返回一个新的文件下载客户端。 md5:37b6062ca04f5dcf
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

// DefaultClient 是默认客户端，所有 Get 方便函数都会使用它。
// md5:3aee22c6e017f0f3
var DefaultClient = X创建客户端()

// Do 发送文件传输请求并返回一个文件传输响应，遵循客户端的HTTPClient配置（例如重定向、Cookie、身份验证等）。
// 
// 类似于http.Get，Do在启动传输时会阻塞，但一旦传输在后台goroutine中开始传输，或者在早期阶段失败，它就会立即返回。
// 
// 如果错误是由客户端策略（如CheckRedirect）引起的，或者存在HTTP协议或IO错误，将通过Response.Err返回错误。Response.Err会在传输完成，无论成功还是失败，都会阻塞调用者。
// md5:45eff2a029f7d8e0
func (c *Client) X下载(下载参数 *Request) *Response {
	// cancel 函数会通过 closeResponse 在所有代码路径上被调用 md5:e670a9d606a72b8c
	ctx, cancel := context.WithCancel(下载参数.I取上下文())
	下载参数 = 下载参数.WithContext(ctx)
	resp := &Response{
		X下载参数:    下载参数,
		X传输开始时间:      time.Now(),
		Done:       make(chan struct{}, 0),
		X文件名:   下载参数.X文件名,
		ctx:        ctx,
		cancel:     cancel,
		bufferSize: 下载参数.X缓冲区大小,
	}
	if resp.bufferSize == 0 {
		// 默认为Client.BufferSize md5:f7214bed22611823
		resp.bufferSize = c.X缓冲区大小
	}

// 当调用者被阻塞以初始化文件传输时，运行状态机。永远不要进入copyFile状态 - 这将在另一个goroutine中接下来发生。
// md5:c59c8fd584f98402
	c.run(resp, c.statFileInfo)

// 在新的goroutine中运行copyFile。如果传输已经完成或失败，copyFile将不执行任何操作。
// md5:0fe588a8a89f07a6
	go c.run(resp, c.copyFile)
	return resp
}

// DoChannel 通过给定的 Request 通道顺序执行所有发送的请求，直到由其他goroutine关闭。调用者会在Request通道关闭且所有传输完成后被阻塞。一旦从远程服务器接收到响应，所有响应会立即通过给定的Response通道发送，并可用于跟踪每个下载的进度。
// 
// 如果Response接收器处理速度慢，会导致worker阻塞，从而延迟已经启动连接的传输开始时间——可能导致服务器超时。调用者需要确保为Response通道使用足够的缓冲大小以防止这种情况。
// 
// 在任何文件传输过程中如果发生错误，可以通过关联的Response.Err函数访问到该错误。
// md5:1e8b3d5e8dee8a95
func (c *Client) DoChannel(reqch <-chan *Request, respch chan<- *Response) {
	// TODO：启用批量任务的取消功能 md5:551d7d31d87cb53d
	for req := range reqch {
		resp := c.X下载(req)
		respch <- resp
		<-resp.Done
	}
}

// DoBatch 使用指定数量的并发工作者执行所有给定的请求。一旦工作者启动，控制权就会返回给调用者。
//
// 如果请求的工作者数量小于1，那么将为每个请求创建一个工作者。即，所有请求将并行执行。
//
// 如果在任何文件传输过程中发生错误，可以通过调用关联的Response.Err来获取该错误。
//
// 返回的Response通道只会在所有给定的Requests完成（无论成功或失败）后关闭。
// md5:a143d0d000672213
func (c *Client) X多线程下载(线程数量 int, 下载参数 ...*Request) <-chan *Response {
	if 线程数量 < 1 {
		线程数量 = len(下载参数)
	}
	reqch := make(chan *Request, len(下载参数))
	respch := make(chan *Response, len(下载参数))
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

// StateFunc是一个动作，它会改变Response的状态，并返回下一个要调用的StateFunc。
// md5:2f895dfc8babb8bb
type stateFunc func(*Response) stateFunc

// run 调用给定的 stateFunc 函数，以及所有后续返回的 stateFuncs，直到某个 stateFunc 返回 nil 或者 Response 的 ctx 被取消。每个 stateFunc 应该修改给定 Response 的状态，直到下载完成或失败。
// md5:dcde558208dbf3d5
func (c *Client) run(resp *Response, f stateFunc) {
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

// statFileInfo 获取与Response.Filename匹配的任何本地文件的FileInfo。
//
// 如果文件不存在，是一个目录，或者其名称未知，则下一个状态函数为headRequest。
//
// 如果文件存在，则设置Response.fi，下一个状态函数为validateLocal。
//
// 如果发生错误，下一个状态函数为closeResponse。
// md5:6285457c42618336
func (c *Client) statFileInfo(resp *Response) stateFunc {
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

// validateLocal 将下载文件的本地副本与远程文件进行比较。
// 
// 如果本地文件大于远程文件，或者Request.SkipExisting 为 true，则返回错误。
// 
// 如果现有文件的大小与远程文件匹配，下一个状态函数是checksumFile。
// 
// 如果本地文件小于远程文件，并且已知远程服务器支持范围请求，下一个状态函数是getRequest。
// md5:fb339cef020987be
func (c *Client) validateLocal(resp *Response) stateFunc {
	if resp.X下载参数.X跳过已存在文件 {
		resp.err = ERR_文件已存在
		return c.closeResponse
	}

	// 确定目标文件大小 md5:b64a2f6b674bfe35
	expectedSize := resp.X下载参数.X预期文件大小
	if expectedSize == 0 && resp.HTTP响应 != nil {
		expectedSize = resp.HTTP响应.ContentLength
	}

	if expectedSize == 0 {
// size 是实际大小，可能是0或未知
// 如果未知，我们向远程服务器查询
// 如果已知为0，我们继续使用GET请求
// md5:b8aed7e2a19d006f
		return c.headRequest
	}

	if expectedSize == resp.fi.Size() {
		// 当本地文件的大小与远程文件相匹配时 - 封装它 md5:9020a4e52bdc768c
		resp.DidResume = true
		resp.bytesResumed = resp.fi.Size()
		return c.checksumFile
	}

	if resp.X下载参数.X不续传 {
		// 当地文件应被覆盖 md5:c6963052a01c5c41
		return c.getRequest
	}

	if expectedSize >= 0 && expectedSize < resp.fi.Size() {
		// 远程大小已知，比本地大小小，并且我们想要恢复下载 md5:4eff80977d3bd5ad
		resp.err = ERR_文件长度不匹配
		return c.closeResponse
	}

	if resp.CanResume {
		// 在GET请求中设置恢复范围 md5:baef5df94e869e30
		resp.X下载参数.Http协议头.Header.Set(
			"Range",
			fmt.Sprintf("bytes=%d-", resp.fi.Size()))
		resp.DidResume = true
		resp.bytesResumed = resp.fi.Size()
		return c.getRequest
	}
	return c.headRequest
}

func (c *Client) checksumFile(resp *Response) stateFunc {
	if resp.X下载参数.hash == nil {
		return c.closeResponse
	}
	if resp.X文件名 == "" {
		panic("grab: developer error: filename not set")
	}
	if resp.X取总字节() < 0 {
		panic("grab: developer error: size unknown")
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
				// err应该是os.PathError类型，并且应包含文件路径 md5:255ad01285b97e27
				resp.err = fmt.Errorf(
					"cannot remove downloaded file with checksum mismatch: %v",
					err)
			}
		}
	}
	return c.closeResponse
}

// doHTTPRequest 发送一个HTTP请求并返回响应 md5:19242108c18d9170
func (c *Client) doHTTPRequest(req *http.Request) (*http.Response, error) {
	if c.HTTP_UA != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.HTTP_UA)
	}
	return c.HTTPClient.Do(req)
}

func (c *Client) headRequest(resp *Response) stateFunc {
	if resp.optionsKnown {
		return c.getRequest
	}
	resp.optionsKnown = true

	if resp.X下载参数.X不续传 {
		return c.getRequest
	}

	if resp.X文件名 != "" && resp.fi == nil {
		// 目标路径已知且不存在 md5:d19d7aa9a062a804
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

// 如果在HEAD请求期间发生重定向，记录最终URL，并在发送后续请求时使用它。
// 这样，我们就避免向原始URL发送可能不被支持的请求，例如"Range"，因为是最终URL表明了其支持。
// md5:5c169e71aa18b9cb
	resp.X下载参数.Http协议头.URL = resp.HTTP响应.Request.URL
	resp.X下载参数.Http协议头.Host = resp.HTTP响应.Request.Host

	return c.readResponse
}

func (c *Client) getRequest(resp *Response) stateFunc {
	resp.HTTP响应, resp.err = c.doHTTPRequest(resp.X下载参数.Http协议头)
	if resp.err != nil {
		return c.closeResponse
	}

	// 待办事项：检查 Content-Range md5:1dd4a612f45654be

	// check status code
	if !resp.X下载参数.X忽略错误状态码 {
		if resp.HTTP响应.StatusCode < 200 || resp.HTTP响应.StatusCode > 299 {
			resp.err = StatusCodeError(resp.HTTP响应.StatusCode)
			return c.closeResponse
		}
	}

	return c.readResponse
}

func (c *Client) readResponse(resp *Response) stateFunc {
	if resp.HTTP响应 == nil {
		panic("grab: developer error: Response.HTTPResponse is nil")
	}

	// check expected size
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
		// Request.Filename将会是空的或者是一个目录 md5:9577464c343c9907
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

// openWriter 为写入打开目标文件，并定位到将继续传输的文件位置。
//
// 需要确保Response.Filename和resp.DidResume已经设置。
// md5:5670da9b8c85a4b1
func (c *Client) openWriter(resp *Response) stateFunc {
	if !resp.X下载参数.X不写入本地文件系统 && !resp.X下载参数.X不自动创建目录 {
		resp.err = mkdirp(resp.X文件名)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	if resp.X下载参数.X不写入本地文件系统 {
		resp.writer = &resp.storeBuffer
	} else {
		// compute write flags
		flag := os.O_CREATE | os.O_WRONLY
		if resp.fi != nil {
			if resp.DidResume {
				flag = os.O_APPEND | os.O_WRONLY
			} else {
// 如果BeforeCopy钩子没有取消，则稍后在copyFile中截断
// md5:8d0fcf3dbf9d4eb0
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

		// seek to start or end
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

	// 下一步是copyFile，但这个操作将在另一个goroutine中稍后被调用。 md5:bdcea988806d0db1
	return nil
}

// copy 将通过Client.do()建立的HTTP连接的内容进行传输 md5:feda3b303e54e72e
func (c *Client) copyFile(resp *Response) stateFunc {
	if resp.X是否已完成() {
		return nil
	}

	// run BeforeCopy hook
	if f := resp.X下载参数.X传输开始之前回调; f != nil {
		resp.err = f(resp)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	var bytesCopied int64
	if resp.transfer == nil {
		panic("grab: developer error: Response.transfer is nil")
	}

// 我们在openWriter()中推迟了截断文件的操作，以确保BeforeCopy不会取消复制。
// 如果这是一个现有的、不打算恢复的文件，那么截断其内容。
// md5:34b779d0e7dfa278
	if t, ok := resp.writer.(truncater); ok && resp.fi != nil && !resp.DidResume {
		t.Truncate(0)
	}

	bytesCopied, resp.err = resp.transfer.copy()
	if resp.err != nil {
		return c.closeResponse
	}
	closeWriter(resp)

	// set file timestamp
	if !resp.X下载参数.X不写入本地文件系统 && !resp.X下载参数.X忽略远程时间 {
		resp.err = setLastModified(resp.HTTP响应, resp.X文件名)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	// 如果之前未知，更新传输大小 md5:d285b51e526307dd
	if resp.X取总字节() < 0 {
		discoveredSize := resp.bytesResumed + bytesCopied
		atomic.StoreInt64(&resp.sizeUnsafe, discoveredSize)
		if resp.X下载参数.X预期文件大小 > 0 && resp.X下载参数.X预期文件大小 != discoveredSize {
			resp.err = ERR_文件长度不匹配
			return c.closeResponse
		}
	}

	// run AfterCopy hook
	if f := resp.X下载参数.X传输完成之后回调; f != nil {
		resp.err = f(resp)
		if resp.err != nil {
			return c.closeResponse
		}
	}

	return c.checksumFile
}

func closeWriter(resp *Response) {
	if closer, ok := resp.writer.(io.Closer); ok {
		closer.Close()
	}
	resp.writer = nil
}

// close 对Response进行最终化处理 md5:5f125695e8212203
func (c *Client) closeResponse(resp *Response) stateFunc {
	if resp.X是否已完成() {
		panic("grab: developer error: response already closed")
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
