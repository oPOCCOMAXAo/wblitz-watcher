package client

import (
	"time"

	"github.com/opoccomaxao/wblitz-watcher/wg/api"
)

type Client struct {
	api *api.API
}

type Config struct {
	ApplicationID string `env:"APP_ID"`
}

func New(config Config) (*Client, error) {
	var err error

	res := Client{}

	res.api, err = api.New(api.Config{
		Timeout:       time.Second * 30,
		ApplicationID: config.ApplicationID,
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}
