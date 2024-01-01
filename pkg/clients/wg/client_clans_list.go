package wg

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type ClansListRequest struct {
	Region models.Region
	Search string
	Page   int
	Limit  int
}

func (c *Client) ClansList(
	ctx context.Context,
	request ClansListRequest,
) ([]*ClanListEntry, error) {
	var res []*ClanListEntry

	err := c.Request(ctx, &Request{
		Region: request.Region,
		App:    AppWotBlitz,
		Method: MethodClansList,
		Data: url.Values{
			"search":  {request.Search},
			"page_no": {strconv.Itoa(request.Page)},
			"limit":   {strconv.Itoa(request.Limit)},
		},
	}, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) FindClanByTag(
	ctx context.Context,
	request ClansListRequest,
) (*ClanListEntry, error) {
	const maxPage = 10

	request.Limit = 100
	request.Page = 1

	cmpStr := strings.ToLower(request.Search)

	for request.Page <= maxPage {
		res, err := c.ClansList(ctx, request)
		if err != nil {
			return nil, err
		}

		for _, clan := range res {
			if strings.ToLower(clan.Tag) == cmpStr {
				return clan, nil
			}
		}

		request.Page++
	}

	return nil, errors.Wrap(ErrLimitExceeded, "clan not found")
}
