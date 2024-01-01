package config

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/server"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/watcher"
)

type Config struct {
	Server    server.Config    `envPrefix:"SERVER_"`
	Discord   discord.Config   `envPrefix:"DISCORD_"`
	Telemetry telemetry.Config `envPrefix:"TELEMETRY_"`
	WG        wg.Config        `envPrefix:"WG_"`
}

//nolint:revive // ctx here is value.
func (c *Config) Provide(
	i *do.Injector,
	appCtx context.Context,
) {
	server.Provide(i, c.Server)
	discord.Provide(i, c.Discord)
	telemetry.Provide(i, appCtx, c.Telemetry)
	watcher.Provide(i)
	wg.Provide(i, c.WG)
}

func Load() (*Config, error) {
	var res Config

	err := env.Parse(&res)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &res, nil
}