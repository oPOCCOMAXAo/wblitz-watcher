package domain

import "context"

const ClanLimitNoPremium = 25

func (s *Service) GetSubscriptionClanLimitForServer(
	ctx context.Context,
	serverID string,
) (int64, error) {
	return ClanLimitNoPremium, nil
}
