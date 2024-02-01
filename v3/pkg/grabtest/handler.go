package grabtest

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	DefaultHandlerContentLength       = 1 << 20
	DefaultHandlerMD5Checksum         = "c35cc7d8d91728a0cb052831bc4ef372"
	DefaultHandlerMD5ChecksumBytes    = MustHexDecodeString(DefaultHandlerMD5Checksum)
	DefaultHandlerSHA256Checksum      = "fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83"
	DefaultHandlerSHA256ChecksumBytes = MustHexDecodeString(DefaultHandlerSHA256Checksum)
)

type StatusCodeFunc func(req *http.Request) int

type handler struct {
	statusCodeFunc     StatusCodeFunc
	methodWhitelist    []string
	headerBlacklist    []string
	contentLength      int
	acceptRanges       bool
	attachmentFilename string
	lastModified       time.Time
	ttfb               time.Duration
	rateLimiter        *time.Ticker
}

func NewHandler(options ...HandlerOption) (http.Handler, error) {
	h := &handler{
		statusCodeFunc:  func(req *http.Request) int { return http.StatusOK },
		methodWhitelist: []string{"GET", "HEAD"},
		contentLength:   DefaultHandlerContentLength,
		acceptRanges:    true,
	}
	for _, option := range options {
		if err := option(h); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func WithTestServer(t *testing.T, f func(url string), options ...HandlerOption) {
	h, err := NewHandler(options...)
	if err != nil {
		t.Fatalf("unable to create test server handler: %v", err)
		return
	}
	s := httptest.NewServer(h)
	defer func() {
		h.(*handler).close()
		s.Close()
	}()
	f(s.URL)
}

func (h *handler) close() {
	if h.rateLimiter != nil {
		h.rateLimiter.Stop()
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// delay response
	if h.ttfb > 0 {
		time.Sleep(h.ttfb)
	}

// 验证请求方法
	allowed := false
	for _, m := range h.methodWhitelist {
		if r.Method == m {
			allowed = true
			break
		}
	}
	if !allowed {
		httpError(w, http.StatusMethodNotAllowed)
		return
	}

// 设置服务器选项
	if h.acceptRanges {
		w.Header().Set("Accept-Ranges", "bytes")
	}

// 设置附件文件名
	if h.attachmentFilename != "" {
		w.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("attachment;filename=\"%s\"", h.attachmentFilename),
		)
	}

// 设置最后修改时间戳
	lastMod := time.Now()
	if !h.lastModified.IsZero() {
		lastMod = h.lastModified
	}
	w.Header().Set("Last-Modified", lastMod.Format(http.TimeFormat))

// 设置内容长度
	offset := 0
	if h.acceptRanges {
		if reqRange := r.Header.Get("Range"); reqRange != "" {
			if _, err := fmt.Sscanf(reqRange, "bytes=%d-", &offset); err != nil {
				httpError(w, http.StatusBadRequest)
				return
			}
			if offset >= h.contentLength {
				httpError(w, http.StatusRequestedRangeNotSatisfiable)
				return
			}
		}
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", h.contentLength-offset))

// 应用头部黑名单
	for _, key := range h.headerBlacklist {
		w.Header().Del(key)
	}

// 发送头部和状态码
	w.WriteHeader(h.statusCodeFunc(r))

	// send body
	if r.Method == "GET" {
// 使用带缓冲的I/O以减少对读取器的开销
		bw := bufio.NewWriterSize(w, 4096)
		for i := offset; !isRequestClosed(r) && i < h.contentLength; i++ {
			bw.Write([]byte{byte(i)})
			if h.rateLimiter != nil {
				bw.Flush()
				w.(http.Flusher).Flush() // 强制服务器将数据发送到客户端
				select {
				case <-h.rateLimiter.C:
				case <-r.Context().Done():
				}
			}
		}
		if !isRequestClosed(r) {
			bw.Flush()
		}
	}
}

// isRequestClosed 返回 true，如果客户端请求已被取消。
func isRequestClosed(r *http.Request) bool {
	return r.Context().Err() != nil
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
