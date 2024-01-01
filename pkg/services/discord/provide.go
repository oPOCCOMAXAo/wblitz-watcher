package discord

import "github.com/samber/do"

func Provide(
	i *do.Injector,
	config Config,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		return New(config)
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}
