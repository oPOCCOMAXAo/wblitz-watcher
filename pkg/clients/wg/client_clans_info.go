package wg

import (
	"context"
	"net/url"

	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

type ClansInfoRequest struct {
	Region Region
	IDs    []int64
}

func (c *Client) ClansInfo(
	ctx context.Context,
	request ClansInfoRequest,
) (map[jsonutils.MaybeInt]*ClanInfo, error) {
	res := map[jsonutils.MaybeInt]*ClanInfo{}

	for _, ids := range lo.Chunk(request.IDs, 100) {
		err := c.Request(ctx, &Request{
			Region: request.Region,
			App:    AppWotBlitz,
			Method: MethodClansInfo,
			Data: url.Values{
				"clan_id": {JoinInt64(ids, ",")},
			},
		}, &res)
		if err != nil {
			return nil, err
		}
	}

	for _, clan := range res {
		clan.Region = request.Region
	}

	return res, nil
}
