package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateInstance(
	ctx context.Context,
	value *models.BotInstance,
) error {
	if value.ID != 0 {
		return errors.WithStack(models.ErrAlreadyExists)
	}

	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO bot_instance (server_id, channel_id, type, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return errors.WithStack(err)
	}
	defer stmt.Close()

	now := time.Now().Unix()

	res, err := stmt.ExecContext(ctx,
		value.ServerID,
		value.ChannelID,
		value.Type,
		now,
		now,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	value.ID, err = res.LastInsertId()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) GetInstance(
	ctx context.Context,
	value *models.BotInstance,
) (*models.BotInstance, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, server_id, channel_id, type
FROM bot_instance
WHERE server_id = ? AND type = ?`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	var res models.BotInstance

	row := stmt.QueryRowContext(ctx,
		value.ServerID,
		value.Type,
	)

	err = row.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = row.
		Scan(
			&res.ID,
			&res.ServerID,
			&res.ChannelID,
			&res.Type,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(models.ErrNotFound)
		}

		return nil, errors.WithStack(err)
	}

	return &res, nil
}

func (r *Repository) UpdateInstance(
	ctx context.Context,
	value *models.BotInstance,
) error {
	if value.ID == 0 {
		return errors.WithStack(models.ErrNotFound)
	}

	stmt, err := r.db.PrepareContext(ctx,
		`UPDATE bot_instance
SET server_id = ?, channel_id = ?, type = ?, updated_at = ?
WHERE id = ?`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	now := time.Now().Unix()

	_, err = stmt.ExecContext(ctx,
		value.ServerID,
		value.ChannelID,
		value.Type,
		now,
		value.ID,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) GetInstancesByType(
	ctx context.Context,
	value models.SubscriptionType,
) ([]*models.BotInstance, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, server_id, channel_id, type
FROM bot_instance
WHERE type = ?`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx,
		value,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var res []*models.BotInstance

	for rows.Next() {
		var item models.BotInstance

		err = rows.
			Scan(
				&item.ID,
				&item.ServerID,
				&item.ChannelID,
				&item.Type,
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

func (r *Repository) SoftDeleteInstancesByServer(
	ctx context.Context,
	value []string,
) error {
	if len(value) == 0 {
		return nil
	}

	stmt, err := r.db.PrepareContext(ctx,
		`UPDATE bot_instance
SET updated_at = ?, deleted_at = ?
WHERE server_id IN (`+r.placeholders(len(value))+`)`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	now := time.Now().Unix()

	args := make([]any, 0, len(value)+2)
	args = append(args, now, now)

	for _, v := range value {
		args = append(args, v)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) HardDeleteCleanedInstances(
	ctx context.Context,
) error {
	// as foreign key is not cascade, we not delete instances with related rows in other tables.
	_, err := r.db.ExecContext(ctx,
		`DELETE IGNORE
FROM bot_instance
WHERE deleted_at > 0`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) GetNonDeletedInstancesServers(
	ctx context.Context,
) ([]string, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT DISTINCT server_id
FROM bot_instance
WHERE deleted_at = 0`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var res []string

	for rows.Next() {
		var item string

		err = rows.Scan(&item)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res = append(res, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
