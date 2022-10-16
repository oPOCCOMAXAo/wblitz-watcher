package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/opoccomaxao/wblitz-watcher/wg/api"
	"github.com/opoccomaxao/wblitz-watcher/wg/types"

	"github.com/opoccomaxao-go/generic-collection/slice"
)

func (c *Client) ClansInfo(
	ctx context.Context,
	region api.Region,
	ids ...int,
) (map[types.MaybeInt]*types.ClanInfo, error) {
	res := map[types.MaybeInt]*types.ClanInfo{}

	for _, ids := range slice.Chunk(ids, 100) {
		err := c.api.Request(ctx, &api.Request{
			Region: region,
			App:    api.AppWotBlitz,
			Method: api.MethodClansInfo,
			Data: url.Values{
				"clan_id": []string{slice.Join(ids, ",", strconv.Itoa)},
			},
		}, &res)
		if err != nil {
			return nil, err
		}
	}

	for _, clan := range res {
		clan.Region = region
	}

	return res, nil
}
