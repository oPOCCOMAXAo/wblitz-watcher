package watcher

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func Provide(
	i *do.Injector,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		telemetry, err := telemetry.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

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

		domain, err := domain.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return NewService(
			telemetry,
			discord,
			wg,
			domain,
		), nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
