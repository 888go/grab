package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cavaliergopher/grab/v3/pkg/grabui"
)

func main() {
// 验证命令参数
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url...\n", os.Args[0])
		os.Exit(1)
	}
	urls := os.Args[1:]

	// download files
	respch, err := grabui.GetBatch(context.Background(), 0, ".", urls...)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

// 返回失败的下载数量作为退出代码
	failed := 0
	for resp := range respch {
		if resp.Err() != nil {
			failed++
		}
	}
	os.Exit(failed)
}
