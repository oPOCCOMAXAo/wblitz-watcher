package models

type DiscordMessage struct {
	ID          int64
	IsProcessed bool
	EventClanID int64
	InstanceID  int64

	EventClan
	BotInstance
}
