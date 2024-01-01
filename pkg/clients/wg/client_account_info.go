package wg

import (
	"context"
	"net/url"

	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

type AccountInfoRequest struct {
	Region Region
	IDs    []int64
}

func (c *Client) AccountInfo(
	ctx context.Context,
	request AccountInfoRequest,
) (map[jsonutils.MaybeInt]*AccountInfo, error) {
	res := map[jsonutils.MaybeInt]*AccountInfo{}

	for _, ids := range lo.Chunk(request.IDs, 100) {
		err := c.Request(ctx, &Request{
			Region: request.Region,
			App:    AppWotBlitz,
			Method: MethodAccoutInfo,
			Data: url.Values{
				"account_id": {JoinInt64(ids, ",")},
			},
		}, &res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
