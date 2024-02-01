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
func X下载(保存目录, 下载链接 string) (*X响应, error) {
	req, err := X生成下载参数(保存目录, 下载链接)
	if err != nil {
		return nil, err
	}

	resp := X默认全局客户端.X下载(req)
	return resp, resp.X等待错误()
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
func X多线程下载(线程数 int, 保存目录 string, 下载链接 ...string) (<-chan *X响应, error) {
	fi, err := os.Stat(保存目录)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("Destination不是目录")
	}

	reqs := make([]*X下载参数, len(下载链接))
	for i := 0; i < len(下载链接); i++ {
		req, err := X生成下载参数(保存目录, 下载链接[i])
		if err != nil {
			return nil, err
		}
		reqs[i] = req
	}

	ch := X默认全局客户端.X多线程下载(线程数, reqs...)
	return ch, nil
}
