package wg

import (
	"context"
	"net/url"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type AccountListRequest struct {
	Region models.Region
	Search string
}

func (c *Client) AccountList(
	ctx context.Context,
	request AccountListRequest,
) ([]*AccountListEntry, error) {
	var res []*AccountListEntry

	err := c.Request(ctx, &Request{
		Region: request.Region,
		App:    AppWotBlitz,
		Method: MethodAccountList,
		Data: url.Values{
			"search": {request.Search},
		},
	}, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
