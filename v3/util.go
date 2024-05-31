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

// setLastModified 根据远程服务器返回的 Last-Modified 头部，设置本地文件的最后修改时间戳。
// md5:fd6bceaff3448251
func setLastModified(resp *http.Response, filename string) error {
// https://tools.ietf.org/html/rfc7232#section-2.2
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
// 
// 这段注释引用了RFC 7232（HTTP/1.1状态码和首部字段）的第2.2节。
// 它可能与HTTP响应头中的"Last-Modified"字段有关，该字段表示资源的最后修改日期，常用于决定是否需要重新获取资源。MDN（Mozilla开发者网络）文档对此字段有详细说明。
// md5:e0a757b6022b2ca0
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

// mkdirp为目标文件路径创建所有缺失的父目录。 md5:e9faeec74648c46e
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

// guessFilename 根据给定的 http.Response 返回一个文件名。如果无法确定，则返回 ErrNoFilename 错误。
//
// TODO: 不需要存储的操作不应该需要文件名。
// md5:d38ac2f1eefa7e95
func guessFilename(resp *http.Response) (string, error) {
	filename := resp.Request.URL.Path
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if val, ok := params["filename"]; ok {
				filename = val
			} // 如果缺少`else filename`指令，就使用URL.Path作为备选。 md5:8823a491b708b59c
		}
	}

	// sanitize
	if filename == "" || strings.HasSuffix(filename, "/") || strings.Contains(filename, "\x00") {
		return "", ERR_无法确定文件名
	}

	filename = filepath.Base(path.Clean("/" + filename))
	if filename == "" || filename == "." || filename == "/" {
		return "", ERR_无法确定文件名
	}

	return filename, nil
}
