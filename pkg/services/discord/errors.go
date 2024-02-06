package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func MapError(
	err error,
) error {
	var discordErr *discordgo.RESTError

	if errors.As(err, &discordErr) {
		if discordErr.Message != nil {
			switch discordErr.Message.Code {
			case discordgo.ErrCodeMissingPermissions:
				return errors.Wrap(models.ErrNoAccess, discordErr.Message.Message)
			default:
				return errors.Wrap(models.ErrNotFound, discordErr.Message.Message)
			}
		}
	}

	return err
}
