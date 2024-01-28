package wg

import (
	"context"
	"net/url"

	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

type AccountInfoRequest struct {
	Region models.Region
	IDs    []int64
}

func (c *Client) AccountInfo(
	ctx context.Context,
	request AccountInfoRequest,
) (map[int64]*AccountInfo, error) {
	tempRes := map[jsonutils.MaybeInt]*AccountInfo{}

	for _, ids := range lo.Chunk(request.IDs, 100) {
		err := c.Request(ctx, &Request{
			Region: request.Region,
			App:    AppWotBlitz,
			Method: MethodAccountInfo,
			Data: url.Values{
				"account_id": {JoinInt64(ids, ",")},
			},
		}, &tempRes)
		if err != nil {
			return nil, err
		}
	}

	res := make(map[int64]*AccountInfo, len(tempRes))

	for _, account := range tempRes {
		if account == nil {
			continue
		}

		res[account.AccountID] = account
	}

	return res, nil
}

func (c *Client) GetAccountByID(
	ctx context.Context,
	id models.WGAccountID,
) (*AccountInfo, error) {
	res, err := c.AccountInfo(ctx, AccountInfoRequest{
		Region: id.Region,
		IDs:    []int64{id.ID},
	})
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		return v, nil
	}

	return nil, models.ErrNotFound
}
