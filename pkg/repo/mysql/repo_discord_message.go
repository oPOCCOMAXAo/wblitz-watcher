package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

//go:embed sql/create_discord_messages_from_event_clan.sql
var sqlCreateDiscordMessagesFromEventClan string

func (r *Repository) CreateDiscordMessagesFromEventClan(
	ctx context.Context,
) (int64, error) {
	res, err := r.db.ExecContext(ctx, sqlCreateDiscordMessagesFromEventClan)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return rows, nil
}

//go:embed sql/get_first_unsent_dm.sql
var sqlGetFirstUnsentDiscordMessage string

func (r *Repository) GetFirstUnsentDiscordMessage(
	ctx context.Context,
) (*models.DiscordMessage, error) {
	var res models.DiscordMessage

	err := r.db.QueryRowContext(ctx, sqlGetFirstUnsentDiscordMessage).
		Scan(
			&res.ID,
			&res.IsProcessed,
			&res.EventClanID,
			&res.InstanceID,
			&res.EventClan.Time,
			&res.EventClan.Type,
			&res.EventClan.Region,
			&res.EventClan.ClanID,
			&res.EventClan.AccountID,
			&res.BotInstance.ChannelID,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &res, nil
}

func (r *Repository) UpdateDiscordMessagesProcessed(
	ctx context.Context,
	ids []int64,
) error {
	if len(ids) == 0 {
		return nil
	}

	stmt, err := r.db.PrepareContext(ctx, `UPDATE discord_message
SET is_processed = 1
WHERE id IN (`+r.placeholders(len(ids))+`)`)
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

func (r *Repository) DeleteDiscordMessagesForDeletedInstances(
	ctx context.Context,
) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM discord_message
WHERE bot_instance_id IN (
	SELECT id
	FROM bot_instance
	WHERE deleted_at != 0
)`)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) DeleteProcessedDiscordMessages(
	ctx context.Context,
) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM discord_message
WHERE is_processed = 1`)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
