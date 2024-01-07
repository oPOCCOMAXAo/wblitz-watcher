package domain

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/repo"
)

func Provide(
	i *do.Injector,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		repo, err := repo.InvokeMySQL(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return NewService(
			repo,
		), nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
