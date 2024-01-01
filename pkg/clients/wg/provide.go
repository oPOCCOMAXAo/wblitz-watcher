package wg

import (
	"github.com/samber/do"
)

func Provide(
	i *do.Injector,
	config Config,
) {
	do.Provide(i, func(i *do.Injector) (*Client, error) {
		return New(config), nil
	})
}

func Invoke(i *do.Injector) (*Client, error) {
	return do.Invoke[*Client](i)
}
