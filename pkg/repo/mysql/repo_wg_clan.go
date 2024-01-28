package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateUpdateWGClan(
	ctx context.Context,
	value *models.WGClan,
) error {
	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO wg_clan (id, region, tag, name, updated_at)
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE tag = VALUES(tag), name = VALUES(name), updated_at = VALUES(updated_at)`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	now := time.Now().Unix()

	_, err = stmt.ExecContext(ctx,
		value.ID,
		value.Region,
		value.Tag,
		value.Name,
		now,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GetWGClan implements repo.Repository.
func (r *Repository) GetWGClan(
	ctx context.Context,
	value *models.WGClan,
) (*models.WGClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, region, tag, name, updated_at, members_updated_at
FROM wg_clan
WHERE id = ? AND region = ?`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	var res models.WGClan

	row := stmt.QueryRowContext(ctx,
		value.ID,
		value.Region,
	)

	err = row.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = row.Scan(
		&res.ID,
		&res.Region,
		&res.Tag,
		&res.Name,
		&res.UpdatedAt,
		&res.MembersUpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(models.ErrNotFound)
		}

		return nil, errors.WithStack(err)
	}

	return &res, nil
}

func (r *Repository) UpdateWGClansMembersUpdatedAt(
	ctx context.Context,
	updatedAt int64,
	ids []models.WGClanID,
) error {
	if len(ids) == 0 {
		return nil
	}

	//nolint:gosec // here placeholders are safe.
	sql := `UPDATE wg_clan
SET members_updated_at = ?
WHERE (id, region) IN (` + r.placeholdersGroup(len(ids), 2) + `)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	args := make([]any, 0, len(ids)*2+1)
	args = append(args, updatedAt)

	for _, id := range ids {
		args = append(args, id.ID, id.Region)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) GetWGClanListByInstance(
	ctx context.Context,
	instanceID int64,
) ([]*models.WGClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT wc.id, wc.region, wc.tag, wc.name, wc.updated_at, wc.members_updated_at
FROM subscription_clan sc
JOIN wg_clan wc ON wc.id = sc.clan_id AND wc.region = sc.region
WHERE sc.instance_id = ?
ORDER BY sc.id ASC`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx,
		instanceID,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var res []*models.WGClan

	for rows.Next() {
		var item models.WGClan

		err = rows.Scan(
			&item.ID,
			&item.Region,
			&item.Tag,
			&item.Name,
			&item.UpdatedAt,
			&item.MembersUpdatedAt,
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res = append(res, &item)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (r *Repository) GetWGClansWithSubscriptions(
	ctx context.Context,
) ([]*models.WGClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT wc.id, wc.region, wc.tag, wc.name, wc.updated_at, wc.members_updated_at
FROM wg_clan wc
JOIN subscription_clan sc ON sc.clan_id = wc.id AND sc.region = wc.region
GROUP BY wc.id, wc.region`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var res []*models.WGClan

	for rows.Next() {
		var item models.WGClan

		err = rows.Scan(
			&item.ID,
			&item.Region,
			&item.Tag,
			&item.Name,
			&item.UpdatedAt,
			&item.MembersUpdatedAt,
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res = append(res, &item)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
