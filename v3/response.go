package 下载类

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

// Response 表示已完成或正在进行的下载请求的响应。
//
//一旦从远程服务器接收到HTTP响应，就可能返回一个响应，但在开始传输主体内容之前。
//
//对Response方法的所有调用都是线程安全的。
// md5:ddbfa271b1d4fd24
type Response struct {
	// 提交以获取此响应的请求。 md5:75fc5e8790cf3c19
	X下载参数 *Request

// HTTPResponse 表示从 HTTP 请求接收到的 HTTP 响应。
// 
// 响应体不应该被使用，因为它会被 grab 消费并关闭。
// md5:de050e04bc066b54
	HTTP响应 *http.Response

// Filename 指定了文件传输在本地存储中的路径。
// md5:2739dccaf6d8bae1
	X文件名 string

	// Size 指定文件传输的总预期大小。 md5:6c3a6da7f95a7aa6
	sizeUnsafe int64

	// Start 指定文件传输开始的时间。 md5:0b7ef91daaa881e6
	X传输开始时间 time.Time

// End 指定文件传输完成的时间。
//
// 在传输完成之前，此函数将返回零值。
// md5:22aad30fb1342e5f
	X传输完成时间 time.Time

// CanResume 指定远程服务器已宣布它可以恢复以前的下载，因为已设置 "Accept-Ranges: bytes" 头。
// md5:d5b949adac4b5a26
	CanResume bool

// DidResume 指定文件传输恢复了先前未完成的传输。
// md5:9e59c01300744b13
	DidResume bool

// Done 在传输完成时关闭，无论是成功还是出现错误。可以通过 Response.Err 获取错误信息。
// md5:a87b182e4efaa89f
	Done chan struct{}

	// ctx是一个控制进行中传输取消的Context md5:a1a2dcb143eb803b
	ctx context.Context

// cancel 是一个取消函数，可以用来取消这个Response的上下文。
// md5:7803555a55294ebb
	cancel context.CancelFunc

// fi 是目标文件在传输开始前已存在的FileInfo。
// md5:80de0f1066a0dd4a
	fi os.FileInfo

// optionsKnown表示已经完成了HEAD请求，并且知道了远程服务器的功能。
// md5:1808cb9d787356c6
	optionsKnown bool

// writer 是用于将下载的文件写入本地存储的文件句柄
// md5:bd2f23db00719ebb
	writer io.Writer

// storeBuffer 接收传输的内容，如果 Request.NoStore 开关启用的话。
// md5:c4f455f8b39f4595
	storeBuffer bytes.Buffer

// bytesCompleted 指定的是在本次传输开始之前已经传输的字节数。
// md5:03aa448103fe412e
	bytesResumed int64

// transfer负责从远程服务器复制数据到本地文件，跟踪进度并允许取消。
// md5:9a7cce1d1337a757
	transfer *transfer

	// bufferSize 指定了传输缓冲区的字节大小。 md5:fc0bc9d5d97bffef
	bufferSize int

// Error 包含文件传输过程中可能发生的任何错误。在 IsComplete 返回 true 之前，不应读取此值。
// md5:5938a32425f8b1ad
	err error
}

// IsComplete 返回如果下载已完成为true。如果在下载过程中发生错误，可以通过Err返回该错误。
// md5:eb0348c8e5031ff2
func (c *Response) X是否已完成() bool {
	select {
	case <-c.Done:
		return true
	default:
		return false
	}
}

// Cancel 通过取消此 Response 的底层上下文来取消文件传输。Cancel 函数会阻塞直到传输关闭，并返回任何错误——通常为 context.Canceled。
// md5:4277ad03dbb17f34
func (c *Response) X取消() error {
	c.cancel()
	return c.X等待错误()
}

// Wait 会阻塞直到下载完成。 md5:2c909b1e4febc570
func (c *Response) X等待完成() {
	<-c.Done
}

// Err会阻塞调用的goroutine，直到底层文件传输完成，并返回可能发生的任何错误。如果下载已经完成，Err会立即返回。
// md5:1eec3a720d62a01e
func (c *Response) X等待错误() error {
	<-c.Done
	return c.err
}

// Size 返回文件传输的大小。如果远程服务器没有指定总大小且传输不完整，返回值为 -1。
// md5:53bd23aba7b5b22a
func (c *Response) X取总字节() int64 {
	return atomic.LoadInt64(&c.sizeUnsafe)
}

// BytesComplete返回已复制到目标的总字节数，包括从先前下载中恢复的任何字节。
// md5:8b9d515c7cdd428f
func (c *Response) X已完成字节() int64 {
	return c.bytesResumed + c.transfer.N()
}

// BytesPerSecond 返回过去五秒内的平均每秒传输字节数。如果下载已完成，则返回整个下载过程中的平均字节数/秒。
// md5:dd425949dfaaf72f
func (c *Response) X取每秒字节() float64 {
	if c.X是否已完成() {
		return float64(c.transfer.N()) / c.X取下载已持续时间().Seconds()
	}
	return c.transfer.BPS()
}

// Progress 返回已下载的总字节数的比例。将返回值乘以100可得到完成的百分比。
// md5:163669af2595df84
func (c *Response) X取进度() float64 {
	size := c.X取总字节()
	if size <= 0 {
		return 0
	}
	return float64(c.X已完成字节()) / float64(size)
}

// Duration 返回文件传输的持续时间。如果传输正在进行中，持续时间将是当前时间和传输开始时间之间的差。如果传输已完成，持续时间将是完成传输过程的开始和结束时间之间的差。
// md5:851e9a4335644cc7
func (c *Response) X取下载已持续时间() time.Duration {
	if c.X是否已完成() {
		return c.X传输完成时间.Sub(c.X传输开始时间)
	}

	return time.Now().Sub(c.X传输开始时间)
}

// ETA 返回给定当前每秒字节数（BytesPerSecond）下载预计完成的时间。如果传输已经完成，将返回实际结束时间。
// md5:7f8fcb12ee64da7f
func (c *Response) X取估计完成时间() time.Time {
	if c.X是否已完成() {
		return c.X传输完成时间
	}
	bt := c.X已完成字节()
	bps := c.transfer.BPS()
	if bps == 0 {
		return time.Time{}
	}
	secs := float64(c.X取总字节()-bt) / bps
	return time.Now().Add(time.Duration(secs) * time.Second)
}

// Open函数会让调用的goroutine阻塞，直到底层文件传输完成，然后打开已传输的文件以供阅读。如果Request.NoStore选项被启用，读取器将从内存中读取。
//
// 如果在传输过程中发生错误，该错误将被返回。
//
// 调用者有责任关闭返回的文件句柄。
// md5:92d1addd1c15e703
func (c *Response) X等待完成后打开文件() (io.ReadCloser, error) {
	if err := c.X等待错误(); err != nil {
		return nil, err
	}
	return c.openUnsafe()
}

func (c *Response) openUnsafe() (io.ReadCloser, error) {
	if c.X下载参数.X不写入本地文件系统 {
		return ioutil.NopCloser(bytes.NewReader(c.storeBuffer.Bytes())), nil
	}
	return os.Open(c.X文件名)
}

// Bytes 使调用的goroutine阻塞，直到底层文件传输完成，然后读取已完成传输的所有字节。如果启用了Request.NoStore，将从内存中读取字节。
// 
// 如果在传输过程中发生错误，将返回该错误。
// md5:af539342438b13c1
func (c *Response) X等待完成后取字节集() ([]byte, error) {
	if err := c.X等待错误(); err != nil {
		return nil, err
	}
	if c.X下载参数.X不写入本地文件系统 {
		return c.storeBuffer.Bytes(), nil
	}
	f, err := c.X等待完成后打开文件()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func (c *Response) requestMethod() string {
	if c == nil || c.HTTP响应 == nil || c.HTTP响应.Request == nil {
		return ""
	}
	return c.HTTP响应.Request.Method
}

func (c *Response) checksumUnsafe() ([]byte, error) {
	f, err := c.openUnsafe()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := newTransfer(c.X下载参数.I取上下文(), nil, c.X下载参数.hash, f, nil)
	if _, err = t.copy(); err != nil {
		return nil, err
	}
	sum := c.X下载参数.hash.Sum(nil)
	return sum, nil
}

func (c *Response) closeResponseBody() error {
	if c.HTTP响应 == nil || c.HTTP响应.Body == nil {
		return nil
	}
	return c.HTTP响应.Body.Close()
}
