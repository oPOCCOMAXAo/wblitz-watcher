package telemetry

import (
	"context"

	"github.com/samber/do"
)

func Provide(
	i *do.Injector,
	appCtx context.Context,
	config Config,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		res, err := NewService(appCtx, config)
		if err != nil {
			return nil, err
		}

		res.SetAsDefault()

		return res, nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
