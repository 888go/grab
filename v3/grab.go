package 下载类

import (
	"fmt"
	"os"
)

// Get 发送一个HTTP请求，并将请求的URL内容下载到给定的目标文件路径。调用者会在下载完成，无论成功还是失败时被阻塞。
// 
// 如果由于客户端策略（如CheckRedirect）导致错误，或者出现HTTP协议或IO错误，会返回一个错误。
// 
// 若要进行非阻塞调用，或者控制HTTP客户端头、重定向策略和其他设置，请创建一个Client。
// md5:1bbde4e6e89ceb15
func X下载(保存目录, 下载链接 string) (*Response, error) {
	req, err := X生成下载参数(保存目录, 下载链接)
	if err != nil {
		return nil, err
	}

	resp := DefaultClient.X下载(req)
	return resp, resp.X等待错误()
}

// GetBatch 发送多个HTTP请求，并使用给定的并发工作goroutine数量将请求的URL内容下载到指定的目标目录。
//
// 每个请求的URL的Response会通过返回的Response通道发送，一旦工作goroutine从远程服务器接收到响应。然后可以使用Response来跟踪下载进度。
//
// Grab会在所有下载完成或失败后关闭返回的Response通道。
//
// 如果在任何下载过程中发生错误，可以通过调用关联的Response.Err获取该错误。
//
// 若要控制HTTP客户端头部、重定向策略和其他设置，应创建一个Client。
// md5:cfa454826e483447
func X多线程下载(线程数 int, 保存目录 string, 下载链接 ...string) (<-chan *Response, error) {
	fi, err := os.Stat(保存目录)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("destination is not a directory")
	}

	reqs := make([]*Request, len(下载链接))
	for i := 0; i < len(下载链接); i++ {
		req, err := X生成下载参数(保存目录, 下载链接[i])
		if err != nil {
			return nil, err
		}
		reqs[i] = req
	}

	ch := DefaultClient.X多线程下载(线程数, reqs...)
	return ch, nil
}
