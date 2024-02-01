package 下载类

import (
	"errors"
	"fmt"
	"net/http"
)

var (
// ErrBadLength 表示服务器响应或已存在的文件内容长度与预期的不匹配。
	ERR_文件长度不匹配 = errors.New("文件长度不匹配")

// ErrBadChecksum 表示下载的文件未能通过校验和验证。
	ERR_文件校验失败 = errors.New("文件校验失败")

// ErrNoFilename 表示无法通过 URL 或服务器响应头自动生成一个合理的文件名。
	ERR_无法确定文件名 = errors.New("无法确定文件名")

// ErrNoTimestamp 表示无法通过远程服务器响应头自动确定时间戳。
	ERR_无法确定时间戳 = errors.New("无法确定时间戳")

// ErrFileExists 表示目标路径已存在。
	ERR_文件已存在 = errors.New("file exists")
)

// StatusCodeError 表示服务器响应的状态码不在 200-299 范围内（在跟随所有重定向之后）。
// 注：此注释描述了一个Go语言错误类型，当HTTP请求的最终状态码不在成功范围（200-299）内时，会返回这个错误。
type X状态码错误 int

func (错误 X状态码错误) Error() string {
	return fmt.Sprintf("服务器返回%d %s", 错误, http.StatusText(int(错误)))
}

// IsStatusCodeError 判断给定的错误是否为 StatusCodeError 类型，如果是则返回 true。
func X是否为状态码错误(错误 error) bool {
	_, ok := 错误.(X状态码错误)
	return ok
}
