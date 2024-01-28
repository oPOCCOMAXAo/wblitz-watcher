package wg

import (
	"context"
	"net/url"
	"strings"

	"github.com/pkg/errors"

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

func (c *Client) FindAccountByName(
	ctx context.Context,
	request AccountListRequest,
) (*AccountListEntry, error) {
	cmpStr := strings.ToLower(request.Search)

	res, err := c.AccountList(ctx, request)
	if err != nil {
		return nil, err
	}

	for _, account := range res {
		if strings.ToLower(account.Nickname) == cmpStr {
			return account, nil
		}
	}

	return nil, errors.Wrap(ErrLimitExceeded, "account not found")
}
