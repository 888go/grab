package grabtest

import (
	"errors"
	"net/http"
	"time"
)

type HandlerOption func(*handler) error


// ff:
// code:
func StatusCodeStatic(code int) HandlerOption {
	return func(h *handler) error {
		return StatusCode(func(req *http.Request) int {
			return code
		})(h)
	}
}


// ff:
// f:
func StatusCode(f StatusCodeFunc) HandlerOption {
	return func(h *handler) error {
		if f == nil {
			return errors.New("status code function cannot be nil")
		}
		h.statusCodeFunc = f
		return nil
	}
}


// ff:
// methods:
func MethodWhitelist(methods ...string) HandlerOption {
	return func(h *handler) error {
		h.methodWhitelist = methods
		return nil
	}
}


// ff:
// headers:
func HeaderBlacklist(headers ...string) HandlerOption {
	return func(h *handler) error {
		h.headerBlacklist = headers
		return nil
	}
}


// ff:
// n:
func ContentLength(n int) HandlerOption {
	return func(h *handler) error {
		if n < 0 {
			return errors.New("content length must be zero or greater")
		}
		h.contentLength = n
		return nil
	}
}


// ff:
// enabled:
func AcceptRanges(enabled bool) HandlerOption {
	return func(h *handler) error {
		h.acceptRanges = enabled
		return nil
	}
}


// ff:
// t:
func LastModified(t time.Time) HandlerOption {
	return func(h *handler) error {
		h.lastModified = t.UTC()
		return nil
	}
}


// ff:
// d:
func TimeToFirstByte(d time.Duration) HandlerOption {
	return func(h *handler) error {
		if d < 1 {
			return errors.New("time to first byte must be greater than zero")
		}
		h.ttfb = d
		return nil
	}
}


// ff:
// bps:
func RateLimiter(bps int) HandlerOption {
	return func(h *handler) error {
		if bps < 1 {
			return errors.New("bytes per second must be greater than zero")
		}
		h.rateLimiter = time.NewTicker(time.Second / time.Duration(bps))
		return nil
	}
}


// ff:
// filename:
func AttachmentFilename(filename string) HandlerOption {
	return func(h *handler) error {
		h.attachmentFilename = filename
		return nil
	}
}
