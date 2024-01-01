package migrations

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func Init(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.BotInstance{},
		&models.SubscriptionClan{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
