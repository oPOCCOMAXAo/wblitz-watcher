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
