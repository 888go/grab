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
// 即使在远程服务器接收到 HTTP 响应，但内容尚未开始传输时，也可能返回一个响应。
//
// 所有 Response 方法调用都是线程安全的。
type X响应 struct {
// 该Response所对应的请求。
	X下载参数 *X下载参数

// HTTPResponse 代表从 HTTP 请求中接收到的 HTTP 响应。
//
// 不应使用响应体（response Body），因为它将被 grab 消耗并关闭。
	HTTP响应 *http.Response

// Filename 指定文件传输在本地存储中的保存路径。
	X文件名 string

// Size 指定文件传输的预期总大小。
	sizeUnsafe int64

// Start 指定文件传输开始的时间。
	X传输开始时间 time.Time

// End 指定文件传输完成的时间。
//
// 在传输尚未完成时，此属性将返回零值。
	X传输完成时间 time.Time

// CanResume 指定远程服务器声明它可以恢复先前的下载，因为已设置了 'Accept-Ranges: bytes' 头部。
	CanResume bool

// DidResume 指定文件传输恢复了先前未完成的传输。
	DidResume bool

// Done 在传输完成（无论成功或出现错误）后关闭。通过 Response.Err 可获取错误信息
	Done chan struct{}

// ctx 是一个 Context，用于控制正在进行的传输的取消
	ctx context.Context

// cancel 是一个取消函数，可用于取消此 Response 的上下文。
	cancel context.CancelFunc

// fi 是目标文件在传输开始前（如果已经存在）的 FileInfo 信息。
	fi os.FileInfo

// optionsKnown 表示已完成 HEAD 请求，并且已知远程服务器的功能。
	optionsKnown bool

// writer 是用于将下载的文件写入本地存储的文件句柄
	writer io.Writer

// storeBuffer 如果 Request.NoStore 被启用，则接收传输的内容。
	storeBuffer bytes.Buffer

// bytesCompleted 指定在这次传输开始之前已经完成传输的字节数。
	bytesResumed int64

// transfer 负责从远程服务器复制数据到本地文件，
// 并跟踪进度以及允许取消操作。
	transfer *transfer

// bufferSize 指定了传输缓冲区的大小（以字节为单位）。
	bufferSize int

// Error 包含了文件传输过程中可能出现的任何错误。
// 请在 IsComplete 返回 true 之前不要读取此内容。
	err error
}

// IsComplete 返回一个布尔值，如果下载已完成则返回true。如果在下载过程中发生错误，可以通过 Err 返回该错误。
func (c *X响应) X是否已完成() bool {
	select {
	case <-c.Done:
		return true
	default:
		return false
	}
}

// Cancel 取消文件传输，通过取消此 Response 对应的基础 Context 来实现。Cancel 会阻塞直到传输关闭并返回任何错误——通常是 context.Canceled。
func (c *X响应) X取消() error {
	c.cancel()
	return c.X等待错误()
}

// Wait会阻塞直到下载完成。
func (c *X响应) X等待完成() {
	<-c.Done
}

// Err 阻塞调用该方法的 goroutine，直到底层文件传输完成，并返回在此期间可能发生的任何错误。如果下载已经完成，Err 将立即返回。
func (c *X响应) X等待错误() error {
	<-c.Done
	return c.err
}

// Size 返回文件传输的大小。如果远程服务器没有指定总大小，并且传输未完成，则返回值为-1。
func (c *X响应) X取总字节() int64 {
	return atomic.LoadInt64(&c.sizeUnsafe)
}

// BytesComplete 返回已复制到目标位置的总字节数，包括从先前下载恢复的所有字节。
func (c *X响应) X已完成字节() int64 {
	return c.bytesResumed + c.transfer.N()
}

// BytesPerSecond 返回过去五秒钟内通过简单移动平均计算出的每秒传输字节数。如果下载已经完成，则返回整个下载过程中平均的每秒字节数。
func (c *X响应) X取每秒字节() float64 {
	if c.X是否已完成() {
		return float64(c.transfer.N()) / c.X取下载已持续时间().Seconds()
	}
	return c.transfer.BPS()
}

// Progress 返回已下载总字节的比例。将返回的值乘以100可得到完成的百分比。
func (c *X响应) X取进度() float64 {
	size := c.X取总字节()
	if size <= 0 {
		return 0
	}
	return float64(c.X已完成字节()) / float64(size)
}

// Duration 返回文件传输的持续时间。如果传输正在进行中，
// 持续时间将是从现在到传输开始之间的时间差。如果传输已完成，
// 持续时间将是整个完成传输过程从开始到结束的时间差。
func (c *X响应) X取下载已持续时间() time.Duration {
	if c.X是否已完成() {
		return c.X传输完成时间.Sub(c.X传输开始时间)
	}

	return time.Now().Sub(c.X传输开始时间)
}

// ETA 返回根据当前每秒字节数估算的下载完成时间。如果传输已完成，将返回实际结束时间。
func (c *X响应) X取估计完成时间() time.Time {
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

// Open 将阻塞调用的 goroutine，直到底层文件传输完成，然后打开已传输的文件以供读取。
// 如果启用了 Request.NoStore，则读取器将从内存中读取数据。
//
// 如果在传输过程中发生错误，将会返回该错误。
//
// 调用者有责任关闭返回的文件句柄。
func (c *X响应) X等待完成后打开文件() (io.ReadCloser, error) {
	if err := c.X等待错误(); err != nil {
		return nil, err
	}
	return c.openUnsafe()
}

func (c *X响应) openUnsafe() (io.ReadCloser, error) {
	if c.X下载参数.X不写入本地文件系统 {
		return ioutil.NopCloser(bytes.NewReader(c.storeBuffer.Bytes())), nil
	}
	return os.Open(c.X文件名)
}

// Bytes 阻塞调用它的 goroutine，直到底层文件传输完成，然后从已完成的传输中读取所有字节。
// 如果启用了 Request.NoStore，则字节将从内存中读取。
//
// 如果在传输过程中发生错误，将会返回该错误。
func (c *X响应) X等待完成后取字节集() ([]byte, error) {
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

func (c *X响应) requestMethod() string {
	if c == nil || c.HTTP响应 == nil || c.HTTP响应.Request == nil {
		return ""
	}
	return c.HTTP响应.Request.Method
}

func (c *X响应) checksumUnsafe() ([]byte, error) {
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

func (c *X响应) closeResponseBody() error {
	if c.HTTP响应 == nil || c.HTTP响应.Body == nil {
		return nil
	}
	return c.HTTP响应.Body.Close()
}
