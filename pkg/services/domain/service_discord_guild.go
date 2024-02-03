package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/diff"
)

func (s *Service) DeleteDiscordGuildData(
	ctx context.Context,
	serverID string,
) error {
	err := s.repo.SoftDeleteInstancesByServer(ctx, []string{serverID})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.CleanupDeletedInstances(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteDiscordGuildsNotInList(
	ctx context.Context,
	currentServers []string,
) error {
	dbServers, err := s.repo.GetNonDeletedInstancesServers(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	diff := diff.Slice(
		currentServers,
		dbServers,
		diff.Strings.GetUniqueID,
		diff.Strings.PrepareToUpdate,
		diff.Strings.Equals,
	)

	err = s.repo.SoftDeleteInstancesByServer(ctx, diff.Deleted)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.CleanupDeletedInstances(ctx)
	if err != nil {
		return err
	}

	return nil
}

type FastFixParams struct {
	ServerID  string
	ChannelID string
}

func (s *Service) FastFixDiscordGuild(
	ctx context.Context,
	params *FastFixParams,
) error {
	err := s.EnsureInstancesForAllTypes(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FastFixDiscordGuildOnce(
	ctx context.Context,
	params *FastFixParams,
) {
	if params.ChannelID == "" {
		return
	}

	go func() {
		s.fastFixMutex.Lock()
		_, prevExec := s.fastFixExecutions[params.ServerID]

		if !prevExec {
			s.fastFixExecutions[params.ServerID] = struct{}{}
		}
		s.fastFixMutex.Unlock()

		if prevExec {
			return
		}

		err := s.FastFixDiscordGuild(ctx, params)
		if err != nil {
			telemetry.RecordError(ctx, err)

			s.fastFixMutex.Lock()
			delete(s.fastFixExecutions, params.ServerID)
			s.fastFixMutex.Unlock()
		}
	}()
}
