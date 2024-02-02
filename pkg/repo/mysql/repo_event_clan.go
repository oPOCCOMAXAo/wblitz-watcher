package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (r *Repository) CreateEventClan(
	ctx context.Context,
	values ...*models.EventClan,
) error {
	if len(values) == 0 {
		return nil
	}

	stmt, err := r.db.PrepareContext(ctx,
		`INSERT INTO event_clan (time, type, region, clan_id, account_id)
VALUES `+r.placeholdersGroup(len(values), 5),
	)
	if err != nil {
		return errors.WithStack(err)
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
		return errors.WithStack(err)
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

	sql := `DELETE FROM event_clan WHERE id IN (` + r.placeholders(len(ids)) + `)`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	args := make([]any, 0, len(ids))
	for _, id := range ids {
		args = append(args, id)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.WithStack(err)
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
		return nil, errors.WithStack(err)
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
			return nil, errors.WithStack(models.ErrNotFound)
		}

		return nil, errors.WithStack(err)
	}

	return &value, nil
}

//go:embed sql/update_event_clan_processed.sql
var sqlUpdateEventClanProcessed string

func (r *Repository) UpdateEventClanProcessed(
	ctx context.Context,
) (int64, error) {
	res, err := r.db.ExecContext(ctx, sqlUpdateEventClanProcessed)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	total, err := res.RowsAffected()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return total, nil
}
