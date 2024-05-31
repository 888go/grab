package grabui

import (
	"context"

	"github.com/888go/grab/v3"
)

func GetBatch(
	ctx context.Context,
	workers int,
	dst string,
	urlStrs ...string,
) (<-chan *下载类.Response, error) {
	reqs := make([]*下载类.Request, len(urlStrs))
	for i := 0; i < len(urlStrs); i++ {
		req, err := 下载类.X生成下载参数(dst, urlStrs[i])
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		reqs[i] = req
	}

	ui := NewConsoleClient(下载类.DefaultClient)
	return ui.Do(ctx, workers, reqs...), nil
}
