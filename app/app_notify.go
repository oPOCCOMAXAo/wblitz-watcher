package app

import (
	"context"
	"log"
	"time"

	"github.com/opoccomaxao/wblitz-watcher/app/diff"
	"github.com/opoccomaxao/wblitz-watcher/sender"
	"github.com/opoccomaxao/wblitz-watcher/sender/discord"
	"github.com/opoccomaxao/wblitz-watcher/wg/types"

	"github.com/opoccomaxao-go/generic-collection/set"
)

func (app *App) NotifyClanDiff(
	ctx context.Context,
	clanInfo *types.ClanInfo,
	clanDiff *diff.Total,
) error {
	log.Printf("%s\n", clanDiff.Pretty())

	playerIDs := set.New[int]()

	for _, diff := range clanDiff.Int {
		switch diff.Type {
		case DiffLeave, DiffEnter, DiffLeader:
			playerIDs.Add(diff.Old)
			playerIDs.Add(diff.New)
		}
	}

	players, err := app.client.AccountInfo(ctx, app.config.Region, playerIDs.Slice()...)
	if err != nil {
		return err
	}

	for _, diff := range clanDiff.Int {
		switch diff.Type {
		case DiffLeave, DiffEnter, DiffLeader:
			event := app.buildClanEvent(clanInfo, players[types.MaybeInt(diff.New)], diff.Type)

			err := app.SendDiscordMessage(ctx, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *App) buildClanEvent(
	clan *types.ClanInfo,
	player *types.AccountInfo,
	eventType diff.Type,
) *discord.Message {
	embed := discord.Embed{
		Type:      discord.EmbedTypeRich,
		Timestamp: time.Now(),
	}

	switch eventType {
	case DiffEnter:
		embed.Color = discord.ColorEnter
		embed.Description = StringEnter
	case DiffLeave:
		embed.Color = discord.ColorLeave
		embed.Description = StringLeave
	default:
		return nil
	}

	if clan != nil {
		embed.Fields = append(embed.Fields,
			discord.Field{
				Inline: false,
				Name:   "Клан",
				Value:  clan.StatName(),
			},
		)
	}

	if player != nil {
		embed.Author = &discord.Author{
			Name: player.AuthorName(),
		}

		embed.Fields = append(embed.Fields,
			discord.Field{
				Inline: true,
				Name:   "Победы",
				Value:  player.StatWins(),
			},
			discord.Field{
				Inline: true,
				Name:   "Урон",
				Value:  player.StatDamage(),
			},
			discord.Field{
				Inline: true,
				Name:   "Бои",
				Value:  player.StatBattles(),
			},
		)
	}

	return &discord.Message{
		Embeds: []discord.Embed{embed},
	}
}

func (app *App) SendDiscordMessage(ctx context.Context, message *discord.Message) error {
	go app.sender.RequestUntilSuccess(ctx, &sender.Request{
		URL:  app.config.DiscordURL,
		Body: message,
	})

	return nil
}
