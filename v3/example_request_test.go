package 下载类

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func ExampleRequest_WithContext() {
	// 使用100毫秒超时创建上下文 md5:20d987ef3ab9f9e9
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 使用上下文创建下载请求 md5:cd71b43cad23f3e2
	req, err := X生成下载参数("", "http://example.com/example.zip")
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)

	// send download request
	resp := DefaultClient.X下载(req)
	if err := resp.X等待错误(); err != nil {
		fmt.Println("error: request cancelled")
	}

	// Output:
	// error: request cancelled
}

func ExampleRequest_SetChecksum() {
	// create download request
	req, err := X生成下载参数("", "http://example.com/example.zip")
	if err != nil {
		panic(err)
	}

	// set request checksum
	sum, err := hex.DecodeString("33daf4c03f86120fdfdc66bddf6bfff4661c7ca11c5da473e537f4d69b470e57")
	if err != nil {
		panic(err)
	}
	req.X设置完成后效验(sha256.New(), sum, true)

	// 下载并验证文件 md5:04a19a87a53ccecb
	resp := DefaultClient.X下载(req)
	if err := resp.X等待错误(); err != nil {
		panic(err)
	}
}
