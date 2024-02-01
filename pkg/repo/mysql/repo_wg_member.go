package mysql

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateUpdateWGClanMembers(
	ctx context.Context,
	values []*models.WGClanMember,
) error {
	if len(values) == 0 {
		return nil
	}

	sql := `INSERT INTO wg_clan_member (region, clan_id, account_id) VALUES ` +
		r.placeholdersGroup(len(values), 3) +
		` ON DUPLICATE KEY UPDATE region = VALUES(region), clan_id = VALUES(clan_id)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	args := make([]any, 0, len(values)*3)
	for _, value := range values {
		args = append(args, value.Region, value.ClanID, value.AccountID)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (r *Repository) DeleteWGClanMembers(
	ctx context.Context,
	values []*models.WGClanMember,
) error {
	if len(values) == 0 {
		return nil
	}

	sql := `DELETE FROM wg_clan_member WHERE (region, clan_id, account_id) IN (` +
		r.placeholdersGroup(len(values), 3) +
		`)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	args := make([]any, 0, len(values)*3)
	for _, value := range values {
		args = append(args, value.Region, value.ClanID, value.AccountID)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

//nolint:funlen
func (r *Repository) GetWGClanMembers(
	ctx context.Context,
	ids []models.WGClanID,
) ([]*models.WGClanMembers, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	sql := `SELECT region, clan_id, account_id FROM wg_clan_member WHERE (clan_id, region) IN (` +
		r.placeholdersGroup(len(ids), 2) +
		`)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	defer stmt.Close()

	args := make([]any, 0, len(ids)*2)
	for _, id := range ids {
		args = append(args, id.ID, id.Region)
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	defer rows.Close()

	res := make([]*models.WGClanMembers, 0, len(ids))
	mapByID := make(map[models.WGClanID]*models.WGClanMembers, len(ids))

	var (
		id        models.WGClanID
		accountID int64
	)

	for rows.Next() {
		err = rows.Scan(
			&id.Region,
			&id.ID,
			&accountID,
		)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		members, ok := mapByID[id]
		if !ok {
			members = &models.WGClanMembers{
				ID: id,
			}

			res = append(res, members)
			mapByID[id] = members
		}

		members.MembersIDs = append(members.MembersIDs, accountID)
	}

	err = rows.Err()
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return res, nil
}
