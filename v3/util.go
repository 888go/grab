package 下载类

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// setLastModified 根据远程服务器返回的 Last-Modified 头部信息，设置本地文件的最后修改时间戳。
func setLastModified(resp *http.Response, filename string) error {
// 参考RFC 7232文档第2.2章节
// 及MDN Web文档中关于HTTP Headers的Last-Modified部分
// RFC 7232是HTTP/1.1协议中定义条件请求语义的标准，其中第2.2节详细说明了Last-Modified头部字段的用法。
// MDN Web文档（Mozilla开发者网络）对HTTP Headers中的Last-Modified做了详细介绍，这是一个用于指示资源最后修改时间的HTTP头部字段。
	header := resp.Header.Get("Last-Modified")
	if header == "" {
		return nil
	}
	lastmod, err := time.Parse(http.TimeFormat, header)
	if err != nil {
		return nil
	}
	return os.Chtimes(filename, lastmod, lastmod)
}

// mkdirp 为目标文件路径创建所有缺失的父级目录。
func mkdirp(path string) error {
	dir := filepath.Dir(path)
	if fi, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error checking destination directory: %v", err)
		}
		if err := os.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("error creating destination directory: %v", err)
		}
	} else if !fi.IsDir() {
		panic("grab: developer error: destination path is not directory")
	}
	return nil
}

// guessFilename 为给定的http.Response返回一个文件名。如果无法确定文件名，则返回ErrNoFilename错误。
//
// TODO: 对于NoStore操作，不应要求提供文件名
func guessFilename(resp *http.Response) (string, error) {
	filename := resp.Request.URL.Path
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if val, ok := params["filename"]; ok {
				filename = val
			} // 如果filename指令缺失，则退回到URL.Path
		}
	}

	// sanitize
	if filename == "" || strings.HasSuffix(filename, "/") || strings.Contains(filename, "\x00") {
		return "", ErrNoFilename
	}

	filename = filepath.Base(path.Clean("/" + filename))
	if filename == "" || filename == "." || filename == "/" {
		return "", ErrNoFilename
	}

	return filename, nil
}
