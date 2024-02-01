package watcher

import (
	"context"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/diff"
)

func (s *Service) TaskWatchClan(
	ctx context.Context,
) error {
	ctx, span := s.taskTracer.Start(ctx, "watchClan")
	defer span.End()

	err := s.taskWatchClan(ctx)
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}

	return err
}

//nolint:funlen,cyclop
func (s *Service) taskWatchClan(
	ctx context.Context,
) error {
	clans, err := s.domain.GetWGClansForProcessing(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	dbClanByID := make(map[models.WGClanID]*models.WGClanExtended, len(clans))
	ids := make([]models.WGClanID, len(clans))

	for i, clan := range clans {
		id := clan.Clan.GetFullClanID()
		ids[i] = id
		dbClanByID[id] = clan
	}

	allClansMembers, err := s.domain.GetClanMembersFromWG(ctx, ids)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	for _, apiClan := range allClansMembers {
		if !apiClan.IsFound {
			telemetry.RecordErrorFail(
				ctx,
				errors.WithMessagef(models.ErrFlowBroken, "%s was not found", apiClan.EntityUniqueID()),
			)

			continue
		}

		dbClan := dbClanByID[apiClan.ID]
		if dbClan == nil {
			telemetry.RecordErrorFail(
				ctx,
				errors.WithMessagef(models.ErrFlowBroken, "%s was not in DB", apiClan.EntityUniqueID()),
			)

			dbClan = &models.WGClanExtended{
				Clan: &models.WGClan{
					ID:     apiClan.ID.ID,
					Region: apiClan.ID.Region,
				},
			}
		}

		idDiff := diff.Calculate(
			apiClan.MembersIDs,
			dbClan.MembersIDs,
			diff.Ints.GetUniqueID,
			diff.Ints.PrepareToUpdate,
		)

		if idDiff.IsEmpty() {
			continue
		}

		memDiff := diff.Diff[*models.WGClanMember]{
			Created: lo.Map(idDiff.Created, apiClan.ID.MapMember),
			Updated: lo.Map(idDiff.Updated, apiClan.ID.MapMember),
			Deleted: lo.Map(idDiff.Deleted, apiClan.ID.MapMember),
		}

		if s.isClanMembersInitialized(dbClan.Clan) {
			err = s.domain.CreateEventClanByMembersDiff(ctx, &memDiff)
			if err != nil {
				//nolint:wrapcheck
				return err
			}

			s.notifyTaskProcessEvents()
		}

		err = s.domain.UpdateWGClanMembers(ctx, &memDiff)
		if err != nil {
			//nolint:wrapcheck
			return err
		}
	}

	return nil
}
