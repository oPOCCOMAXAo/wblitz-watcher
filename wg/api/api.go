package api

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

type API struct {
	client  *fasthttp.Client
	limiter *rate.Limiter
	config  Config
}

type Config struct {
	Timeout       time.Duration
	ApplicationID string
}

func New(config Config) (*API, error) {
	if config.ApplicationID == "" {
		return nil, errors.New("ApplicationID required")
	}

	if config.Timeout < 0 {
		config.Timeout = 0
	}

	return &API{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			ReadTimeout:              config.Timeout,
			WriteTimeout:             config.Timeout,
		},
		limiter: rate.NewLimiter(rate.Limit(20), 10),
		config:  config,
	}, nil
}

func (a *API) makeURL(request *Request) string {
	if request.Data == nil {
		request.Data = url.Values{}
	}

	request.Data.Add("application_id", a.config.ApplicationID)

	builder := strings.Builder{}
	builder.Grow(128)
	builder.WriteString("https://")
	builder.WriteString(request.Region.Host())
	builder.WriteString("/")
	builder.WriteString(string(request.App))
	builder.WriteString("/")
	builder.WriteString(string(request.Method))
	builder.WriteString("/?")
	builder.WriteString(request.Data.Encode())

	return builder.String()
}

func (a *API) Request(
	ctx context.Context,
	request *Request,
	resPtr interface{},
) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	err := a.limiter.Wait(ctx)
	if err != nil {
		return err
	}

	req.SetRequestURI(a.makeURL(request))

	err = a.client.Do(req, res)
	if err != nil {
		return err
	}

	code := res.StatusCode()
	if code != fasthttp.StatusOK {
		return errors.Errorf("status: %d", code)
	}

	resObject := Response{
		Data: resPtr,
	}

	err = json.Unmarshal(res.Body(), &resObject)
	if err != nil {
		return err
	}

	if resObject.Error != nil {
		return resObject.Error.GetError()
	}

	return nil
}
