package wg

import (
	"context"
	"net/url"

	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

type ClansInfoRequest struct {
	Region models.Region
	IDs    []int64
}

func (c *Client) ClansInfo(
	ctx context.Context,
	request ClansInfoRequest,
) (map[int64]*ClanInfo, error) {
	tempRes := map[jsonutils.MaybeInt]*ClanInfo{}

	for _, ids := range lo.Chunk(request.IDs, 100) {
		err := c.Request(ctx, &Request{
			Region: request.Region,
			App:    AppWotBlitz,
			Method: MethodClansInfo,
			Data: url.Values{
				"clan_id": {JoinInt64(ids, ",")},
			},
		}, &tempRes)
		if err != nil {
			return nil, err
		}
	}

	res := make(map[int64]*ClanInfo, len(tempRes))

	for _, clan := range tempRes {
		if clan == nil {
			continue
		}

		res[clan.ClanID] = clan
		clan.Region = request.Region
	}

	return res, nil
}

func (c *Client) GetClanByID(
	ctx context.Context,
	id models.WGClanID,
) (*ClanInfo, error) {
	res, err := c.ClansInfo(ctx, ClansInfoRequest{
		Region: id.Region,
		IDs:    []int64{id.ID},
	})
	if err != nil {
		return nil, err
	}

	clan, ok := res[id.ID]
	if !ok {
		return nil, models.ErrNotFound
	}

	return clan, nil
}
