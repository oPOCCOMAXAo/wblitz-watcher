package watcher

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

func Provide(
	i *do.Injector,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		discord, err := discord.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		wg, err := wg.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return NewService(
			discord,
			wg,
		), nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
