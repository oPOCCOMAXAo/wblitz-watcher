package config

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo"
	"github.com/opoccomaxao/wblitz-watcher/pkg/server"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/watcher"
)

type Config struct {
	Environment models.Environment `env:"ENVIRONMENT"      envDefault:"development"`
	Server      server.Config      `envPrefix:"SERVER_"`
	Discord     discord.Config     `envPrefix:"DISCORD_"`
	Telemetry   telemetry.Config   `envPrefix:"TELEMETRY_"`
	WG          wg.Config          `envPrefix:"WG_"`
	DB          repo.Config        `envPrefix:"DB_"`
}

//nolint:revive // ctx here is value.
func (c *Config) Provide(
	i *do.Injector,
	appCtx context.Context,
) {
	server.Provide(i, c.Server)
	discord.Provide(i, c.Discord, c.Environment)
	telemetry.Provide(i, appCtx, c.Telemetry)
	watcher.Provide(i)
	wg.Provide(i, c.WG)
	repo.Provide(i, c.DB)
	domain.Provide(i)
}

func Load() (*Config, error) {
	var res Config

	err := env.Parse(&res)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &res, nil
}
