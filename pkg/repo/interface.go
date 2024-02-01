package repo

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/mysql"
)

var _ Repository = (*mysql.Repository)(nil)

type Repository interface {
	Instance
	SubscriptionClan
	WGClan
	WGClanMember
	EventClan
	DiscordMessage
}

type Instance interface {
	CreateInstance(context.Context, *models.BotInstance) error
	GetInstance(context.Context, *models.BotInstance) (*models.BotInstance, error)
	UpdateInstance(context.Context, *models.BotInstance) error
	GetInstancesByType(context.Context, models.SubscriptionType) ([]*models.BotInstance, error)
}

type SubscriptionClan interface {
	CreateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	GetSubscriptionClan(context.Context, *models.SubscriptionClan) (*models.SubscriptionClan, error)
	UpdateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	DeleteSubscriptionClan(context.Context, *models.SubscriptionClan) error
	GetSubscriptionClanListByInstance(context.Context, int64) ([]*models.SubscriptionClan, error)
	UpdateIsDisabledSubscriptionClanByID(
		_ context.Context,
		isDisabled bool,
		ids []int64,
	) error
}

type WGClan interface {
	CreateUpdateWGClan(context.Context, *models.WGClan) error
	GetWGClan(context.Context, *models.WGClan) (*models.WGClan, error)
	UpdateWGClansMembersUpdatedAt(context.Context, int64, []models.WGClanID) error
	GetWGClanListByInstance(context.Context, int64) ([]*models.WGClan, error)
	GetWGClansWithSubscriptions(context.Context) ([]*models.WGClan, error)
}

type WGClanMember interface {
	CreateUpdateWGClanMembers(context.Context, []*models.WGClanMember) error
	DeleteWGClanMembers(context.Context, []*models.WGClanMember) error
	GetWGClanMembers(context.Context, []models.WGClanID) ([]*models.WGClanMembers, error)
}

type EventClan interface {
	CreateEventClan(context.Context, ...*models.EventClan) error
	GetEventClanByID(context.Context, int64) (*models.EventClan, error)
	UpdateEventClanProcessed(context.Context) error
	DeleteEventClansByID(context.Context, []int64) error
}

type DiscordMessage interface {
	CreateDiscordMessagesFromEventClan(context.Context) error
	GetFirstUnsentDiscordMessage(context.Context) (*models.DiscordMessage, error)
	UpdateDiscordMessagesProcessed(context.Context, []int64) error
}
