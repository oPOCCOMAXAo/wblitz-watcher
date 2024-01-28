package watcher

import "github.com/opoccomaxao/wblitz-watcher/pkg/models"

func (s *Service) isClanMembersInitialized(
	clan *models.WGClan,
) bool {
	return clan.MembersUpdatedAt > s.now()-ClanInitializationIntervalSeconds
}
