package 下载类

import (
	"fmt"
	"sync"
)

func ExampleClient_Do() {
	client := X创建客户端()
	req, err := X生成下载参数("/tmp", "http://example.com/example.zip")
	if err != nil {
		panic(err)
	}

	resp := client.X下载(req)
	if err := resp.X等待错误(); err != nil {
		panic(err)
	}

	fmt.Println("Download saved to", resp.X文件名)
}

// 此示例使用DoChannel创建一个生产者/消费者模型，用于并发下载多个文件。
// 这与DoBatch在内部使用DoChannel的工作方式类似，不同之处在于它允许调用者持续发送新请求，
// 直到他们希望关闭请求通道为止。
// md5:8cfd63343a82362c
func ExampleClient_DoChannel() {
	// 创建一个请求和一个缓冲的响应通道 md5:c4b7d02b204abf79
	reqch := make(chan *Request)
	respch := make(chan *Response, 10)

	// start 4 workers
	client := X创建客户端()
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			client.DoChannel(reqch, respch)
			wg.Done()
		}()
	}

	go func() {
		// send requests
		for i := 0; i < 10; i++ {
			url := fmt.Sprintf("http://example.com/example%d.zip", i+1)
			req, err := X生成下载参数("/tmp", url)
			if err != nil {
				panic(err)
			}
			reqch <- req
		}
		close(reqch)

		// 等待工人完成 md5:b758a6709f2c0bc0
		wg.Wait()
		close(respch)
	}()

	// check each response
	for resp := range respch {
		// block until complete
		if err := resp.X等待错误(); err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded %s to %s\n", resp.X下载参数.X取下载链接(), resp.X文件名)
	}
}

func ExampleClient_DoBatch() {
	// 创建多个下载请求 md5:763389fc7a54db5a
	reqs := make([]*Request, 0)
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("http://example.com/example%d.zip", i+1)
		req, err := X生成下载参数("/tmp", url)
		if err != nil {
			panic(err)
		}
		reqs = append(reqs, req)
	}

	// 使用4个工人开始下载 md5:f23bbcd916bb7cfc
	client := X创建客户端()
	respch := client.X多线程下载(4, reqs...)

	// check each response
	for resp := range respch {
		if err := resp.X等待错误(); err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded %s to %s\n", resp.X下载参数.X取下载链接(), resp.X文件名)
	}
}
