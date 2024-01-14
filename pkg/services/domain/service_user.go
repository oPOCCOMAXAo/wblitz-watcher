package domain

import (
	"context"
	"strings"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

type UserStatsRequest struct {
	ServerID string
	Username string
	Region   models.Region
}

func (s *Service) UserStats(
	ctx context.Context,
	request *UserStatsRequest,
) (*wg.AccountInfo, error) {
	list, err := s.wg.AccountList(ctx, wg.AccountListRequest{
		Search: request.Username,
		Region: request.Region,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	if len(list) == 0 {
		return nil, models.ErrNotFound
	}

	var id int64

	request.Username = strings.ToLower(request.Username)

	for _, user := range list {
		if strings.ToLower(user.Nickname) == request.Username {
			id = user.AccountID

			break
		}
	}

	if id == 0 {
		return nil, models.ErrNotFound
	}

	info, err := s.wg.AccountInfo(ctx, wg.AccountInfoRequest{
		Region: request.Region,
		IDs:    []int64{id},
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	account := info[jsonutils.MaybeInt(id)]
	if account == nil {
		return nil, models.ErrNotFound
	}

	return account, nil
}
