package domain

import "context"

func (s *Service) DeleteDiscordGuildData(
	ctx context.Context,
	serverID string,
) error {
	err := s.repo.SoftDeleteInstancesByServer(ctx, serverID)
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
