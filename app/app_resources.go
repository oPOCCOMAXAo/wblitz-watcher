package app

import (
	"github.com/opoccomaxao/wblitz-watcher/repo"
	"github.com/opoccomaxao/wblitz-watcher/sender"
	"github.com/opoccomaxao/wblitz-watcher/wg/api"
	"github.com/opoccomaxao/wblitz-watcher/wg/client"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	ApplicationID string        `env:"APP_ID,required"`
	Clans         []int         `env:"CLANS" envSeparator:","`
	DB            repo.Config   `envPrefix:"DB_"`
	Sender        sender.Config `envPrefix:"SENDER_"`
	DiscordURL    string        `env:"DISCORD_URL,required"`
	Region        api.Region
}

func (app *App) initConfig() error {
	_ = godotenv.Load(".env")

	err := env.Parse(&app.config)
	if err != nil {
		return errors.WithStack(err)
	}

	app.config.Region = api.RegionEU

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
