package repo

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/mysql"
)

var _ Repository = (*mysql.Repository)(nil)

type Repository interface {
	CreateInstance(context.Context, *models.BotInstance) error
	GetInstance(context.Context, *models.BotInstance) (*models.BotInstance, error)
	UpdateInstance(context.Context, *models.BotInstance) error

	CreateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	GetSubscriptionClan(context.Context, *models.SubscriptionClan) (*models.SubscriptionClan, error)
	UpdateSubscriptionClan(context.Context, *models.SubscriptionClan) error
	DeleteSubscriptionClan(context.Context, *models.SubscriptionClan) error
}
