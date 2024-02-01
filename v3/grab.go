package 下载类

import (
	"fmt"
	"os"
)

// Get通过发送HTTP请求并下载请求URL的内容到指定的目标文件路径。调用方会一直阻塞，直到下载完成（无论成功与否）。
//
// 若因客户端策略（例如CheckRedirect）导致错误，或者发生HTTP协议错误或IO错误，将会返回错误。
//
// 对于非阻塞调用，或需要控制HTTP客户端头信息、重定向策略和其他设置的情况，请创建一个Client实例代替。
func Get(dst, urlStr string) (*Response, error) {
	req, err := NewRequest(dst, urlStr)
	if err != nil {
		return nil, err
	}

	resp := DefaultClient.Do(req)
	return resp, resp.Err()
}

// GetBatch 发送多个 HTTP 请求，并使用指定数量的并发工作 goroutine 将请求的 URL 内容下载到给定的目标目录。
//
// 当一个工作线程从远程服务器接收到响应时，每个请求的 URL 的 Response 会立即通过返回的 Response 通道发送出去。这样，在下载过程中就可以利用这个 Response 来跟踪下载进度。
//
// 只有当所有下载任务完成或失败后，由 Grab 返回的 Response 通道才会被关闭。
//
// 如果在任何下载过程中发生错误，可以通过调用相关的 Response.Err 来获取该错误信息。
//
// 如果需要控制 HTTP 客户端头、重定向策略以及其他设置，请创建一个 Client 对象来代替。
func GetBatch(workers int, dst string, urlStrs ...string) (<-chan *Response, error) {
	fi, err := os.Stat(dst)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("destination is not a directory")
	}

	reqs := make([]*Request, len(urlStrs))
	for i := 0; i < len(urlStrs); i++ {
		req, err := NewRequest(dst, urlStrs[i])
		if err != nil {
			return nil, err
		}
		reqs[i] = req
	}

	ch := DefaultClient.DoBatch(workers, reqs...)
	return ch, nil
}
