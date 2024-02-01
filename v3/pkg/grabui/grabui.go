package grabui

import (
	"context"

	"github.com/cavaliergopher/grab/v3"
)

func GetBatch(
	ctx context.Context,
	workers int,
	dst string,
	urlStrs ...string,
) (<-chan *下载类.X响应, error) {
	reqs := make([]*下载类.X下载参数, len(urlStrs))
	for i := 0; i < len(urlStrs); i++ {
		req, err := 下载类.X生成下载参数(dst, urlStrs[i])
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		reqs[i] = req
	}

	ui := NewConsoleClient(下载类.X默认全局客户端)
	return ui.Do(ctx, workers, reqs...), nil
}
