package sender

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type Client struct {
	client        *fasthttp.Client
	retryInterval time.Duration
	config        Config
}

type Config struct {
	Base string `env:"BASE,required"`
}

func New(config Config) *Client {
	return &Client{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
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
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(c.config.Base)
	req.Header.SetMethod(fasthttp.MethodPost)

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req.SetBodyRaw(data)

	err = c.client.Do(req, res)
	if err != nil {
		return err
	}

	code := res.StatusCode()
	if code != fasthttp.StatusOK {
		return errors.Errorf("status: %d", code)
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
