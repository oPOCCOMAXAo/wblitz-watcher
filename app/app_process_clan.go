package app

import (
	"context"
	"log"
	"wblitz-watcher/app/diff"
	"wblitz-watcher/wg/types"

	"github.com/pkg/errors"
)

func (app *App) ProcessClans(ctx context.Context) error {
	clansAPI, err := app.client.ClansInfo(ctx, app.config.Region, app.config.Clans...)
	if err != nil {
		return errors.WithStack(err)
	}

	regionName := app.config.Region.Name()

	clansDB, err := app.repo.GetClans(ctx, regionName, app.config.Clans)
	if err != nil {
		return errors.WithStack(err)
	}

	clansDBMap := map[int]*types.ClanInfo{}
	for _, clan := range clansDB {
		clansDBMap[clan.ClanID] = clan
	}

	for _, id := range app.config.Clans {
		newClan := clansAPI[types.MaybeInt(id)]
		diff := DiffClan(clansDBMap[id], newClan)

		if diff.Len() > 0 {
			err = app.NotifyClanDiff(ctx, newClan, &diff)
			if err != nil {
				return errors.WithStack(err)
			}

			newClan.Region = regionName

			err = app.repo.SaveClan(ctx, newClan)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}
	}

	return nil
}

func DiffClan(oldClan, newClan *types.ClanInfo) diff.Total {
	res := diff.Total{}

	if newClan == nil {
		return res
	}

	if oldClan == nil {
		res.Void = append(res.Void, diff.Diff[diff.Void]{Type: DiffCreated})

		return res
	}

	diff.DetectSingleValue(DiffLeader, oldClan.LeaderID, newClan.LeaderID, &res.Int)
	diff.DetectSingleValue(DiffName, oldClan.Name, newClan.Name, &res.String)
	diff.DetectSingleValue(DiffTag, oldClan.Tag, newClan.Tag, &res.String)
	diff.DetectSetNew(DiffEnter, oldClan.MembersIDs, newClan.MembersIDs, &res.Int)
	diff.DetectSetNew(DiffLeave, newClan.MembersIDs, oldClan.MembersIDs, &res.Int)

	return res
}
