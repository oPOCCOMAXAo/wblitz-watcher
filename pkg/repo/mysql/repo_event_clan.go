package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateEventClan(
	ctx context.Context,
	value *models.EventClan,
) error {
	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO event_clan (time, type, region, clan_id, account_id)
VALUES (?, ?, ?, ?, ?)`,
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	now := time.Now().Unix()

	res, err := stmt.ExecContext(ctx,
		now,
		value.Type,
		value.Region,
		value.ClanID,
		value.AccountID,
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	value.ID = id

	return nil
}

func (r *Repository) DeleteEventClansByID(
	ctx context.Context,
	ids []int64,
) error {
	//nolint:gosec // here placeholders are safe.
	sql := `DELETE FROM event_clan WHERE id IN (` + r.placeholders(len(ids)) + `)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	args := make([]any, 0, len(ids))
	for _, id := range ids {
		args = append(args, id)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (r *Repository) GetEventClanByID(
	ctx context.Context,
	id int64,
) (*models.EventClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, time, type, region, clan_id, account_id, is_processed
FROM event_clan
WHERE id = ?`,
	)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	defer stmt.Close()

	var value models.EventClan

	err = stmt.
		QueryRowContext(ctx, id).
		Scan(
			&value.ID,
			&value.Time,
			&value.Type,
			&value.Region,
			&value.ClanID,
			&value.AccountID,
			&value.IsProcessed,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotFound
		}

		//nolint:wrapcheck
		return nil, err
	}

	return &value, nil
}

func (r *Repository) GetEventClanForProcessing(
	ctx context.Context,
) (*models.EventClan, error) {
	stmt, err := r.db.PrepareContext(ctx,
		`SELECT id, time, type, region, clan_id, account_id, is_processed
FROM event_clan
WHERE is_processed = 0
LIMIT 1`,
	)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	defer stmt.Close()

	var value models.EventClan

	err = stmt.
		QueryRowContext(ctx).
		Scan(
			&value.ID,
			&value.Time,
			&value.Type,
			&value.Region,
			&value.ClanID,
			&value.AccountID,
			&value.IsProcessed,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotFound
		}

		//nolint:wrapcheck
		return nil, err
	}

	return &value, nil
}

func (r *Repository) UpdateEventClanProcessed(
	ctx context.Context,
	value *models.EventClan,
) error {
	stmt, err := r.db.PrepareContext(ctx,
		`UPDATE event_clan
SET is_processed = ?
WHERE id = ?`,
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		value.IsProcessed,
		value.ID,
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
