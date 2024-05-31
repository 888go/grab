package main

import (
	"context"
	"fmt"
	"os"

	"github.com/888go/grab/v3/pkg/grabui"
)

func main() {
	// validate command args
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

	// 作为退出代码，返回失败下载的数量 md5:b33b9e56ad5e3d93
	failed := 0
	for resp := range respch {
		if resp.X等待错误() != nil {
			failed++
		}
	}
	os.Exit(failed)
}
