package discord

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func Provide(
	i *do.Injector,
	config Config,
	env models.Environment,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		telemetry, err := telemetry.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return New(
			config,
			env,
			telemetry,
		)
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
