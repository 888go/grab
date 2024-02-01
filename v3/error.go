package 下载类

import (
	"errors"
	"fmt"
	"net/http"
)

var (
// ErrBadLength 表示服务器响应或已存在的文件内容长度与预期的不匹配。
	ErrBadLength = errors.New("bad content length")

// ErrBadChecksum 表示下载的文件未能通过校验和验证。
	ErrBadChecksum = errors.New("checksum mismatch")

// ErrNoFilename 表示无法通过 URL 或服务器响应头自动生成一个合理的文件名。
	ErrNoFilename = errors.New("no filename could be determined")

// ErrNoTimestamp 表示无法通过远程服务器响应头自动确定时间戳。
	ErrNoTimestamp = errors.New("no timestamp could be determined for the remote file")

// ErrFileExists 表示目标路径已存在。
	ErrFileExists = errors.New("file exists")
)

// StatusCodeError 表示服务器响应的状态码不在 200-299 范围内（在跟随所有重定向之后）。
// 注：此注释描述了一个Go语言错误类型，当HTTP请求的最终状态码不在成功范围（200-299）内时，会返回这个错误。
type StatusCodeError int

func (err StatusCodeError) Error() string {
	return fmt.Sprintf("server returned %d %s", err, http.StatusText(int(err)))
}

// IsStatusCodeError 判断给定的错误是否为 StatusCodeError 类型，如果是则返回 true。
func IsStatusCodeError(err error) bool {
	_, ok := err.(StatusCodeError)
	return ok
}
