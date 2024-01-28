package telemetry

import (
	"context"

	"github.com/samber/do"
	"go.opentelemetry.io/otel/trace"
)

//nolint:revive // context here is parameter.
func Provide(
	i *do.Injector,
	appCtx context.Context,
	config Config,
) {
	do.Provide(i, func(i *do.Injector) (*Service, error) {
		res, err := NewService(appCtx, config)
		if err != nil {
			return nil, err
		}

		res.SetAsDefault()

		return res, nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.Invoke[*Service](i)
}

func InvokeSpan(i *do.Injector, name string) (context.Context, trace.Span, error) {
	svc, err := do.Invoke[*Service](i)
	if err != nil {
		return nil, nil, err
	}

	ctx, span := svc.TaskTracer().Start(context.Background(), name)

	return ctx, span, nil
}
