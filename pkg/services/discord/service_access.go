package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

func (s *Service) VerifyAccess(
	event *discordgo.InteractionCreate,
) error {
	if event.Member == nil {
		return models.ErrNoAccess
	}

	if s.isOwner(event) {
		return nil
	}

	if s.isAdmin(event) {
		return nil
	}

	return models.ErrNoAccess
}

func (s *Service) isOwner(
	event *discordgo.InteractionCreate,
) bool {
	var user *discordgo.User

	if event.Member != nil {
		user = event.Member.User
	}

	if user == nil {
		user = event.User
	}

	return user.ID == s.owner.ID
}

func (s *Service) isAdmin(
	event *discordgo.InteractionCreate,
) bool {
	if event.Member == nil {
		return false
	}

	if du.HasPermissions(event.Member.Permissions, discordgo.PermissionAdministrator) {
		return true
	}

	return false
}
