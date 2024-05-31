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

	// 验证请求方法 md5:71e60e72d4e0cfd0
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

	// set server options
	if h.acceptRanges {
		w.Header().Set("Accept-Ranges", "bytes")
	}

	// 设置附件文件名 md5:853840c1bcd20ee7
	if h.attachmentFilename != "" {
		w.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("attachment;filename=\"%s\"", h.attachmentFilename),
		)
	}

	// 设置最后修改时间戳 md5:78efc5dd8a42b4cd
	lastMod := time.Now()
	if !h.lastModified.IsZero() {
		lastMod = h.lastModified
	}
	w.Header().Set("Last-Modified", lastMod.Format(http.TimeFormat))

	// set content-length
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

	// apply header blacklist
	for _, key := range h.headerBlacklist {
		w.Header().Del(key)
	}

	// 发送头和状态码 md5:502df63c08265698
	w.WriteHeader(h.statusCodeFunc(r))

	// send body
	if r.Method == "GET" {
		// 使用缓冲输入流来减少读者的开销 md5:8df643db318497f8
		bw := bufio.NewWriterSize(w, 4096)
		for i := offset; !isRequestClosed(r) && i < h.contentLength; i++ {
			bw.Write([]byte{byte(i)})
			if h.rateLimiter != nil {
				bw.Flush()
				w.(http.Flusher).Flush() // 强制服务器将数据发送到客户端 md5:dfabe937f3096ed4
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

// isRequestClosed 判断客户端请求是否已取消，如果已取消则返回true。 md5:9265ac4addb1f469
func isRequestClosed(r *http.Request) bool {
	return r.Context().Err() != nil
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
