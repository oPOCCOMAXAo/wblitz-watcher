package telemetry

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	provider *sdktrace.TracerProvider
}

type Config struct {
	Service  string `env:"SERVICE"`
	Endpoint string `env:"ENDPOINT"`
	APIKEY   string `env:"APIKEY"`
	Insecure bool   `env:"INSECURE"`
}

func NewService(ctx context.Context, config Config) (*Service, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(config.Endpoint),
	}

	if config.APIKEY != "" {
		opts = append(opts, otlptracehttp.WithHeaders(map[string]string{
			"api-key": config.APIKEY,
		}))
	}

	if config.Insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	client := otlptracehttp.NewClient(opts...)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	service := Service{}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", config.Service),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	service.provider = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	return &service, nil
}

func (s *Service) SetAsDefault() {
	otel.SetTracerProvider(s.provider)
	errTracer = s.provider.Tracer("error")
}

func (s *Service) Tracer() trace.Tracer {
	return s.provider.Tracer(
		"trading",
	)
}

func (s *Service) TracerWithOptions(name string, opts ...trace.TracerOption) trace.Tracer {
	return s.provider.Tracer(name, opts...)
}

func (s *Service) PackageTracer(pkgName string) trace.Tracer {
	return NewTracerWithOptions(s.Tracer(), pkgName, []trace.SpanStartOption{})
}

func (s *Service) TaskTracer() trace.Tracer {
	return NewTracerWithOptions(s.Tracer(), "task", []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
	})
}

func (s *Service) Close(ctx context.Context) error {
	err := s.provider.Shutdown(ctx)

	return errors.WithStack(err)
}

func (s *Service) TestContext(t *testing.T) context.Context {
	ctx, span := s.Tracer().Start(context.Background(), "test."+t.Name())
	t.Cleanup(func() { span.End() })

	return ctx
}
