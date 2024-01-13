package models

type BotInstance struct {
	ID        int64
	ServerID  string
	ChannelID string
	Type      SubscriptionType
}
