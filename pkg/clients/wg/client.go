package wg

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/time/rate"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type Client struct {
	config  Config
	client  *http.Client
	limiter *rate.Limiter
}

type Config struct {
	ApplicationID string `env:"APPLICATION_ID,required"`
}

func New(config Config) *Client {
	res := Client{
		config: config,
		client: &http.Client{
			Timeout:   30 * time.Second,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
		limiter: rate.NewLimiter(rate.Limit(20), 10),
	}

	return &res
}

func (c *Client) getHost(region models.Region) string {
	switch region {
	case models.RegionAsia:
		return "api.wotblitz.asia"
	case models.RegionEU:
		return "api.wotblitz.eu"
	case models.RegionNA:
		return "api.wotblitz.com"
	case models.RegionRU:
		return "api.tanki.su"
	default:
		return "api.wotblitz.eu"
	}
}

func (c *Client) getFullURL(request *Request) string {
	if request.Data == nil {
		request.Data = url.Values{}
	}

	request.Data.Add("application_id", c.config.ApplicationID)

	builder := strings.Builder{}
	builder.Grow(256)
	builder.WriteString("https://")
	builder.WriteString(c.getHost(request.Region))
	builder.WriteString("/")
	builder.WriteString(string(request.App))
	builder.WriteString("/")
	builder.WriteString(string(request.Method))
	builder.WriteString("/?")
	builder.WriteString(request.Data.Encode())

	return builder.String()
}

func (c *Client) Request(
	ctx context.Context,
	request *Request,
	resPtr any,
) error {
	err := c.limiter.Wait(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getFullURL(request), http.NoBody)
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.Errorf("status: %d", res.StatusCode)
	}

	resObject := Response{
		Data: resPtr,
	}

	err = json.NewDecoder(res.Body).Decode(&resObject)
	if err != nil {
		return errors.WithStack(err)
	}

	if resObject.Error != nil {
		return resObject.Error.GetError()
	}

	return nil
}
