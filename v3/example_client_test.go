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

// 此示例使用 DoChannel 创建一个用于并发下载多个文件的生产者/消费者模型。这与 DoBatch 在底层使用 DoChannel 的方式类似，但不同之处在于它允许调用者持续发送新的请求，直到他们希望关闭请求通道为止。
func ExampleClient_DoChannel() {
// 创建一个请求和一个缓冲的响应通道
	reqch := make(chan *X下载参数)
	respch := make(chan *X响应, 10)

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

// 等待所有工作者完成
		wg.Wait()
		close(respch)
	}()

// 检查每个响应
	for resp := range respch {
// 等待直到完成
		if err := resp.X等待错误(); err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded %s to %s\n", resp.X下载参数.X取下载链接(), resp.X文件名)
	}
}

func ExampleClient_DoBatch() {
// 创建多个下载请求
	reqs := make([]*X下载参数, 0)
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("http://example.com/example%d.zip", i+1)
		req, err := X生成下载参数("/tmp", url)
		if err != nil {
			panic(err)
		}
		reqs = append(reqs, req)
	}

// 使用4个工人开始下载
	client := X创建客户端()
	respch := client.X多线程下载(4, reqs...)

// 检查每个响应
	for resp := range respch {
		if err := resp.X等待错误(); err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded %s to %s\n", resp.X下载参数.X取下载链接(), resp.X文件名)
	}
}
