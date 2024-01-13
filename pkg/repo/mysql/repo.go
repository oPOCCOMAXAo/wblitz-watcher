package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type Repository struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}

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

func (r *Repository) CreateSubscriptionClan(
	ctx context.Context,
	value *models.SubscriptionClan,
) error {
	if value.ID != 0 {
		return errors.WithStack(models.ErrAlreadyExists)
	}

	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO subscription_clan (instance_id, clan_id, region, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	now := time.Now().Unix()

	res, err := stmt.ExecContext(ctx,
		value.InstanceID,
		value.ClanID,
		value.Region,
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

func (r *Repository) GetSubscriptionClan(
	ctx context.Context,
	value *models.SubscriptionClan,
) (*models.SubscriptionClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, instance_id, clan_id, region
FROM subscription_clan
WHERE instance_id = ? AND clan_id = ? AND region = ?`,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer stmt.Close()

	var res models.SubscriptionClan

	row := stmt.QueryRowContext(ctx,
		value.InstanceID,
		value.ClanID,
		value.Region,
	)

	err = row.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = row.
		Scan(
			&res.ID,
			&res.InstanceID,
			&res.ClanID,
			&res.Region,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(models.ErrNotFound)
		}

		return nil, errors.WithStack(err)
	}

	return &res, nil
}

func (r *Repository) UpdateSubscriptionClan(
	ctx context.Context,
	value *models.SubscriptionClan,
) error {
	if value.ID == 0 {
		return errors.WithStack(models.ErrNotFound)
	}

	stmt, err := r.db.PrepareContext(ctx,
		`UPDATE subscription_clan
SET instance_id = ?, clan_id = ?, region = ?, updated_at = ?
WHERE id = ?`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	now := time.Now().Unix()

	_, err = stmt.ExecContext(ctx,
		value.InstanceID,
		value.ClanID,
		value.Region,
		now,
		value.ID,
	)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) DeleteSubscriptionClan(
	ctx context.Context,
	value *models.SubscriptionClan,
) error {
	stmt, err := r.db.PrepareContext(ctx,
		`DELETE FROM subscription_clan
WHERE instance_id = ? AND clan_id = ? AND region = ?`,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		value.InstanceID,
		value.ClanID,
		value.Region,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
