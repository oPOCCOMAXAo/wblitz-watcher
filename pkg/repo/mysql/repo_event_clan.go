package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateEventClan(
	ctx context.Context,
	values ...*models.EventClan,
) error {
	if len(values) == 0 {
		return nil
	}

	//nolint:gosec // here placeholders are safe.
	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO event_clan (time, type, region, clan_id, account_id)
VALUES `+r.placeholdersGroup(len(values), 5),
	)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	defer stmt.Close()

	now := time.Now().Unix()

	args := make([]any, 0, len(values)*5)
	for _, value := range values {
		args = append(args,
			now,
			value.Type,
			value.Region,
			value.ClanID,
			value.AccountID,
		)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (r *Repository) DeleteEventClansByID(
	ctx context.Context,
	ids []int64,
) error {
	if len(ids) == 0 {
		return nil
	}

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

//go:embed sql/update_event_clan_processed.sql
var sqlUpdateEventClanProcessed string

func (r *Repository) UpdateEventClanProcessed(
	ctx context.Context,
) error {
	_, err := r.db.ExecContext(ctx, sqlUpdateEventClanProcessed)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
