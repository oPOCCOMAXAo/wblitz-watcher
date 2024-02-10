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

//nolint:funlen
func (s *Service) taskWatchClan(
	ctx context.Context,
) error {
	clans, err := s.domain.GetWGClansForProcessing(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	notifyRequired := false
	defer func() {
		if notifyRequired {
			s.notifyTaskProcessEvents()
		}
	}()

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

	updatedIDs := make([]models.WGClanID, 0, len(ids))

	defer func() {
		err := s.domain.UpdateWGClansMembersUpdateTime(ctx, updatedIDs)
		if err != nil {
			telemetry.RecordError(ctx, err)
		}
	}()

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

		idDiff := diff.Slice(
			apiClan.MembersIDs,
			dbClan.MembersIDs,
			diff.Ints.GetUniqueID,
			diff.Ints.PrepareToUpdate,
			diff.Ints.Equals,
		)

		err = s.updateClanByIDDiff(ctx, dbClan, &idDiff, &notifyRequired)
		if err != nil {
			return err
		}

		updatedIDs = append(updatedIDs, apiClan.ID)
	}

	return nil
}

func (s *Service) updateClanByIDDiff(
	ctx context.Context,
	dbClan *models.WGClanExtended,
	idDiff *diff.Diff[int64],
	notifyRequired *bool,
) error {
	if idDiff.IsEmpty() {
		return nil
	}

	mapper := dbClan.Clan.GetFullClanID().MapMember

	memDiff := diff.Diff[*models.WGClanMember]{
		Created: lo.Map(idDiff.Created, mapper),
		Updated: lo.Map(idDiff.Updated, mapper),
		Deleted: lo.Map(idDiff.Deleted, mapper),
	}

	if s.isClanMembersInitialized(dbClan.Clan) {
		total, err := s.domain.CreateEventClanByMembersDiff(ctx, &memDiff)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		if total > 0 {
			*notifyRequired = true
		}
	}

	err := s.domain.UpdateWGClanMembers(ctx, &memDiff)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
