package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/diff"
)

func (s *Service) EnsureWGClan(
	ctx context.Context,
	value *models.WGClan,
) error {
	err := s.repo.CreateUpdateWGClan(ctx, value)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (s *Service) GetWGClansForProcessing(
	ctx context.Context,
) ([]*models.WGClanExtended, error) {
	values, err := s.repo.GetWGClansWithSubscriptions(ctx)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	res := make([]*models.WGClanExtended, len(values))
	ids := make([]models.WGClanID, len(values))
	byID := make(map[models.WGClanID]*models.WGClanExtended, len(values))

	for i, value := range values {
		id := value.GetFullClanID()
		item := &models.WGClanExtended{
			Clan: value,
		}

		res[i] = item
		ids[i] = id
		byID[id] = item
	}

	allMembers, err := s.repo.GetWGClanMembers(ctx, ids)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	for _, members := range allMembers {
		clan := byID[members.ID]
		if clan == nil {
			continue
		}

		clan.MembersIDs = members.MembersIDs
	}

	return res, nil
}

func (s *Service) GetClanMembersFromWG(
	ctx context.Context,
	clans []models.WGClanID,
) ([]*models.WGClanMembers, error) {
	byRegion := map[models.Region][]int64{}
	for _, clan := range clans {
		byRegion[clan.Region] = append(byRegion[clan.Region], clan.ID)
	}

	res := make([]*models.WGClanMembers, len(clans))
	byClanID := map[models.WGClanID]*models.WGClanMembers{}

	for i, clanID := range clans {
		res[i] = &models.WGClanMembers{
			ID:         clanID,
			MembersIDs: []int64{},
			IsFound:    false,
		}

		byClanID[clanID] = res[i]
	}

	for region, ids := range byRegion {
		clansInfo, err := s.wg.ClansInfo(ctx, wg.ClansInfoRequest{
			Region: region,
			IDs:    ids,
		})
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		for _, clanInfo := range clansInfo {
			members := byClanID[clanInfo.WGClanID()]
			if members == nil {
				continue
			}

			members.MembersIDs = clanInfo.MembersIDs
			members.IsFound = true
		}
	}

	return res, nil
}

func (s *Service) UpdateWGClanMembers(
	ctx context.Context,
	values *diff.Diff[*models.WGClanMember],
) error {
	err := s.repo.CreateUpdateWGClanMembers(ctx,
		append(values.Created, values.Updated...),
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.DeleteWGClanMembers(ctx, values.Deleted)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (s *Service) UpdateWGClansMembersUpdateTime(
	ctx context.Context,
	clans []models.WGClanID,
) error {
	err := s.repo.UpdateWGClansMembersUpdatedAt(ctx, s.now(), clans)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
