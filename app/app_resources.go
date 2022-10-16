package app

import (
	"github.com/opoccomaxao/wblitz-watcher/repo"
	"github.com/opoccomaxao/wblitz-watcher/sender"
	"github.com/opoccomaxao/wblitz-watcher/utils"
	"github.com/opoccomaxao/wblitz-watcher/wg/api"
	"github.com/opoccomaxao/wblitz-watcher/wg/client"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type ClansConfig struct {
	EU   []int `env:"EU" envSeparator:","`
	RU   []int `env:"RU" envSeparator:","`
	NA   []int `env:"NA" envSeparator:","`
	Asia []int `env:"ASIA" envSeparator:","`
}

func (c *ClansConfig) GetRegion(region api.Region) []int {
	return utils.GetFromMap(map[api.Region][]int{
		api.RegionAsia: c.Asia,
		api.RegionEU:   c.EU,
		api.RegionNA:   c.NA,
		api.RegionRU:   c.RU,
	}, region, api.RegionUnknown)
}

type Config struct {
	ApplicationID string        `env:"APP_ID,required"`
	Clans         ClansConfig   `envPrefix:"CLANS_"`
	DB            repo.Config   `envPrefix:"DB_"`
	Sender        sender.Config `envPrefix:"SENDER_"`
	DiscordURL    string        `env:"DISCORD_URL,required"`
}

func (app *App) initConfig() error {
	_ = godotenv.Load(".env")

	err := env.Parse(&app.config)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (app *App) initClient() error {
	var err error

	app.client, err = client.New(client.Config{
		ApplicationID: app.config.ApplicationID,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (app *App) initRepo() error {
	var err error

	app.repo, err = repo.New(app.config.DB)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (app *App) initSender() error {
	app.sender = sender.New(app.config.Sender)

	return nil
}

func (app *App) Close() error {
	return nil
}
