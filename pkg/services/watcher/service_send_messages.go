package watcher

import (
	"context"
	"errors"
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) notifyTaskSendMessages() {
	select {
	case s.chanSendMessages <- struct{}{}:
	default:
	}
}

func (s *Service) TaskSendMessages(
	ctx context.Context,
) error {
	ctx, span := s.taskTracer.Start(ctx, "sendMessages")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			s.notifyTaskSendMessages()

			return nil
		default:
			err := s.taskSendMessageFirst(ctx)
			if err == nil {
				continue
			}

			if !errors.Is(err, models.ErrNotFound) {
				telemetry.RecordErrorFail(ctx, err)

				return err
			}

			return nil
		}
	}
}

func (s *Service) taskSendMessageFirst(
	ctx context.Context,
) error {
	data, err := s.domain.GetFirstUnsentDiscordMessage(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	res := discord.Response{}

	res.Embeds = append(res.Embeds, s.embedClanEvent(
		&data.Message.EventClan,
		data.Clan,
		data.User,
	))

	err = s.discord.SendMessage(ctx, data.Message.ChannelID, &res)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.domain.UpdateDiscordMessageProcessed(ctx, data.Message.ID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
