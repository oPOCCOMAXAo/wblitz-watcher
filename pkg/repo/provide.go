package repo

import (
	"github.com/pkg/errors"
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/db"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/migrations"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo/mysql"
)

func Provide(
	i *do.Injector,
	config Config,
) {
	do.ProvideNamed[Repository](i, "MySQL", func(i *do.Injector) (Repository, error) {
		db, err := db.OpenMySQL(config.RepoDSN)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = migrations.Init(db)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return mysql.New(db), nil
	})
}

func InvokeMySQL(i *do.Injector) (Repository, error) {
	return do.InvokeNamed[Repository](i, "MySQL")
}
