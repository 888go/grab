package 下载类

import (
	"errors"
	"fmt"
	"net/http"
)

var (
// ErrBadLength表示服务器响应或现有文件的长度与预期内容长度不符。
// md5:da55fb8042067108
	ERR_文件长度不匹配 = errors.New("bad content length")

// ErrBadChecksum 表示下载的文件未能通过校验和验证。
// md5:ac12cf077b834024
	ERR_文件校验失败 = errors.New("checksum mismatch")

// ErrNoFilename表示无法使用服务器的URL或响应头自动确定一个合理的文件名。
// md5:9a298d198ab7a9e5
	ERR_无法确定文件名 = errors.New("no filename could be determined")

// ErrNoTimestamp表示无法使用远程服务器响应头自动确定时间戳。
// md5:b68398288f9bf056
	ERR_无法确定时间戳 = errors.New("no timestamp could be determined for the remote file")

	// ErrFileExists 表示目标路径已经存在。 md5:aa4e242a296ca268
	ERR_文件已存在 = errors.New("file exists")
)

// StatusCodeError表示服务器响应的状态码不在200-299范围内（在跟踪任何重定向后）。
// md5:d0a3bd375863dac3
type StatusCodeError int

func (err StatusCodeError) Error() string {
	return fmt.Sprintf("server returned %d %s", err, http.StatusText(int(err)))
}

// IsStatusCodeError 如果给定的错误是StatusCodeError类型，则返回true。 md5:ce211b2217978d1f
func X是否为状态码错误(错误 error) bool {
	_, ok := 错误.(StatusCodeError)
	return ok
}
