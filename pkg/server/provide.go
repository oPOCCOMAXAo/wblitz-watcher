package server

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func Provide(
	i *do.Injector,
	config Config,
) {
	do.Provide(i, func(i *do.Injector) (*Server, error) {
		telemetry, err := telemetry.Invoke(i)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return New(
			config,
			telemetry,
		), nil
	})
}

func Invoke(i *do.Injector) (*Server, error) {
	return do.Invoke[*Server](i)
}
