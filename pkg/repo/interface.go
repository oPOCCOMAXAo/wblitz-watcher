package repo

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/mysql"
)

var _ Repository = (*mysql.Repository)(nil)

//nolint:interfacebloat
type Repository interface {
	CreateInstance(context.Context, *models.BotInstance) error
	GetInstance(context.Context, *models.BotInstance) (*models.BotInstance, error)
	UpdateInstance(context.Context, *models.BotInstance) error

	CreateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	GetSubscriptionClan(context.Context, *models.SubscriptionClan) (*models.SubscriptionClan, error)
	UpdateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	DeleteSubscriptionClan(context.Context, *models.SubscriptionClan) error

	CreateUpdateWGClan(context.Context, *models.WGClan) error
	GetWGClan(context.Context, *models.WGClan) (*models.WGClan, error)
	UpdateWGClansMembersUpdatedAt(context.Context, int64, []models.WGClanID) error
	GetWGClanListByInstance(context.Context, int64) ([]*models.WGClan, error)
}
