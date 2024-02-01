package 下载类

import "context"

// RateLimiter 是一个接口，任何用于限制下载传输速度的第三方限速器都必须满足此接口。
//
// 推荐使用的令牌桶实现可以在 https://godoc.org/golang.org/x/time/rate#Limiter 找到。
type RateLimiter interface {
	WaitN(ctx context.Context, n int) (err error)
}
