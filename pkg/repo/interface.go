package repo

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/mysql"
)

var _ Repository = (*mysql.Repository)(nil)

type Repository interface {
	CreateInstance(context.Context, *models.BotInstance) error
	GetInstanceByServer(context.Context, string) (*models.BotInstance, error)
	UpdateInstance(context.Context, *models.BotInstance) error
}
