package 下载类

import "context"

// RateLimiter 是一个接口，任何第三方的速率限制器都必须实现这个接口，以限制下载传输速度。
//
// 推荐的令牌桶实现可以在 https://godoc.org/golang.org/x/time/rate#Limiter 找到。
// md5:cd1978873813ac86
type RateLimiter interface {
	WaitN(ctx context.Context, n int) (err error)
}
