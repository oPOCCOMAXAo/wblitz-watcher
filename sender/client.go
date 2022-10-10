package sender

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	client        *http.Client
	retryInterval time.Duration
	config        Config
}

type Config struct {
	Base string `env:"BASE,required"`
}

func New(config Config) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Minute,
		},
		retryInterval: 15 * time.Second,
		config:        config,
	}
}

type Request struct {
	URL  string      `json:"url"`
	Body interface{} `json:"body"`
}

func (c *Client) Request(ctx context.Context, request *Request) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.Base, bytes.NewReader(body))
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.Errorf("status: %d", res.StatusCode)
	}

	return nil
}

func (c *Client) RequestUntilSuccess(ctx context.Context, request *Request) {
	for {
		err := c.Request(ctx, request)
		if err == nil {
			return
		}

		log.Printf("%+v\n", err)
		time.Sleep(c.retryInterval)
	}
}
