package discord

import (
	"context"
)

type EventName string

const (
	EventGuildDelete EventName = "onGuildDelete"
)

type Event struct {
	GuildID string
}

type EventHandler func(context.Context, *Event) error
