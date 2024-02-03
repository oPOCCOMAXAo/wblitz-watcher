package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type EventName string

const (
	EventGuildCreate EventName = "onGuildCreate"
	EventGuildDelete EventName = "onGuildDelete"
	EventReady       EventName = "onReady"
)

type Event struct {
	Guild  *discordgo.Guild
	Guilds []*discordgo.Guild
}

type EventHandler func(context.Context, *Event) error
